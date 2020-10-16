package engine

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/NgeKaworu/stock/src/models"
	"github.com/NgeKaworu/stock/src/resultor"
	"github.com/NgeKaworu/stock/src/stock"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

// FetchAnnals 爬年报 并返回
func (d *DbEngine) FetchAnnals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	allReport := make([]interface{}, 0)
	now := time.Now().Local()
	pool := make(chan bool, 10)
	for k, v := range stock.Stocks {
		pool <- true
		go func(key, val string) {
			s := stock.NewStock(key, val)
			log.Println("FetchEnterprise current code: " + key)
			s.FetchMainIndicator()

			for _, enterprise := range *s.Enterprise {
				enterprise.CreateDate = now
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

// ListAnnals 年报列表
func (d *DbEngine) ListAnnals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
