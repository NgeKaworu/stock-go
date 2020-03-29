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
	// 风险收益率(Rate of Risked Return)
	// 假设10年内 > 80% 30年内 < 20%
	RRR := 0.086
	// 通货
	CPI := 0.052
	// 无风险利率 (The risk-free rate of interest)
	RFR := 0.0285
	discount := RRR + CPI + RFR
	log.Println(discount)
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
		s.FetchCurrentInfo()
		s.FetchMainIndicator()
		s.FetchClassify()
		s.Calc()

		log.Printf("%+v\n", s)
		break
	}
	// for k, v := range constants.Ss50 {
	// 	log.Println(k, v)
	//
	// }

}
