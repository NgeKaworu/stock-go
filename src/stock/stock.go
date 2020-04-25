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
	PB          float64              `json:"PB,omitempty" bson:"pb,omitempty"`                    //市净率
	PE          float64              `json:"PE,omitempty" bson:"pe,omitempty"`                    //市盈率
	PEG         float64              `json:"PEG,omitempty" bson:"peg,omitempty"`                  //市盈增长比
	ROE         float64              `json:"ROE,omitempty" bson:"roe,omitempty"`                  //净资产收益率
	DPE         float64              `json:"DPE,omitempty" bson:"dpe,omitempty"`                  //动态利润估值
	DCE         float64              `json:"DCE,omitempty" bson:"dce,omitempty"`                  //动态现金估值
	AAGR        float64              `json:"AAGR,omitempty" bson:"aagr,omitempty"`                //平均年增长率
	Grade       float64              `json:"grade,omitempty" bson:"grade,omitempty"`              //评分
	CreateDate  time.Time            `json:"createDate" bson:"create_date"`                       //创建时间
	OriginDate  time.Time            `json:"originDate" bson:"origin_date"`                       //ci时间
	Classify    string               `json:"classify" bson:"classify"`                            //板块
	Name        string               `json:"name,omitempty" bson:"name,omitempty"`                //股票名字
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
		val := reflect.Indirect(reflect.ValueOf(s))
		s1 := reflect.Indirect(val.Index(i)).FieldByName(key).Float()
		s2 := reflect.Indirect(val.Index(j)).FieldByName(key).Float()
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
func WeightSort(weights []Weights, s *[]*Stock) {
	var total float64
	for _, v := range weights {
		total += v.Weight
	}

	l := len(*s)
	pool := make(chan bool, 100)
	for _, v := range weights {
		rate := v.Weight / total
		CusSort(*s, v.Name, v.Gt)
		for i := 0; i < l; i++ {
			pool <- true
			go func(i int) {

				(*s)[i].Grade += float64(l-i) * rate
				<-pool
			}(i)
		}

	}
	CusSort(*s, "Grade", true)
}
