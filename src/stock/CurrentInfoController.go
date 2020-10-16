package stock

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"strings"

	"github.com/NgeKaworu/stock/src/models"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// FetchCurrentInfo 获取当前值
func (s *Stock) FetchCurrentInfo() {
	params := *s.Bourse + s.Code
	resp, err := http.Get("http://hq.sinajs.cn/list=" + params)
	if err != nil {
		panic(err)

	}
	utf8Reader := transform.NewReader(resp.Body,
		simplifiedchinese.GBK.NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// 股票名称、今日开盘价、昨日收盘价、当前价格、今日最高价、今日最低价、竞买价、竞卖价、成交股数、成交金额、买1手、买1报价、买2手、买2报价、…、买5报价、…、卖5报价、日期、时间
	strArr := strings.Split(string(body), ",")

	ci := &models.CurrentInfo{}

	st := reflect.ValueOf(ci).Elem()
	for k, v := range strArr[:len(strArr)-3] {
		if k == 0 {
			st.Field(k).SetString(strings.Split(v, "\"")[1])
			continue
		}
		st.Field(k).SetString(v)

	}

	s.CurrentInfo = ci
}

// FetchClassify 获取分类(板块)
func (s *Stock) FetchClassify() {
	params := s.Code + *s.BourseCode
	resp, err := http.Get("https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart2Get?fc=" + params)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Println(err)
	}

	if r, ok := result["Result"].(map[string]interface{}); ok {
		if tiCaiXiangQingList, ok := r["TiCaiXiangQingList"]; ok {
			for _, tiCaiXiangQing := range tiCaiXiangQingList.([]interface{}) {
				if keyWord, ok := tiCaiXiangQing.(map[string]interface{})["KeyWord"].(string); ok {
					s.CurrentInfo.Classify = keyWord
				}
				break
			}
		}

	}

}
