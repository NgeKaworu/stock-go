/*
 * @Author: fuRan NgeKaworu@gmail.com
 * @Date: 2023-01-30 15:19:24
 * @LastEditors: fuRan NgeKaworu@gmail.com
 * @LastEditTime: 2023-02-04 18:23:21
 * @FilePath: /stock/stock-go/src/position/position.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */

package position

import (
	"time"

	"github.com/NgeKaworu/stock/src/bitmask"
	"github.com/NgeKaworu/stock/src/stock"
)

// Position 持仓
type Position struct {
	BourseCode    *string      `json:"bourseCode,omitempty" bson:"_id,omitempty"`               // 交易所代码
	errorCode     bitmask.Bits `json:"-" bson:"errorCode"`                                      // 错误码
	Stock         *stock.Stock `json:"stock,omitempty" bson:"stock,omitempty"`                  // 股票
	TotalShare    *float64     `json:"totalShare,omitempty" bson:"total_share,omitempty"`       // 总股份
	TotalCapital  *float64     `json:"totalCapital,omitempty" bson:"total_capital,omitempty"`   // 总投入
	TotalDividend *float64     `json:"totalDividend,omitempty" bson:"total_dividend,omitempty"` // 总派息
	StopProfit    *float64     `json:"stopProfit,omitempty" bson:"stop_profit,omitempty"`       // 止盈点
	StopLoss      *float64     `json:"stopLoss,omitempty" bson:"stop_loss,omitempty"`           // 止损点
	CreateAt      *time.Time   `json:"createAt" bson:"createAt"`                                // 创建时间
	UpdateAt      *time.Time   `json:"updateAt" bson:"updateAt"`                                // 更新时间
}
