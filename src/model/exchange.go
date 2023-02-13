/*
 * @Author: fuRan NgeKaworu@gmail.com
 * @Date: 2023-01-30 15:19:17
 * @LastEditors: fuRan NgeKaworu@gmail.com
 * @LastEditTime: 2023-02-13 20:50:48
 * @FilePath: /stock/stock-go/src/model/exchange.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package model

import (
	"time"

	"github.com/NgeKaworu/stock/src/bitmask"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var TExchange = "t_exchange"

// Exchange 交易记录
type Exchange struct {
	ID               *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	BourseCode       *string             `json:"bourseCode,omitempty" bson:"bourse_code,omitempty"`             // 交易所代码
	errorCode        bitmask.Bits        `json:"-" bson:"errorCode"`                                            // 错误码
	CreateAt         *time.Time          `json:"createAt" bson:"createAt"`                                      // 创建时间
	TransactionPrice *float64            `json:"transactionPrice,omitempty" bson:"transaction_price,omitempty"` // 成交价
	CurrentShare     *float64            `json:"currentShare,omitempty" bson:"current_share,omitempty"`         // 本次股份
	CurrentDividend  *float64            `json:"currentDividend,omitempty" bson:"current_dividend,omitempty"`   // 本次派息
}
