package engine

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/NgeKaworu/stock/src/models"
	"github.com/NgeKaworu/stock/src/resultor"
	"github.com/NgeKaworu/stock/src/stock"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchCurrent 爬取当前信息
func (d *DbEngine) FetchCurrent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	allMarket := make([]interface{}, 0)
	now := time.Now().Local()
	pool := make(chan bool, 10)
	for k, v := range stock.Stocks {
		pool <- true
		go func(key, val string) {
			s := stock.NewStock(key, val)
			ciPar := *s.Bourse + *s.Code
			ciRes, err := http.Get("http://hq.sinajs.cn/list=" + ciPar)
			if err != nil {
				log.Println(err.Error())
			}
			// 中文编码
			utf8Reader := transform.NewReader(ciRes.Body, simplifiedchinese.GBK.NewDecoder())
			body, err := ioutil.ReadAll(utf8Reader)
			if err != nil {
				log.Fatal(err)
			}
			defer ciRes.Body.Close()
			// 股票名称、今日开盘价、昨日收盘价、当前价格、今日最高价、今日最低价、竞买价、竞卖价、成交股数、成交金额、买1手、买1报价、买2手、买2报价、…、买5报价、…、卖5报价、日期、时间
			strArr := strings.Split(string(body), ",")

			ci := models.CurrentInfo{}

			st := reflect.ValueOf(ci).Elem()
			for k, v := range strArr[:len(strArr)-3] {
				if k == 0 {
					st.Field(k).SetPointer(unsafe.Pointer(&strings.Split(v, "\"")[1]))
					continue
				}
				st.Field(k).SetPointer(unsafe.Pointer(&v))

			}

			clsPar := *s.Code + *s.BourseCode
			clsRes, err := http.Get("https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart2Get?fc=" + clsPar)

			body, err = ioutil.ReadAll(clsRes.Body)
			defer clsRes.Body.Close()

			result := map[string]interface{}{}
			err = json.Unmarshal(body, &result)

			if err != nil {
				log.Println(err)
			}

			if r, ok := result["Result"].(map[string]interface{}); ok {
				if tiCaiXiangQingList, ok := r["TiCaiXiangQingList"]; ok {
					for _, tiCaiXiangQing := range tiCaiXiangQingList.([]interface{}) {
						if keyWord, ok := tiCaiXiangQing.(map[string]interface{})["KeyWord"].(string); ok {
							ci.Classify = &keyWord
							break
						}
					}
				}

			}

			ci.Code = s.Code
			ci.CreateDate = &now
			allMarket = append(allMarket, ci)
			<-pool
		}(k, v)
	}

	tCurrentInfo := d.GetColl(models.TCurrentInfo)
	_, err := tCurrentInfo.InsertMany(context.Background(), allMarket)

	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	resultor.RetOk(w, &allMarket)
}

// ListCurrent 现值列表
func (d *DbEngine) ListCurrent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	date, err := time.Parse(time.RFC3339, ps.ByName("date"))
	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}
	t := d.GetColl(models.TCurrentInfo)
	res, err := t.Find(context.Background(), bson.M{"create_date": date})

	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	a := make([]models.CurrentInfo, 0)

	err = res.All(context.Background(), &a)

	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	resultor.RetOk(w, &a)
}

// ListInfoTime 爬取时间列表
func (d *DbEngine) ListInfoTime(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	query := []bson.M{
		{"$group": bson.M{
			"_id": "$create_date",
		}},
	}
	tCurrentInfo := d.GetColl(models.TCurrentInfo)
	re, err := tCurrentInfo.Aggregate(context.Background(), query, options.Aggregate())
	if err != nil {
		resultor.RetFail(w, err.Error())
	}
	times := make([]map[string]time.Time, 0)
	err = re.All(context.Background(), &times)

	if err != nil {
		resultor.RetFail(w, err.Error())
	}

	resultor.RetOk(w, times)
}
