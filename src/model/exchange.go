/*
 * @Author: fuRan NgeKaworu@gmail.com
 * @Date: 2023-01-30 15:19:17
 * @LastEditors: fuRan NgeKaworu@gmail.com
 * @LastEditTime: 2023-02-19 20:26:47
 * @FilePath: /stock/stock-go/src/model/exchange.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var TExchange = "t_exchange"

// Exchange 交易记录
type Exchange struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Code             *string            `json:"code,omitempty" bson:"code,omitempty" validate:"required"`              // 交易所代码
	CreateAt         MyTime             `json:"createAt" bson:"createAt" validate:"required,datetime"`                 // 创建时间
	UpdateAt         MyTime             `json:"updateAt" bson:"updateAt" validate:"required,datetime"`                 // 创建时间
	TransactionPrice float64            `json:"transactionPrice" bson:"transaction_price" validate:"required,numeric"` // 成交价
	CurrentShare     float64            `json:"currentShare" bson:"current_share" validate:"required,numeric"`         // 本次股份
	CurrentDividend  float64            `json:"currentDividend" bson:"current_dividend" validate:"required,numeric"`   // 本次派息
}
