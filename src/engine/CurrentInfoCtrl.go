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
			log.Println("FetchCurrent current code: " + key)
			s.FetchCurrentInfo()
			s.FetchClassify()
			s.CurrentInfo.Code = s.Code
			s.CurrentInfo.CreateDate = now
			allMarket = append(allMarket, *s.CurrentInfo)
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
