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
	ID          *primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`                   //id
	Code        string               `json:"code" bson:"code"`                                    //股票代码
	Bourse      *string              `json:"bourse,omitempty" bson:"bourse,omitempty"`            //交易所名字
	BourseCode  *string              `json:"bourseCode,omitempty" bson:"bourse_code,omitempty"`   //交易所代码
	Enterprise  *[]models.Enterprise `json:"enterprise,omitempty" bson:"enterprise,omitempty"`    //年报列表
	CurrentInfo *models.CurrentInfo  `json:"currentInfo,omitempty" bson:"current_info,omitempty"` //当前信息
	PB          float64              `json:"PB" bson:"pb"`                                        //市净率
	PE          float64              `json:"PE" bson:"pe"`                                        //市盈率
	PEG         float64              `json:"PEG" bson:"peg"`                                      //市盈增长比
	ROE         float64              `json:"ROE" bson:"roe"`                                      //净资产收益率
	DPE         float64              `json:"DPE" bson:"dpe"`                                      //动态利润估值
	DCE         float64              `json:"DCE" bson:"dce"`                                      //动态现金估值
	AAGR        float64              `json:"AAGR" bson:"aagr"`                                    //平均年增长率
	Grade       float64              `json:"grade" bson:"grade"`                                  //评分
	CreateDate  time.Time            `json:"createDate" bson:"create_date"`                       //创建时间
}

// TWeight 权重表名
const TWeight = "t_weight"

// Weights 权重结构
type Weights struct {
	Name       string    `json:"name" bson:"name"`                       //权重名字
	Weight     float64   `json:"weight" bson:"weight"`                   //权重
	Gt         bool      `json:"gt" bson:"gt"`                           //是否大于
	CreateDate time.Time `json:"createDate" bson:"create_date,omitzero"` //创建时间
}

// NewStock retrun new stock
func NewStock(code, bourseCode string) *Stock {
	s := &Stock{
		Code:       code,
		BourseCode: &bourseCode,
	}
	var bourse string
	switch bourseCode {
	case "01":
		bourse = "sh"
	case "02":
		bourse = "sz"
	default:
		break
	}

	s.Bourse = &bourse
	return s
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
func WeightSort(weights []Weights, s *[]Stock, total float64) {
	l := len(*s)
	for _, v := range weights {
		rate := v.Weight / total
		CusSort(*s, v.Name, v.Gt)
		for i := 0; i < l; i++ {
			(*s)[i].Grade += float64(l-i) * rate
		}
	}
	CusSort(*s, "Grade", true)
}
