package models

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
	Enterprise Enterprise `json:"ZhuYaoZhiBiaoList_QiYe"`
	// YinHang   YinHang   `json:"ZhuYaoZhiBiaoList_YinHang"`
	// QuanShang QuanShang `json:"ZhuYaoZhiBiaoList_QuanShang"`
	// BaoXian   BaoXian   `json:"ZhuYaoZhiBiaoList_BaoXian"`
}

// Enterprise 企业指标
type Enterprise struct {
	ReportDate                   string `json:"ReportDate" bson:"report-date"`                                    //报告日期
	Title                        string `json:"Title" bson:"title"`                                               //报告名称
	Epsjb                        string `json:"Epsjb" bson:"epsjb"`                                               //基本每股收益(元)
	Epskcjb                      string `json:"Epskcjb" bson:"epskcjb"`                                           //扣非每股收益(元)
	Epsxs                        string `json:"Epsxs" bson:"epsxs"`                                               //稀释每股收益(元)
	Bps                          string `json:"Bps" bson:"bps"`                                                   //每股净资产(元)
	Mgzbgj                       string `json:"Mgzbgj" bson:"mgzbgj"`                                             //每股资本公积(元)
	Mgwfplr                      string `json:"Mgwfplr" bson:"mgwfplr"`                                           //每股未分配利润(元)
	Mgjyxjje                     string `json:"Mgjyxjje" bson:"mgjyxjje"`                                         //每股经营现金流(元)
	Totalincome                  string `json:"Totalincome" bson:"totalincome"`                                   //营业总收入(元)
	Grossprofit                  string `json:"Grossprofit" bson:"grossprofit"`                                   //毛利润(元)
	Parentnetprofit              string `json:"Parentnetprofit" bson:"parentnetprofit"`                           //归属净利润(元)
	Bucklenetprofit              string `json:"Bucklenetprofit" bson:"bucklenetprofit"`                           //扣非净利润(元)
	Totalincomeyoy               string `json:"Totalincomeyoy" bson:"totalincomeyoy"`                             //营业总收入同比增长
	Parentnetprofityoy           string `json:"Parentnetprofityoy" bson:"parentnetprofityoy"`                     //归属净利润同比增长
	Bucklenetprofityoy           string `json:"Bucklenetprofityoy" bson:"bucklenetprofityoy"`                     //扣非净利润同比增长
	Totalincomerelativeratio     string `json:"Totalincomerelativeratio" bson:"totalincomerelativeratio"`         //营业总收入滚动环比增长
	Parentnetprofitrelativeratio string `json:"Parentnetprofitrelativeratio" bson:"parentnetprofitrelativeratio"` //归属净利润滚动环比增长
	Bucklenetprofitrelativeratio string `json:"Bucklenetprofitrelativeratio" bson:"bucklenetprofitrelativeratio"` //扣非净利润滚动环比增长
	Roejq                        string `json:"Roejq" bson:"roejq"`                                               //净资产收益率(加权)
	Roekcjq                      string `json:"Roekcjq" bson:"roekcjq"`                                           //净资产收益率(扣非/加权)
	Allcapitalearningsrate       string `json:"Allcapitalearningsrate" bson:"allcapitalearningsrate"`             //总资产收益率(加权)
	Grossmargin                  string `json:"Grossmargin" bson:"grossmargin"`                                   //毛利率
	Netinterest                  string `json:"Netinterest" bson:"netinterest"`                                   //净利率
	Accountsrate                 string `json:"Accountsrate" bson:"accountsrate"`                                 //预收账款/营业收入
	Salesrate                    string `json:"Salesrate" bson:"salesrate"`                                       //销售净现金流/营业收入
	Operatingrate                string `json:"Operatingrate" bson:"operatingrate"`                               //经营净现金流/营业收入
	Taxrate                      string `json:"Taxrate" bson:"taxrate"`                                           //实际税率
	Liquidityratio               string `json:"Liquidityratio" bson:"liquidityratio"`                             //流动比率
	Quickratio                   string `json:"Quickratio" bson:"quickratio"`                                     //速动比率
	Cashflowratio                string `json:"Cashflowratio" bson:"cashflowratio"`                               //现金流量比率
	Assetliabilityratio          string `json:"Assetliabilityratio" bson:"assetliabilityratio"`                   //资产负债率
	Equitymultiplier             string `json:"Equitymultiplier" bson:"equitymultiplier"`                         //权益乘数
	Equityratio                  string `json:"Equityratio" bson:"equityratio"`                                   //产权比率
	Totalassetsdays              string `json:"Totalassetsdays" bson:"totalassetsdays"`                           //总资产周转天数(天)
	Inventorydays                string `json:"Inventorydays" bson:"inventorydays"`                               //存货周转天数(天)
	Accountsreceivabledays       string `json:"Accountsreceivabledays" bson:"accountsreceivabledays"`             //应收账款周转天数(天)
	Totalassetrate               string `json:"Totalassetrate" bson:"totalassetrate"`                             //总资产周转率(次)
	Inventoryrate                string `json:"Inventoryrate" bson:"inventoryrate"`                               //存货周转率(次)
	Accountsreceiveablerate      string `json:"Accountsreceiveablerate" bson:"accountsreceiveablerate"`           //应收账款周转率(次)
}

// type YinHang struct {

// }
// type QuanShang struct {

// }
// type BaoXian struct {

// }
