package main

import (
	"flag"
	"log"
	"os"
	"stock/src/dbengin"
)

type Some struct {
	I int     `bson:"test_one"`
	B float32 `bson:"-"`
	P *int    `bson:"test_ponint,omitempty"`
}

func main() {
	var (
		dbinit = flag.Bool("i", false, "init database flag")
		mongo  = flag.String("m", "mongodb://localhost:27017", "mongod addr flag")
		db     = flag.String("db", "stock", "database name")
		pb     = flag.Int("pb", 2, "pb weight")
		// pe     = flag.Int("pe", 5, "pe weight")
		// peg    = flag.Int("peg", 6, "peg weight")
		// roe    = flag.Int("roe", 8, "roe weight")
		// dpe    = flag.Int("dpe", 5, "dpe weight")
		// dce    = flag.Int("dce", 8, "dce weight")
		// aagr   = flag.Int("aagr", 8, "aagr weight")
	)
	flag.Parse()

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(*pb)

	eng := dbengin.NewDbEngine()
	err := eng.Open(*mongo, *db, *dbinit)

	if err != nil {
		log.Println(err.Error())
	}
	// 风险收益率(Rate of Risked Return)
	// 假设10年内 > 80% 30年内 < 20%
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

	// l := len(stocks)
	// ss := make([]stock.Stock, l)

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
	// 	s.FetchCurrentInfo()
	// 	s.FetchMainIndicator()
	// 	s.FetchClassify()
	// 	s.Calc()
	// 	s.Discount(discount)

	// 	ss = append(ss, *s)
	// }

	// stock.WeightSort(weights, &ss, total)

	// d := eng.GetColl("test")

}
