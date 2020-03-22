package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"stock/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

// FetchMainIndicator 获取主要指标
func (d *DbEngine) FetchMainIndicator() {
	curIndicator := &models.MainIndicatorReq{Fc: "60001901", CorpType: "4", LatestCount: 5, ReportDateType: 0}

	reqBody, err := json.Marshal(curIndicator)

	url := "https://emh5.eastmoney.com/api/CaiWuFenXi/GetZhuYaoZhiBiaoList"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var result models.MainIndicatorRes

	err = json.Unmarshal(body, &result)

	enterpriseList := result.Result.Enterprise

	enterprise := d.GetColl(models.TEnterpriseIndicator)

	bson.Marshal(enterpriseList)

	ret, err := enterprise.InsertMany(context.Background())

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(ret)

}
