package stock

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"stock/src/models"

	"strings"
)

// FetchCurrentInfo 获取当前值
func (s *Stock) FetchCurrentInfo() {
	params := s.Bourse + s.Code
	resp, err := http.Get("http://hq.sinajs.cn/list=" + params)
	if err != nil {
		panic(err)

	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	// 股票名称、今日开盘价、昨日收盘价、当前价格、今日最高价、今日最低价、竞买价、竞卖价、成交股数、成交金额、买1手、买1报价、买2手、买2报价、…、买5报价、…、卖5报价、日期、时间
	strArr := strings.Split(string(body), ",")

	ci := &models.CurrentInfo{}

	st := reflect.ValueOf(ci).Elem()
	for k, v := range strArr[1 : len(strArr)-2] {
		st.Field(k).SetString(v)
	}

	s.CurrentInfo = ci
}

// FetchClassify 获取分类(板块)
func (s *Stock) FetchClassify() {
	params := s.Code + s.BourseCode
	resp, err := http.Get("https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart2Get?fc=" + params)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	s.Classify = result["Result"].(map[string]interface{})["TiCaiXiangQingList"].([]interface{})[0].(map[string]interface{})["KeyWord"].(string)

	if err != nil {
		log.Println(err)
	}
}
