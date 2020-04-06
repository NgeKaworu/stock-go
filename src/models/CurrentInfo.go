package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TCurrentInfo 表名
const TCurrentInfo = "t_current_info"

// CurrentInfo 当前股票信息
type CurrentInfo struct {
	TodayOpeningPrice     string             `json:"todayOpeningPrice,omitempty" bson:"today_opening_price,omitempty"`         //今日开盘价
	YesterdayOpeningPrice string             `json:"yesterdayOpeningPrice,omitempty" bson:"yesterday_opening_price,omitempty"` //昨日收盘价
	CurrentPrice          string             `json:"currentPrice,omitempty" bson:"current_price,omitempty"`                    //当前价格
	TopPrice              string             `json:"topPrice,omitempty" bson:"top_price,omitempty"`                            //今日最高价
	FloorPrice            string             `json:"floorPrice,omitempty" bson:"floor_price,omitempty"`                        //今日最低价
	BidPrice              string             `json:"bidPrice,omitempty" bson:"bid_price,omitempty"`                            //竞买价
	AuctionPrice          string             `json:"auctionPrice,omitempty" bson:"auction_price,omitempty"`                    //竞卖价
	Vol                   string             `json:"vol,omitempty" bson:"vol,omitempty"`                                       //成交股数
	Amount                string             `json:"amount,omitempty" bson:"amount,omitempty"`                                 //成交金额
	Buy1Num               string             `json:"buy1Num,omitempty" bson:"buy_1_num,omitempty"`                             //买1手
	Buy1Price             string             `json:"buy1Price,omitempty" bson:"buy_1_price,omitempty"`                         //买1报价
	Buy2Num               string             `json:"buy2Num,omitempty" bson:"buy_2_num,omitempty"`                             //买2手
	Buy2Price             string             `json:"buy2Price,omitempty" bson:"buy_2_price,omitempty"`                         //买2报价
	Buy3Num               string             `json:"buy3Num,omitempty" bson:"buy_3_num,omitempty"`                             //买3手
	Buy3Price             string             `json:"buy3Price,omitempty" bson:"buy_3_price,omitempty"`                         //买3报价
	Buy4Num               string             `json:"buy4Num,omitempty" bson:"buy_4_num,omitempty"`                             //买4手
	Buy4Price             string             `json:"buy4Price,omitempty" bson:"buy_4_price,omitempty"`                         //买4报价
	Buy5Num               string             `json:"buy5Num,omitempty" bson:"buy_5_num,omitempty"`                             //买5手
	Buy5Price             string             `json:"buy5Price,omitempty" bson:"buy_5_price,omitempty"`                         //买5报价
	Sell1Num              string             `json:"sell1Num,omitempty" bson:"sell_1_num,omitempty"`                           //卖1手
	Sell1Price            string             `json:"sell1Price,omitempty" bson:"sell_1_price,omitempty"`                       //卖1报价
	Sell2Num              string             `json:"sell2Num,omitempty" bson:"sell_2_num,omitempty"`                           //卖2手
	Sell2Price            string             `json:"sell2Price,omitempty" bson:"sell_2_price,omitempty"`                       //卖2报价
	Sell3Num              string             `json:"sell3Num,omitempty" bson:"sell_3_num,omitempty"`                           //卖3手
	Sell3Price            string             `json:"sell3Price,omitempty" bson:"sell_3_price,omitempty"`                       //卖3报价
	Sell4Num              string             `json:"sell4Num,omitempty" bson:"sell_4_num,omitempty"`                           //卖4手
	Sell4Price            string             `json:"sell4Price,omitempty" bson:"sell_4_price,omitempty"`                       //卖4报价
	Sell5Num              string             `json:"sell5Num,omitempty" bson:"sell_5_num,omitempty"`                           //卖5手
	Sell5Price            string             `json:"sell5Price,omitempty" bson:"sell_5_price,omitempty"`                       //卖5报价
	Date                  string             `json:"date,omitempty" bson:"date,omitempty"`                                     //日期
	Time                  string             `json:"time,omitempty" bson:"time,omitempty"`                                     //时间
	ID                    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`                                        //id
	CreateDate            time.Time          `json:"createDate" bson:"create_date"`                                            //创建时间
}
