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

// Discount 计算估值
func (d *DbEngine) Discount(ctx context.Context, args struct {
	DiscountRate float64         `json:"discountRate" bson:"discount_rate"`
	CreateDate   string          `json:"createDate" bson:"create_date" formatter:"local"`
	Weights      []stock.Weights `json:"weights,omitempty" bson:"weights,omitempty"`
}) (string, error) {
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
		{"$match": bson.M{"create_date": m["create_date"], "code": "002450"}},
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

	stocks := make([]*stock.Stock, 0)
	err = re.All(ctx, &stocks)
	if err != nil {
		return "All 失败", err
	}

	now := time.Now().Local()
	pool := make(chan bool, 100)
	ss := make([]interface{}, 0)
	for _, s := range stocks {
		pool <- true
		go func(s *stock.Stock) {
			s.CreateDate = now
			s.OriginDate = m["create_date"].(time.Time)
			s.Code = s.CurrentInfo.Code
			s.Classify = s.CurrentInfo.Classify
			s.Name = s.CurrentInfo.Name
			s.Calc()
			s.Discount(m["discount_rate"].(float64))
			s.Enterprise = nil
			s.CurrentInfo = nil
			ss = append(ss, s)
			<-pool
		}(s)
	}

	stock.WeightSort(args.Weights, &stocks)

	for _, v := range m["weights"].([]interface{}) {
		v.(map[string]interface{})["create_date"] = now
	}

	tStock := d.GetColl(stock.TStock)
	if _, err := tStock.InsertMany(context.Background(), ss); err != nil {
		return "InsertMany s 失败", err
	}

	tWeight := d.GetColl(stock.TWeight)
	if _, err := tWeight.InsertMany(context.Background(), m["weights"].([]interface{})); err != nil {
		return "InsertMany w 失败", err
	}

	return "成功", nil
}

// FetchStockTime 获取所有 爬取时间
func (d *DbEngine) FetchStockTime(ctx context.Context) ([]*graphql.Time, error) {
	query := []bson.M{
		{"$group": bson.M{
			"_id": "$create_date",
		}},
	}
	tStock := d.GetColl(stock.TStock)
	re, err := tStock.Aggregate(ctx, query, options.Aggregate())
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

// GetStockByTime 通过时间查估值
func (d *DbEngine) GetStockByTime(ctx context.Context, args struct {
	CreateDate string `json:"createDate" bson:"create_date" formatter:"local"`
	Limit      *int32 `json:"limit,omitempty" bson:"limit,omitempty"`
}) ([]*stock.Stock, error) {
	m, err := d.Mapper.Conver2Map(args)

	if err != nil {
		return nil, err
	}

	query := []bson.M{
		{"$match": bson.M{"create_date": m["create_date"]}},
		{"$sort": bson.M{"grade": -1}},
	}

	if limit, ok := m["limit"]; ok {
		query = append(query, bson.M{
			"$limit": limit,
		})
	}

	tStock := d.GetColl(stock.TStock)
	re, err := tStock.Aggregate(ctx, query, options.Aggregate())
	if err != nil {
		return nil, err
	}
	s := make([]*stock.Stock, 0)

	err = re.All(ctx, &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
