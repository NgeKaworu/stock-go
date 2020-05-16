package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"stock/src/cors"
	"stock/src/dbengin"
	"syscall"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		addr   = flag.String("l", ":8000", "绑定Host地址")
		dbinit = flag.Bool("i", false, "init database flag")
		mongo  = flag.String("m", "mongodb://localhost:27017", "mongod addr flag")
		db     = flag.String("db", "stock", "database name")
	)
	flag.Parse()

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	eng := dbengin.NewDbEngine()
	err = eng.Open(*mongo, *db, *dbinit)

	if err != nil {
		log.Println(err.Error())
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, "schema.graphql"))
	if err != nil {
		log.Fatal(err.Error())
	}

	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
		graphql.MaxParallelism(1000),
		//生产环境需启动禁调试功能
		//	opts = append(opts, graphql.DisableIntrospection())
	}

	schema := graphql.MustParseSchema(string(b), eng, opts...)
	mux := http.NewServeMux()
	mux.Handle("/", &relay.Handler{Schema: schema})

	srv := &http.Server{Handler: cors.CORS(mux), ErrorLog: nil}
	srv.Addr = *addr

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("server on http port", srv.Addr)

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	cleanup := make(chan bool)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for range signalChan {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			go func() {
				_ = srv.Shutdown(ctx)
				cleanup <- true
			}()
			<-cleanup
			eng.Close()
			fmt.Println("safe exit")
			cleanupDone <- true
		}
	}()
	<-cleanupDone

}
