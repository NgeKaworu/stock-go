package stock

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"stock/src/models"
)

// FetchMainIndicator 获取主要指标
func (s *Stock) FetchMainIndicator() {
	curIndicator := &models.MainIndicatorReq{Fc: s.Code + *s.BourseCode, CorpType: "4", LatestCount: 12, ReportDateType: 0}

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

	s.Enterprise = &result.Result.Enterprise

	if err != nil {
		log.Println(err.Error())
	}

}
