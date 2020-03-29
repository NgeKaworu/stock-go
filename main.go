package main

import (
	"flag"
	"log"
	"os"
	"stock/src/constants"
	"stock/src/dbengin"
	"stock/src/stock"
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

	eng := dbengin.NewDbEngine()
	err := eng.Open(*mongo, *db, *dbinit)

	if err != nil {
		log.Println(err.Error())
	}

	stocks := utils.Merge(constants.Ss50, constants.Hs300)
	for k, v := range stocks {
		s := &stock.Stock{
			Code:       k,
			BourseCode: v,
		}
		switch v {
		case "01":
			s.Bourse = "sh"
		case "02":
			s.Bourse = "sz"
		default:
			break
		}
		// s.FetchClassify()
		log.Printf("%+v\n", s)
		break
	}
	// for k, v := range constants.Ss50 {
	// 	log.Println(k, v)
	//
	// }

}
