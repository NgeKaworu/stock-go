package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TEnterpriseIndicator 表名
const TEnterpriseIndicator = "t_enterprise_indicator"

//MainIndicatorReq 主要指标请求体
type MainIndicatorReq struct {
	Fc             string `json:"fc"`
	CorpType       string `json:"corpType"`
	LatestCount    int    `json:"latestCount"`
	ReportDateType int    `json:"reportDateType"`
}

// MainIndicatorRes 主要指标返回体
type MainIndicatorRes struct {
	Result Result `json:"Result"`
}

// Result 主要指标返回值
type Result struct {
	Enterprise []Enterprise `json:"ZhuYaoZhiBiaoList_QiYe"`
	// YinHang   YinHang   `json:"ZhuYaoZhiBiaoList_YinHang"`
	// QuanShang QuanShang `json:"ZhuYaoZhiBiaoList_QuanShang"`
	// BaoXian   BaoXian   `json:"ZhuYaoZhiBiaoList_BaoXian"`
}

// Enterprise 企业指标
type Enterprise struct {
	ID                           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreateDate                   time.Time          `json:"CreateDate,omitempty" bson:"create-date,omitempty"`
	Code                         string             `json:"code,omitempty" bson:"code,omitempty"`                                                 //股票代号
	ReportDate                   string             `json:"ReportDate,omitempty" bson:"report-date,omitempty"`                                    //报告日期
	Title                        string             `json:"Title,omitempty" bson:"title,omitempty"`                                               //报告名称
	Epsjb                        string             `json:"Epsjb,omitempty" bson:"epsjb,omitempty"`                                               //基本每股收益(元)
	Epskcjb                      string             `json:"Epskcjb,omitempty" bson:"epskcjb,omitempty"`                                           //扣非每股收益(元)
	Epsxs                        string             `json:"Epsxs,omitempty" bson:"epsxs,omitempty"`                                               //稀释每股收益(元)
	Bps                          string             `json:"Bps,omitempty" bson:"bps,omitempty"`                                                   //每股净资产(元)
	Mgzbgj                       string             `json:"Mgzbgj,omitempty" bson:"mgzbgj,omitempty"`                                             //每股资本公积(元)
	Mgwfplr                      string             `json:"Mgwfplr,omitempty" bson:"mgwfplr,omitempty"`                                           //每股未分配利润(元)
	Mgjyxjje                     string             `json:"Mgjyxjje,omitempty" bson:"mgjyxjje,omitempty"`                                         //每股经营现金流(元)
	Totalincome                  string             `json:"Totalincome,omitempty" bson:"totalincome,omitempty"`                                   //营业总收入(元)
	Grossprofit                  string             `json:"Grossprofit,omitempty" bson:"grossprofit,omitempty"`                                   //毛利润(元)
	Parentnetprofit              string             `json:"Parentnetprofit,omitempty" bson:"parentnetprofit,omitempty"`                           //归属净利润(元)
	Bucklenetprofit              string             `json:"Bucklenetprofit,omitempty" bson:"bucklenetprofit,omitempty"`                           //扣非净利润(元)
	Totalincomeyoy               string             `json:"Totalincomeyoy,omitempty" bson:"totalincomeyoy,omitempty"`                             //营业总收入同比增长
	Parentnetprofityoy           string             `json:"Parentnetprofityoy,omitempty" bson:"parentnetprofityoy,omitempty"`                     //归属净利润同比增长
	Bucklenetprofityoy           string             `json:"Bucklenetprofityoy,omitempty" bson:"bucklenetprofityoy,omitempty"`                     //扣非净利润同比增长
	Totalincomerelativeratio     string             `json:"Totalincomerelativeratio,omitempty" bson:"totalincomerelativeratio,omitempty"`         //营业总收入滚动环比增长
	Parentnetprofitrelativeratio string             `json:"Parentnetprofitrelativeratio,omitempty" bson:"parentnetprofitrelativeratio,omitempty"` //归属净利润滚动环比增长
	Bucklenetprofitrelativeratio string             `json:"Bucklenetprofitrelativeratio,omitempty" bson:"bucklenetprofitrelativeratio,omitempty"` //扣非净利润滚动环比增长
	Roejq                        string             `json:"Roejq,omitempty" bson:"roejq,omitempty"`                                               //净资产收益率(加权)
	Roekcjq                      string             `json:"Roekcjq,omitempty" bson:"roekcjq,omitempty"`                                           //净资产收益率(扣非/加权)
	Allcapitalearningsrate       string             `json:"Allcapitalearningsrate,omitempty" bson:"allcapitalearningsrate,omitempty"`             //总资产收益率(加权)
	Grossmargin                  string             `json:"Grossmargin,omitempty" bson:"grossmargin,omitempty"`                                   //毛利率
	Netinterest                  string             `json:"Netinterest,omitempty" bson:"netinterest,omitempty"`                                   //净利率
	Accountsrate                 string             `json:"Accountsrate,omitempty" bson:"accountsrate,omitempty"`                                 //预收账款/营业收入
	Salesrate                    string             `json:"Salesrate,omitempty" bson:"salesrate,omitempty"`                                       //销售净现金流/营业收入
	Operatingrate                string             `json:"Operatingrate,omitempty" bson:"operatingrate,omitempty"`                               //经营净现金流/营业收入
	Taxrate                      string             `json:"Taxrate,omitempty" bson:"taxrate,omitempty"`                                           //实际税率
	Liquidityratio               string             `json:"Liquidityratio,omitempty" bson:"liquidityratio,omitempty"`                             //流动比率
	Quickratio                   string             `json:"Quickratio,omitempty" bson:"quickratio,omitempty"`                                     //速动比率
	Cashflowratio                string             `json:"Cashflowratio,omitempty" bson:"cashflowratio,omitempty"`                               //现金流量比率
	Assetliabilityratio          string             `json:"Assetliabilityratio,omitempty" bson:"assetliabilityratio,omitempty"`                   //资产负债率
	Equitymultiplier             string             `json:"Equitymultiplier,omitempty" bson:"equitymultiplier,omitempty"`                         //权益乘数
	Equityratio                  string             `json:"Equityratio,omitempty" bson:"equityratio,omitempty"`                                   //产权比率
	Totalassetsdays              string             `json:"Totalassetsdays,omitempty" bson:"totalassetsdays,omitempty"`                           //总资产周转天数(天)
	Inventorydays                string             `json:"Inventorydays,omitempty" bson:"inventorydays,omitempty"`                               //存货周转天数(天)
	Accountsreceivabledays       string             `json:"Accountsreceivabledays,omitempty" bson:"accountsreceivabledays,omitempty"`             //应收账款周转天数(天)
	Totalassetrate               string             `json:"Totalassetrate,omitempty" bson:"totalassetrate,omitempty"`                             //总资产周转率(次)
	Inventoryrate                string             `json:"Inventoryrate,omitempty" bson:"inventoryrate,omitempty"`                               //存货周转率(次)
	Accountsreceiveablerate      string             `json:"Accountsreceiveablerate,omitempty" bson:"accountsreceiveablerate,omitempty"`           //应收账款周转率(次)
}

// type YinHang struct {

// }
// type QuanShang struct {

// }
// type BaoXian struct {

// }
