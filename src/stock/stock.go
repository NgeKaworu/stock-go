package stock

import (
	"github.com/NgeKaworu/stock/src/bitmask"
	"github.com/NgeKaworu/stock/src/models"
	"github.com/NgeKaworu/stock/src/utils"
)

// Stocks 全部股票集合
var Stocks = utils.Merge(Ss50, Hs300)

// Stock 股票基本结构
type Stock struct {
	Code                *string                       `json:"code" bson:"code"`                                  //股票代码
	Bourse              *string                       `json:"bourse,omitempty" bson:"bourse,omitempty"`          //交易所名字
	BourseCode          *string                       `json:"bourseCode,omitempty" bson:"bourse_code,omitempty"` //交易所代码
	*models.CurrentInfo `json:"currentinfo" bson:"-"` //当前信息
	Enterprise          []models.Enterprise           `json:"enterprise" bson:"-"`                          // 年报
	errorCode           bitmask.Bits                  `json:"-" bson:"errorCode"`                           // 错误码
	PB                  float64                       `json:"PB,omitempty" bson:"pb,omitempty"`             //市净率
	PE                  float64                       `json:"PE,omitempty" bson:"pe,omitempty"`             //市盈率
	PEG                 float64                       `json:"PEG,omitempty" bson:"peg,omitempty"`           //市盈增长比
	ROE                 float64                       `json:"ROE,omitempty" bson:"roe,omitempty"`           //净资产收益率
	DPE                 float64                       `json:"DPE,omitempty" bson:"dpe,omitempty"`           //动态利润估值
	DPER                float64                       `json:"DPER,omitempty" bson:"dper,omitempty"`         //动态利润估值率
	DCE                 float64                       `json:"DCE,omitempty" bson:"dce,omitempty"`           //动态现金估值
	DCER                float64                       `json:"DCER,omitempty" bson:"dcer,omitempty"`         //动态现金估值率
	AAGR                float64                       `json:"AAGR,omitempty" bson:"aagr,omitempty"`         //平均年增长率
	Grade               float64                       `json:"grade,omitempty" bson:"grade,omitempty"`       //评分
	Classify            *string                       `json:"classify,omitempty" bson:"classify,omitempty"` //板块
	Name                *string                       `json:"name,omitempty" bson:"name,omitempty"`         //股票名字
}

// NewStock retrun new stock
func NewStock(code, bourseCode string) *Stock {
	s := &Stock{
		Code:       &code,
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
