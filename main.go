package main

import (
	"flag"
	"log"
	"os"
	"stock/src/constants"
	"stock/src/controllers"
	"stock/src/utils"
)

func main() {
	var (
		dbinit = flag.Bool("i", false, "init database flag")
		mongo  = flag.String("m", "mongodb://localhost:27017", "mongod addr flag")
		db     = flag.String("db", "stock", "database name")
	)
	flag.Parse()

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	eng := controllers.NewDbEngine()
	err := eng.Open(*mongo, *db, *dbinit)

	if err != nil {
		log.Println(err.Error())
	}

	stocks := utils.Merge(constants.Ss50, constants.Hs300)
	log.Println(stocks)
	// for k, v := range constants.Ss50 {
	// 	log.Println(k, v)
	// 	// eng.FetchMainIndicator(v)
	// }

}
