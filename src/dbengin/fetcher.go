package dbengin

import (
	"context"
	"log"
	"stock/src/constants"
	"stock/src/models"
	"stock/src/stock"
	"stock/src/utils"
	"time"
)

// FetchEnterprise 爬年报+写库
func (d *DbEngine) FetchEnterprise() (string, error) {
	stocks := utils.Merge(constants.Ss50, constants.Hs300)
	allReport := make([]interface{}, 0)
	now := time.Now().Local()
	pool := make(chan bool, 10)
	for k, v := range stocks {
		pool <- true
		go func(key, val string) {
			s := &stock.Stock{
				Code:       key,
				BourseCode: val,
			}
			switch val {
			case "01":
				s.Bourse = "sh"
			case "02":
				s.Bourse = "sz"
			default:
				break
			}
			log.Println("current code: " + key)
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
	if _, err := tEnterpriseIndicator.InsertMany(context.Background(), allReport); err != nil {
		return "成功", nil
	} else {
		return "错误", err
	}

}
