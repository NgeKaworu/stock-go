package stock

import (
	"reflect"
	"sort"
	"stock/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TStock 表名
const TStock = "t_stock"

// Stock 股票基本结构
type Stock struct {
	ID          *primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`                 //id
	Code        string               `json:"code,omitempty" bson:"code,omitempty"`              //股票代码
	Bourse      string               `json:"bourse,omitempty" bson:"bourse,omitempty"`          //交易所名字
	BourseCode  string               `json:"bourseCode,omitempty" bson:"bourse_code,omitempty"` //交易所代码
	Enterprise  *[]models.Enterprise `json:"enterprise,omitempty" bson:"-,omitempty"`           //年报列表
	CurrentInfo *models.CurrentInfo  `json:"currentInfo,omitempty" bson:"-,omitempty"`          //当前信息
	Classify    string               `json:"classify,omitempty" bson:"classify,omitempty"`      //板块
	PB          float64              `json:"PB,omitempty" bson:"pb,omitempty"`                  //市净率
	PE          float64              `json:"PE,omitempty" bson:"pe,omitempty"`                  //市盈率
	PEG         float64              `json:"PEG,omitempty" bson:"peg,omitempty"`                //市盈增长比
	ROE         float64              `json:"ROE,omitempty" bson:"roe,omitempty"`                //净资产收益率
	DPE         float64              `json:"DPE,omitempty" bson:"dpe,omitempty"`                //动态利润估值
	DCE         float64              `json:"DCE,omitempty" bson:"dce,omitempty"`                //动态现金估值
	AAGR        float64              `json:"AAGR,omitempty" bson:"aagr,omitempty"`              //平均年增长率
	Grade       float64              `json:"grade,omitempty" bson:"grade,omitempty"`            //评分
	CreateDate  time.Time            `json:"createDate,omitempty" bson:"create_date,omitempty"` //创建时间
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
