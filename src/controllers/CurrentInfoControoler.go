package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// FetchCurrentInfo 获取当前值
// func (d *DbEngine) FetchCurrentInfo(code, bourse string) {
func (d *DbEngine) FetchCurrentInfo() {
	resp, err := http.Get("http://hq.sinajs.cn/list=sh600000")
	if err != nil {
		panic(err)

	}
	s, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	// 股票名称、今日开盘价、昨日收盘价、当前价格、今日最高价、今日最低价、竞买价、竞卖价、成交股数、成交金额、买1手、买1报价、买2手、买2报价、…、买5报价、…、卖5报价、日期、时间
	strArr := strings.Split(string(s), ",")
	for _, v := range strArr[1 : len(strArr)-2] {
		log.Println(string(v))
	}
}
