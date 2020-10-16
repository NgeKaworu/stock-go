package stock

import (
	"time"

	"github.com/NgeKaworu/stock/src/constants"
	"github.com/NgeKaworu/stock/src/models"
	"github.com/NgeKaworu/stock/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TStock 表名
const TStock = "t_stock"

// Stocks 全部股票集合
var Stocks = utils.Merge(constants.Ss50, constants.Hs300)

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
	DPER        float64              `json:"DPER,omitempty" bson:"dper,omitempty"`                //动态利润估值率
	DCE         float64              `json:"DCE,omitempty" bson:"dce,omitempty"`                  //动态现金估值
	DCER        float64              `json:"DCER,omitempty" bson:"dcer,omitempty"`                //动态现金估值率
	AAGR        float64              `json:"AAGR,omitempty" bson:"aagr,omitempty"`                //平均年增长率
	Grade       float64              `json:"grade,omitempty" bson:"grade,omitempty"`              //评分
	CreateDate  time.Time            `json:"createDate" bson:"create_date"`                       //创建时间
	OriginDate  time.Time            `json:"originDate" bson:"origin_date"`                       //ci时间
	Classify    string               `json:"classify" bson:"classify"`                            //板块
	Name        string               `json:"name,omitempty" bson:"name,omitempty"`                //股票名字
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
