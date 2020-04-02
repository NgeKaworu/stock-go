package stock

import (
	"reflect"
	"sort"
	"stock/src/models"
)

// Stock 股票基本结构
type Stock struct {
	Code        string               `bson:"code"`        //股票代码
	Bourse      string               `bson:"bourse"`      //交易所名字
	BourseCode  string               `bson:"bourse_code"` //交易所代码
	Enterprise  *[]models.Enterprise `bson:"-"`           //年报列表
	CurrentInfo *models.CurrentInfo  `bson:"-"`           //当前信息
	Classify    string               `bson:"classify"`    //板块
	PB          float64              `bson:"pb"`          //市净率
	PE          float64              `bson:"pe"`          //市盈率
	PEG         float64              `bson:"peg"`         //市盈增长比
	ROE         float64              `bson:"roe"`         //净资产收益率
	DPE         float64              `bson:"dpe"`         //动态利润估值
	DCE         float64              `bson:"dce"`         //动态现金估值
	AAGR        float64              `bson:"aagr"`        //平均年增长率
	Grade       float64              `bson:"grade"`       //评分
}

// CusSort 自定义 排序
func CusSort(s interface{}, key string, gt bool) {

	sort.Slice(s, func(i, j int) bool {
		var isGt bool
		val := reflect.ValueOf(s)
		s1 := val.Index(i).FieldByName(key).Interface().(float64)
		s2 := val.Index(j).FieldByName(key).Interface().(float64)
		if s1 > s2 {
			isGt = true
		}
		if gt {
			return isGt
		}

		return !isGt
	})
}

// WeightSort 权重排序
func WeightSort(weights map[string][]interface{}, s *[]Stock, total float64) {
	// CusSort(*s, "PB", true)
	l := len(*s)

	for k, v := range weights {
		weight, gt := v[0], v[1]
		rate := float64(weight.(int)) / total
		CusSort(*s, k, gt.(bool))
		for i := 0; i < l; i++ {
			(*s)[i].Grade += float64(l-i) * rate
		}
	}
	CusSort(*s, "Grade", true)
}
