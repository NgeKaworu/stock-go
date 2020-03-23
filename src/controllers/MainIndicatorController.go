package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"stock/src/models"
	"time"
)

// FetchMainIndicator 获取主要指标
func (d *DbEngine) FetchMainIndicator(code, bourse string) {
	curIndicator := &models.MainIndicatorReq{Fc: code + bourse, CorpType: "4", LatestCount: 5, ReportDateType: 0}

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

	var enterpriseListTemp []interface{}

	for _, v := range enterpriseList {
		createDate := time.Now().Local()
		v.CreateDate = createDate
		v.Code = code
		enterpriseListTemp = append(enterpriseListTemp, v)
	}

	ret, err := enterprise.InsertMany(context.Background(), enterpriseListTemp)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(ret)

}
