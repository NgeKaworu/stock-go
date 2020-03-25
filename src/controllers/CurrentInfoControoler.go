package controllers

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"stock/src/models"
	"strings"
)

// FetchCurrentInfo 获取当前值
// func (d *DbEngine) FetchCurrentInfo(code, bourse string) {
func (d *DbEngine) FetchCurrentInfo() {

	// https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart2Get?fc=60000001&color=w

	resp, err := http.Get("http://hq.sinajs.cn/list=sh600000")
	if err != nil {
		panic(err)

	}
	s, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	// 股票名称、今日开盘价、昨日收盘价、当前价格、今日最高价、今日最低价、竞买价、竞卖价、成交股数、成交金额、买1手、买1报价、买2手、买2报价、…、买5报价、…、卖5报价、日期、时间
	strArr := strings.Split(string(s), ",")

	currentInfo := map[string]interface{}{}
	for k, v := range strArr[1 : len(strArr)-2] {
		currentInfo[models.CurrentInfoMap[k]] = v
	}
	log.Println(currentInfo)
	currentInfoMap := d.GetColl(models.TCurrentInfo)
	ret, err := currentInfoMap.InsertOne(context.Background(), currentInfo)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(ret)
}
