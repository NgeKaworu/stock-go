package stock

import "stock/src/models"

// Stock 股票基本结构
type Stock struct {
	Code        string               //股票代码
	Bourse      string               //交易所名字
	BourseCode  string               //交易所代码
	Enterprise  *[]models.Enterprise //年报列表
	CurrentInfo *models.CurrentInfo  //当前信息
	Classify    string               //板块
	PB          float64              //市净率
	PE          float64              //市盈率
	PEG         float64              //市盈增长比
	ROE         float64              //净资产收益率
	DPE         float64              //动态利润估值
	DCE         float64              //动态现金估值
	AAGR        float64              //平均年增长率
	Grade       int                  //评分
}
