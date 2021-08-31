package stock

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/NgeKaworu/stock/src/bitmask"
	"github.com/NgeKaworu/stock/src/models"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func (s *Stock) FetchCurrentInfor() (*Stock, error) {
	ciPar := *s.Bourse + *s.Code
	s.errorCode = bitmask.Toggle(s.errorCode, CUR_ERR)

	ciRes, err := http.Get("http://hq.sinajs.cn/list=" + ciPar)
	if err != nil {
		return nil, err
	}
	// 中文编码
	utf8Reader := transform.NewReader(ciRes.Body, simplifiedchinese.GBK.NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		return nil, err
	}

	defer ciRes.Body.Close()
	// 股票名称、今日开盘价、昨日收盘价、当前价格、今日最高价、今日最低价、竞买价、竞卖价、成交股数、成交金额、买1手、买1报价、买2手、买2报价、…、买5报价、…、卖5报价、日期、时间
	strArr := strings.Split(string(body), ",")
	s.CurrentInfo = &models.CurrentInfo{}
	st := reflect.ValueOf(s.CurrentInfo).Elem()
	for k, v := range strArr[:len(strArr)-3] {
		if k == 0 {
			st.Field(k).Set(reflect.ValueOf(&strings.Split(v, "\"")[1]))
			continue
		}
		// 创建临时变量来接指针
		value := v
		st.Field(k).Set(reflect.ValueOf(&value))

	}

	clsPar := *s.Code + *s.BourseCode
	clsRes, err := http.Get("https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart2Get?fc=" + clsPar)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(clsRes.Body)
	if err != nil {
		return nil, err
	}

	defer clsRes.Body.Close()

	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	if r, ok := result["Result"].(map[string]interface{}); ok {
		if tiCaiXiangQingList, ok := r["TiCaiXiangQingList"]; ok {
			for _, tiCaiXiangQing := range tiCaiXiangQingList.([]interface{}) {
				if keyWord, ok := tiCaiXiangQing.(map[string]interface{})["KeyWord"].(string); ok {
					s.CurrentInfo.Classify = &keyWord
					break
				}
			}
		}

	}

	s.errorCode = bitmask.Toggle(s.errorCode, CUR_ERR)
	s.Name = s.CurrentInfo.Name
	s.Classify = s.CurrentInfo.Classify
	return s, nil
}

func (s *Stock) FetchEnterPrise() (*Stock, error) {
	curIndicator := map[string]interface{}{
		"fc":             *s.Code + *s.BourseCode,
		"corpType":       "4",
		"latestCount":    12,
		"reportDateType": 0,
	}
	s.errorCode = bitmask.Toggle(s.errorCode, YEAR_ERR)

	reqBody, err := json.Marshal(curIndicator)
	if err != nil {
		return nil, err
	}

	url := "https://emh5.eastmoney.com/api/CaiWuFenXi/GetZhuYaoZhiBiaoList"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var result models.MainIndicatorRes

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	s.Enterprise = make([]models.Enterprise, 0)

	s.Enterprise = append(s.Enterprise, result.Result.Enterprise...)

	s.errorCode = bitmask.Toggle(s.errorCode, YEAR_ERR)
	return s, nil
}
