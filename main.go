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
		db     = flag.String("db", "stock_test", "database name")
		// pb     = flag.Int("pb", 2, "pb weight")
		// pe     = flag.Int("pe", 5, "pe weight")
		// peg    = flag.Int("peg", 6, "peg weight")
		// roe    = flag.Int("roe", 8, "roe weight")
		// dpe    = flag.Int("dpe", 8, "dpe weight")
		// dce    = flag.Int("dce", 5, "dce weight")
		// aagr   = flag.Int("aagr", 8, "aagr weight")
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

	// // 风险收益率(Rate of Risked Return)
	// // 假设10年内 > 80% 30年内 < 20%
	// RRR := 0.086
	// // 通货
	// CPI := 0.052
	// // 无风险利率 (The risk-free rate of interest)
	// RFR := 0.0285
	// discount := RRR + CPI + RFR

	// weights := map[string][]interface{}{
	// 	"PB":   {*pb, false},
	// 	"PE":   {*pe, true},
	// 	"PEG":  {*peg, true},
	// 	"ROE":  {*roe, true},
	// 	"DPE":  {*dpe, true},
	// 	"DCE":  {*dce, true},
	// 	"AAGR": {*aagr, true},
	// }
	// total := float64(*pb + *pe + *peg + *roe + *dpe + *dce + *aagr)

	// stocks := utils.Merge(constants.Ss50, constants.Hs300)

	// allStock := make([]stock.Stock, 0)
	// allReport := make([]interface{}, 0)
	// allMarket := make([]interface{}, 0)
	// now := time.Now().Local()
	// for k, v := range stocks {
	// 	s := &stock.Stock{
	// 		Code:       k,
	// 		BourseCode: v,
	// 	}
	// 	switch v {
	// 	case "01":
	// 		s.Bourse = "sh"
	// 	case "02":
	// 		s.Bourse = "sz"
	// 	default:
	// 		break
	// 	}
	// 	log.Println("current code: " + k)
	// 	s.FetchCurrentInfo()
	// 	s.FetchMainIndicator()
	// 	s.FetchClassify()
	// 	s.Calc()
	// 	s.Discount(discount)
	// 	s.CreateDate = now

	// 	s.CurrentInfo.Code = s.Code
	// 	s.CurrentInfo.CreateDate = now

	// 	allStock = append(allStock, *s)
	// }

	// stock.WeightSort(weights, &allStock, total)

	// insertStock := make([]interface{}, 0)
	// for _, v := range allStock {
	// 	insertStock = append(insertStock, v)
	// }

	// tStock := eng.GetColl(stock.TStock)
	// if ret, err := tStock.InsertMany(context.Background(), insertStock); err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Println(ret)
	// }

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
