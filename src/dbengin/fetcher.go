package dbengin

import (
	"context"
	"log"
	"stock/src/constants"
	"stock/src/models"
	"stock/src/stock"
	"stock/src/utils"
	"time"

	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var stocks = utils.Merge(constants.Ss50, constants.Hs300)

// FetchEnterprise 爬年报+写库
func (d *DbEngine) FetchEnterprise() (string, error) {
	allReport := make([]interface{}, 0)
	now := time.Now().Local()
	pool := make(chan bool, 10)
	for k, v := range stocks {
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
		return "删除错误", err
	}

	_, err = tEnterpriseIndicator.InsertMany(context.Background(), allReport)

	if err != nil {
		return "错误", err
	}

	return "成功", nil

}

// FetchCurrent 爬取当前信息
func (d *DbEngine) FetchCurrent() (string, error) {
	allMarket := make([]interface{}, 0)
	now := time.Now().Local()
	pool := make(chan bool, 10)
	for k, v := range stocks {
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
		return "FetchCurrent 失败", nil
	}

	return "成功", nil
}

// FetchInfoTime 获取所有 爬取时间
func (d *DbEngine) FetchInfoTime(ctx context.Context) ([]*graphql.Time, error) {
	query := []bson.M{
		{"$group": bson.M{
			"_id": "$create_date",
		}},
	}
	tCurrentInfo := d.GetColl(models.TCurrentInfo)
	re, err := tCurrentInfo.Aggregate(ctx, query, options.Aggregate())
	if err != nil {
		return nil, err
	}
	times := make([]map[string]time.Time, 0)
	err = re.All(ctx, &times)

	if err != nil {
		return nil, err
	}

	gqlTimes := make([]*graphql.Time, 0)

	for _, v := range times {
		if time, ok := v["_id"]; ok {
			gqlTime := &graphql.Time{Time: time.Local()}
			gqlTimes = append(gqlTimes, gqlTime)
		}

	}
	return gqlTimes, nil
}

// DiscountQuery 估值入参结构
type DiscountQuery struct {
	DiscountRate float64         `json:"discountRate" bson:"discount_rate"`
	CreateDate   string          `json:"createDate" bson:"create_date" formatter:"local"`
	Weights      []stock.Weights `json:"weights,omitempty" bson:"weights,omitempty"`
}

// Discount 计算估值
func (d *DbEngine) Discount(ctx context.Context, args DiscountQuery) (string, error) {
	// // 风险收益率(Rate of Risked Return)
	// // 假设10年内 > 80% 30年内 < 20%
	// RRR := 0.086
	// // 通货
	// CPI := 0.052
	// // 无风险利率 (The risk-free rate of interest)
	// RFR := 0.0285
	// discount := RRR + CPI + RFR
	m, err := d.Mapper.Conver2Map(args)

	if err != nil {
		return "Conver2Map 失败", err
	}

	tInfo := d.GetColl(models.TCurrentInfo)

	query := []bson.M{
		{"$match": bson.M{"create_date": m["create_date"]}},
		{"$project": bson.M{"_id": 0, "current_info": "$$ROOT"}},
		{
			"$lookup": bson.M{
				"from":         "t_enterprise_indicator",
				"localField":   "current_info.code",
				"foreignField": "code",
				"as":           "enterprise",
			},
		},
	}

	re, err := tInfo.Aggregate(ctx, query, options.Aggregate())
	if err != nil {
		return "Aggregate 失败", err
	}

	stocks := make([]stock.Stock, 0)
	err = re.All(ctx, &stocks)
	if err != nil {
		return "All 失败", err
	}

	stock.WeightSort(args.Weights, &stocks)

	now := time.Now().Local()

	s, err := d.Mapper.Conver(stocks)

	if err != nil {
		return "s Conver2Map 失败", err
	}

	for _, v := range s.([]interface{}) {
		v.(map[string]interface{})["create_date"] = now
		delete(v.(map[string]interface{}), "enterprise")
		delete(v.(map[string]interface{}), "current_info")
	}

	for _, v := range m["weights"].([]interface{}) {
		v.(map[string]interface{})["create_date"] = now
	}

	tStock := d.GetColl(stock.TStock)
	if _, err := tStock.InsertMany(context.Background(), s.([]interface{})); err != nil {
		return "InsertMany s Conver2Map 失败", err
	}

	tWeight := d.GetColl(stock.TWeight)
	if _, err := tWeight.InsertMany(context.Background(), m["weights"].([]interface{})); err != nil {
		return "InsertMany w Conver2Map 失败", err
	}

	return "成功", nil
}
