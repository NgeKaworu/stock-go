package engine

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/NgeKaworu/stock/src/models"
	"github.com/NgeKaworu/stock/src/resultor"
	"github.com/NgeKaworu/stock/src/stock"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

// FetchEnterprise 爬年报 并返回
func (d *DbEngine) FetchEnterprise(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	allReport := make([]interface{}, 0)
	now := time.Now().Local()
	pool := make(chan bool, 10)
	for k, v := range stock.Stocks {
		pool <- true
		go func(key, val string) {
			s := stock.NewStock(key, val)
			curIndicator := map[string]interface{}{
				"fc":             *s.Code + *s.BourseCode,
				"corpType":       "4",
				"latestCount":    12,
				"reportDateType": 0,
			}

			reqBody, err := json.Marshal(curIndicator)
			if err != nil {
				log.Println(err.Error())
			}

			url := "https://emh5.eastmoney.com/api/CaiWuFenXi/GetZhuYaoZhiBiaoList"

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
			if err != nil {
				log.Println(err.Error())
			}

			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				log.Println(err.Error())
			}

			body, err := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			if err != nil {
				log.Println(err.Error())
			}

			var result models.MainIndicatorRes

			err = json.Unmarshal(body, &result)
			if err != nil {
				log.Println(err.Error())
			}

			for _, enterprise := range result.Result.Enterprise {
				enterprise.CreateDate = &now
				enterprise.Code = s.Code
				allReport = append(allReport, enterprise)
			}

			<-pool
		}(k, v)
	}

	tEnterpriseIndicator := d.GetColl(models.TEnterpriseIndicator)
	_, err := tEnterpriseIndicator.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	_, err = tEnterpriseIndicator.InsertMany(context.Background(), allReport)

	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	resultor.RetOk(w, &allReport)
}

// ListEnterprise 年报列表
func (d *DbEngine) ListEnterprise(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := d.GetColl(models.TEnterpriseIndicator)
	res, err := t.Find(context.Background(), bson.M{})

	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	a := make([]models.Enterprise, 0)

	err = res.All(context.Background(), &a)

	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	resultor.RetOk(w, &a)
}
