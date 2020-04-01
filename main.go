package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"stock/src/dbengin"
	"stock/src/stock"
)

func main() {
	var (
		dbinit = flag.Bool("i", false, "init database flag")
		mongo  = flag.String("m", "mongodb://localhost:27017", "mongod addr flag")
		db     = flag.String("db", "stock", "database name")
		pb     = flag.Int("pb", 2, "pb weight")
		pe     = flag.Int("pe", 8, "pe weight")
		peg    = flag.Int("peg", 0, "peg weight")
		roe    = flag.Int("roe", 0, "roe weight")
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

	s := make([]stock.Stock, 5)
	s[0] = stock.Stock{Code: "01", Bourse: "sh", BourseCode: "01", Enterprise: nil, CurrentInfo: nil, Classify: "nil", PB: 100.0, PE: 20.0, PEG: 10.0, ROE: 13.0, DPE: 56.0, DCE: 47.0, AAGR: 23.0, Grade: 0}
	s[1] = stock.Stock{Code: "01", Bourse: "sh", BourseCode: "01", Enterprise: nil, CurrentInfo: nil, Classify: "nil", PB: 80.0, PE: 40.0, PEG: 80.0, ROE: 1456.0, DPE: 78.0, DCE: 15.0, AAGR: 14.0, Grade: 0}
	s[2] = stock.Stock{Code: "01", Bourse: "sh", BourseCode: "01", Enterprise: nil, CurrentInfo: nil, Classify: "nil", PB: 90.0, PE: 50.0, PEG: 90.0, ROE: 456.0, DPE: 36.0, DCE: 47.0, AAGR: 665.0, Grade: 0}
	s[3] = stock.Stock{Code: "01", Bourse: "sh", BourseCode: "01", Enterprise: nil, CurrentInfo: nil, Classify: "nil", PB: 60.0, PE: 30.0, PEG: 20.0, ROE: 312.0, DPE: 65.0, DCE: 53.0, AAGR: 51.0, Grade: 0}
	s[4] = stock.Stock{Code: "01", Bourse: "sh", BourseCode: "01", Enterprise: nil, CurrentInfo: nil, Classify: "nil", PB: 50.0, PE: 10.0, PEG: 30.0, ROE: 35.0, DPE: 85.0, DCE: 41.0, AAGR: 45.0, Grade: 0}
	fmt.Println("初始化结果:\n", s)

	weights := map[string][]interface{}{
		"pb":  {*pb, false},
		"pe":  {*pe, true},
		"peg": {*peg, true},
		"roe": {*roe, true},
	}
	total := float64(*pb + *pe + *peg + *roe)
	// 从小到大排序(不稳定排序)
	stock.WeightSort(weights, s, total)
	fmt.Println("\n从小到大排序结果:")
	fmt.Println(s)
	// log.Println(discount)
	// stocks := utils.Merge(constants.Ss50, constants.Hs300)
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
}
