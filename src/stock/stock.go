package stock

import (
	"github.com/NgeKaworu/stock/src/utils"
)

// Stocks 全部股票集合
var Stocks = utils.Merge(Ss50, Hs300)

// Stock 股票基本结构
type Stock struct {
	Code       *string `json:"code" bson:"code"`                                  //股票代码
	Bourse     *string `json:"bourse,omitempty" bson:"bourse,omitempty"`          //交易所名字
	BourseCode *string `json:"bourseCode,omitempty" bson:"bourse_code,omitempty"` //交易所代码
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
