package main

import (
	"flag"
	"log"
	"os"
	"reflect"
	"stock/src/controllers"
	"stock/src/models"
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

	// resp, err := http.Get("https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart2Get?fc=60000001&color=w")

	// body, err := ioutil.ReadAll(resp.Body)
	// defer resp.Body.Close()

	// result := map[string]interface{}{}
	// err = json.Unmarshal(body, &result)

	// log.Println(result["Result"].(map[string]interface{})["TiCaiXiangQingList"].([]interface{})[0].(map[string]interface{})["KeyWord"])

	ci := &models.CurrentInfo{}

	st := reflect.TypeOf(ci).Elem()
	for i := 0; i < st.NumField(); i++ {
		log.Println(st.Field(i).Tag)
	}

	// stocks := utils.Merge(constants.Ss50, constants.Hs300)
	// log.Println(stocks)
	// eng.FetchCurrentInfo()
	// for k, v := range constants.Ss50 {
	// 	log.Println(k, v)
	// 	// eng.FetchMainIndicator(v)
	// }

}
