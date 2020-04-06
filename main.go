package main

import (
	"context"
	"flag"
	"log"
	"os"
	"stock/src/dbengin"
	"stock/src/models"
	"stock/src/stock"
	"time"
)

func main() {
	var (
		dbinit = flag.Bool("i", false, "init database flag")
		mongo  = flag.String("m", "mongodb://localhost:27017", "mongod addr flag")
		db     = flag.String("db", "stock", "database name")
		pb     = flag.Int("pb", 2, "pb weight")
		pe     = flag.Int("pe", 5, "pe weight")
		peg    = flag.Int("peg", 6, "peg weight")
		roe    = flag.Int("roe", 8, "roe weight")
		dpe    = flag.Int("dpe", 5, "dpe weight")
		dce    = flag.Int("dce", 8, "dce weight")
		aagr   = flag.Int("aagr", 8, "aagr weight")
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

	weights := map[string][]interface{}{
		"PB":   {*pb, false},
		"PE":   {*pe, true},
		"PEG":  {*peg, true},
		"ROE":  {*roe, true},
		"DPE":  {*dpe, true},
		"DCE":  {*dce, true},
		"AAGR": {*aagr, true},
	}
	total := float64(*pb + *pe + *peg + *roe + *dpe + *dce + *aagr)

	// stocks := utils.Merge(constants.Ss50, constants.Hs300)

	stocks := map[string]string{
		"600000": "01", //浦发银行
	}

	allStock := make([]stock.Stock, 0)
	allReport := make([]interface{}, 0)
	allMarket := make([]interface{}, 0)
	now := time.Now().Local()
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
		log.Println("current code: " + k)
		s.FetchCurrentInfo()
		s.FetchMainIndicator()
		s.FetchClassify()
		s.Calc()
		s.Discount(discount)
		s.CreateDate = now

		s.CurrentInfo.Code = s.Code
		s.CurrentInfo.CreateDate = now

		allStock = append(allStock, *s)
		allMarket = append(allMarket, *s.CurrentInfo)

		for _, enterprise := range *s.Enterprise {
			enterprise.CreateDate = time.Now().Local()
			enterprise.Code = s.Code
			allReport = append(allReport, enterprise)
		}

	}

	stock.WeightSort(weights, &allStock, total)

	insertStock := make([]interface{}, 0)
	for _, v := range allStock {
		insertStock = append(insertStock, v)
	}

	log.Printf("%+v\n", insertStock)
	log.Printf("%+v\n", allReport)
	log.Printf("%+v\n", allMarket)

	tStock := eng.GetColl(stock.TStock)
	if ret, err := tStock.InsertMany(context.Background(), insertStock); err != nil {
		log.Println(err)
	} else {
		log.Println(ret)
	}

	tCurrentInfo := eng.GetColl(models.TCurrentInfo)
	if ret, err := tCurrentInfo.InsertMany(context.Background(), allMarket); err != nil {
		log.Println(err)
	} else {
		log.Println(ret)
	}

	tEnterpriseIndicator := eng.GetColl(models.TEnterpriseIndicator)
	if ret, err := tEnterpriseIndicator.InsertMany(context.Background(), allReport); err != nil {
		log.Println(err)
	} else {
		log.Println(ret)
	}

}
