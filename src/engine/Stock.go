package engine

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/NgeKaworu/stock/src/resultor"
	"github.com/NgeKaworu/stock/src/stock"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *DbEngine) StockCrawlMany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	allStock := make([]interface{}, 0)
	pool := make(chan bool, 10)
	now := time.Now().Local()
	format, _ := time.Parse("2006-01-02 15:03:05", now.Format("2006-01-02 00:00:00"))

	for k, v := range stock.Stocks {
		pool <- true
		go func(key, val string) {
			s := stock.NewStock(key, val)
			s.CreateAt = &format
			s.Crawl()
			allStock = append(allStock, s)
			<-pool
		}(k, v)

	}

	t := d.GetColl(stock.TStock)
	_, err := t.DeleteMany(context.Background(), bson.M{
		"createAt": &format,
	})

	if err != nil {
		resultor.RetFail(w, err)
		return
	}

	res, err := t.InsertMany(context.Background(), allStock)
	if err != nil {
		resultor.RetFail(w, err)
		return
	}

	resultor.RetOk(w, &res)
}

func (d *DbEngine) StockList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dataTime := r.URL.Query().Get("dataTime")
	times := make([]time.Time, 2)

	err := json.Unmarshal([]byte(dataTime), &times)
	if err != nil {
		resultor.RetFail(w, err)
		return
	}

	if len(times) != 2 {
		resultor.RetFail(w, errors.New("dataTime must a range"))
		return
	}

	query := bson.M{
		"createAt": bson.M{
			"$gte": times[0],
			"$lte": times[1],
		},
	}

	t := d.GetColl(stock.TStock)

	c, err := t.Find(context.Background(), &query)
	if err != nil {
		resultor.RetFail(w, err)
		return
	}

	res := make([]*stock.Stock, 0)
	err = c.All(context.Background(), &res)

	if err != nil {
		resultor.RetFail(w, err)
		return
	}

	resultor.RetOkWithTotal(w, res, int64(len(res)))
}
