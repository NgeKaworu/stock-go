package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"stock/src/models"
)

// FetchMainIndicator 获取主要指标
func FetchMainIndicator() {
	curIndicator := &models.MainIndicatorReq{"60001901", "4", 5, 0}

	reqBody, err := json.Marshal(curIndicator)

	url := "https://emh5.eastmoney.com/api/CaiWuFenXi/GetZhuYaoZhiBiaoList"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(string(body))
}
