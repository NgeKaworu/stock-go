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
	ID          primitive.ObjectID   `json:"id" bson:"_id"`                 //id
	Code        string               `json:"code" bson:"code"`              //股票代码
	Bourse      string               `json:"bourse" bson:"bourse"`          //交易所名字
	BourseCode  string               `json:"bourseCode" bson:"bourse_code"` //交易所代码
	Enterprise  *[]models.Enterprise `json:"enterprise" bson:"-"`           //年报列表
	CurrentInfo *models.CurrentInfo  `json:"currentInfo" bson:"-"`          //当前信息
	Classify    string               `json:"classify" bson:"classify"`      //板块
	PB          float64              `json:"PB" bson:"pb"`                  //市净率
	PE          float64              `json:"PE" bson:"pe"`                  //市盈率
	PEG         float64              `json:"PEG" bson:"peg"`                //市盈增长比
	ROE         float64              `json:"ROE" bson:"roe"`                //净资产收益率
	DPE         float64              `json:"DPE" bson:"dpe"`                //动态利润估值
	DCE         float64              `json:"DCE" bson:"dce"`                //动态现金估值
	AAGR        float64              `json:"AAGR" bson:"aagr"`              //平均年增长率
	Grade       float64              `json:"grade" bson:"grade"`            //评分
	CreateDate  time.Time            `json:"createDate" bson:"create_date"` //创建时间
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
func WeightSort(weights map[string][]interface{}, s *[]interface{}, total float64) {
	// CusSort(*s, "PB", true)
	l := len(*s)

	for k, v := range weights {
		weight, gt := v[0], v[1]
		rate := float64(weight.(int)) / total
		CusSort(*s, k, gt.(bool))
		for i := 0; i < l; i++ {
			(*s)[i].(*Stock).Grade += float64(l-i) * rate
		}
	}
	CusSort(*s, "Grade", true)
}
