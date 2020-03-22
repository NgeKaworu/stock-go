package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TEnterpriseIndicator 表名
const TEnterpriseIndicator = "t-enterprise-indicator"

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
	ID                           primitive.ObjectID `json:"id,omitempty" bson:"omitempty,_id"`
	CreateDate                   time.Time          `json:"CreateDate,omitempty" bson:"omitempty,create-date"`
	ReportDate                   string             `json:"ReportDate,omitempty" bson:"omitempty,report-date"`                                    //报告日期
	Title                        string             `json:"Title,omitempty" bson:"omitempty,title"`                                               //报告名称
	Epsjb                        string             `json:"Epsjb,omitempty" bson:"omitempty,epsjb"`                                               //基本每股收益(元)
	Epskcjb                      string             `json:"Epskcjb,omitempty" bson:"omitempty,epskcjb"`                                           //扣非每股收益(元)
	Epsxs                        string             `json:"Epsxs,omitempty" bson:"omitempty,epsxs"`                                               //稀释每股收益(元)
	Bps                          string             `json:"Bps,omitempty" bson:"omitempty,bps"`                                                   //每股净资产(元)
	Mgzbgj                       string             `json:"Mgzbgj,omitempty" bson:"omitempty,mgzbgj"`                                             //每股资本公积(元)
	Mgwfplr                      string             `json:"Mgwfplr,omitempty" bson:"omitempty,mgwfplr"`                                           //每股未分配利润(元)
	Mgjyxjje                     string             `json:"Mgjyxjje,omitempty" bson:"omitempty,mgjyxjje"`                                         //每股经营现金流(元)
	Totalincome                  string             `json:"Totalincome,omitempty" bson:"omitempty,totalincome"`                                   //营业总收入(元)
	Grossprofit                  string             `json:"Grossprofit,omitempty" bson:"omitempty,grossprofit"`                                   //毛利润(元)
	Parentnetprofit              string             `json:"Parentnetprofit,omitempty" bson:"omitempty,parentnetprofit"`                           //归属净利润(元)
	Bucklenetprofit              string             `json:"Bucklenetprofit,omitempty" bson:"omitempty,bucklenetprofit"`                           //扣非净利润(元)
	Totalincomeyoy               string             `json:"Totalincomeyoy,omitempty" bson:"omitempty,totalincomeyoy"`                             //营业总收入同比增长
	Parentnetprofityoy           string             `json:"Parentnetprofityoy,omitempty" bson:"omitempty,parentnetprofityoy"`                     //归属净利润同比增长
	Bucklenetprofityoy           string             `json:"Bucklenetprofityoy,omitempty" bson:"omitempty,bucklenetprofityoy"`                     //扣非净利润同比增长
	Totalincomerelativeratio     string             `json:"Totalincomerelativeratio,omitempty" bson:"omitempty,totalincomerelativeratio"`         //营业总收入滚动环比增长
	Parentnetprofitrelativeratio string             `json:"Parentnetprofitrelativeratio,omitempty" bson:"omitempty,parentnetprofitrelativeratio"` //归属净利润滚动环比增长
	Bucklenetprofitrelativeratio string             `json:"Bucklenetprofitrelativeratio,omitempty" bson:"omitempty,bucklenetprofitrelativeratio"` //扣非净利润滚动环比增长
	Roejq                        string             `json:"Roejq,omitempty" bson:"omitempty,roejq"`                                               //净资产收益率(加权)
	Roekcjq                      string             `json:"Roekcjq,omitempty" bson:"omitempty,roekcjq"`                                           //净资产收益率(扣非/加权)
	Allcapitalearningsrate       string             `json:"Allcapitalearningsrate,omitempty" bson:"omitempty,allcapitalearningsrate"`             //总资产收益率(加权)
	Grossmargin                  string             `json:"Grossmargin,omitempty" bson:"omitempty,grossmargin"`                                   //毛利率
	Netinterest                  string             `json:"Netinterest,omitempty" bson:"omitempty,netinterest"`                                   //净利率
	Accountsrate                 string             `json:"Accountsrate,omitempty" bson:"omitempty,accountsrate"`                                 //预收账款/营业收入
	Salesrate                    string             `json:"Salesrate,omitempty" bson:"omitempty,salesrate"`                                       //销售净现金流/营业收入
	Operatingrate                string             `json:"Operatingrate,omitempty" bson:"omitempty,operatingrate"`                               //经营净现金流/营业收入
	Taxrate                      string             `json:"Taxrate,omitempty" bson:"omitempty,taxrate"`                                           //实际税率
	Liquidityratio               string             `json:"Liquidityratio,omitempty" bson:"omitempty,liquidityratio"`                             //流动比率
	Quickratio                   string             `json:"Quickratio,omitempty" bson:"omitempty,quickratio"`                                     //速动比率
	Cashflowratio                string             `json:"Cashflowratio,omitempty" bson:"omitempty,cashflowratio"`                               //现金流量比率
	Assetliabilityratio          string             `json:"Assetliabilityratio,omitempty" bson:"omitempty,assetliabilityratio"`                   //资产负债率
	Equitymultiplier             string             `json:"Equitymultiplier,omitempty" bson:"omitempty,equitymultiplier"`                         //权益乘数
	Equityratio                  string             `json:"Equityratio,omitempty" bson:"omitempty,equityratio"`                                   //产权比率
	Totalassetsdays              string             `json:"Totalassetsdays,omitempty" bson:"omitempty,totalassetsdays"`                           //总资产周转天数(天)
	Inventorydays                string             `json:"Inventorydays,omitempty" bson:"omitempty,inventorydays"`                               //存货周转天数(天)
	Accountsreceivabledays       string             `json:"Accountsreceivabledays,omitempty" bson:"omitempty,accountsreceivabledays"`             //应收账款周转天数(天)
	Totalassetrate               string             `json:"Totalassetrate,omitempty" bson:"omitempty,totalassetrate"`                             //总资产周转率(次)
	Inventoryrate                string             `json:"Inventoryrate,omitempty" bson:"omitempty,inventoryrate"`                               //存货周转率(次)
	Accountsreceiveablerate      string             `json:"Accountsreceiveablerate,omitempty" bson:"omitempty,accountsreceiveablerate"`           //应收账款周转率(次)
}

// type YinHang struct {

// }
// type QuanShang struct {

// }
// type BaoXian struct {

// }
