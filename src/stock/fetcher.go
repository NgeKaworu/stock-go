package stock

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/NgeKaworu/stock/src/bitmask"
	"github.com/NgeKaworu/stock/src/models"
)

var BOURSE_CODE_MAP = map[string]string{
	"01": "01",
	"02": "02",
}

func (s *Stock) FetchCurrentInform() error {
	s.errorCode = bitmask.Toggle(s.errorCode, CUR_ERR)

	u, err := url.Parse("https://push2.eastmoney.com/api/qt/stock/get")
	if err != nil {
		return err
	}

	q := u.Query()
	q.Add("fields", "f43,f58")
	q.Add("secid", BOURSE_CODE_MAP[*s.BourseCode]+"."+*s.Code)

	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var r struct {
		Data *struct {
			CurrentPrice *float64 `json:"f43,omitempty"`
			Name         *string  `json:"f58,omitempty"`
		} `json:"data,omitempty"`
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	s.Name = r.Data.Name
	s.CurrentPrice = r.Data.CurrentPrice

	clsPar := *s.Code + *s.BourseCode
	clsRes, err := http.Get("https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart2Get?fc=" + clsPar)
	if err != nil {
		return err
	}

	body, err = ioutil.ReadAll(clsRes.Body)
	if err != nil {
		return err
	}

	defer clsRes.Body.Close()

	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		return err
	}

	if r, ok := result["Result"].(map[string]interface{}); ok {
		if tiCaiXiangQingList, ok := r["TiCaiXiangQingList"]; ok {
			for _, tiCaiXiangQing := range tiCaiXiangQingList.([]interface{}) {
				if keyWord, ok := tiCaiXiangQing.(map[string]interface{})["KeyWord"].(string); ok {
					s.Classify = &keyWord
					break
				}
			}
		}

	}

	s.errorCode = bitmask.Toggle(s.errorCode, CUR_ERR)
	return nil
}

func (s *Stock) FetchEnterPrise() error {
	curIndicator := map[string]interface{}{
		"fc":             *s.Code + *s.BourseCode,
		"corpType":       "4",
		"latestCount":    12,
		"reportDateType": 0,
	}
	s.errorCode = bitmask.Toggle(s.errorCode, YEAR_ERR)

	reqBody, err := json.Marshal(curIndicator)
	if err != nil {
		return err
	}

	url := "https://emh5.eastmoney.com/api/CaiWuFenXi/GetZhuYaoZhiBiaoList"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	var result models.MainIndicatorRes

	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	s.Enterprise = make([]models.Enterprise, 0)

	s.Enterprise = append(s.Enterprise, result.Result.Enterprise...)

	s.errorCode = bitmask.Toggle(s.errorCode, YEAR_ERR)
	return nil
}
