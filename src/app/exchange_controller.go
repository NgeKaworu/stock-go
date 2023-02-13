/*
 * @Author: fuRan NgeKaworu@gmail.com
 * @Date: 2023-01-30 18:05:33
 * @LastEditors: fuRan NgeKaworu@gmail.com
 * @LastEditTime: 2023-02-13 20:59:02
 * @FilePath: /stock/stock-go/src/app/exchange_controller.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package app

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/NgeKaworu/stock/src/bitmask"
	"github.com/NgeKaworu/stock/src/model"
	"github.com/NgeKaworu/stock/src/util"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *App) ExchangeList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (app *App) ExchangeUpsert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	tPostion := app.mongo.GetColl(model.TPosition)
	tPostion.FindOne(context.Background(), bson.M{"_id": "123"})
	// tExchange := app.mongo.GetColl(exchange.TExchange)

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		util.RetFail(w, err)
		return
	}

	if len(body) == 0 {
		util.RetFail(w, errors.New("not has body"))
		return
	}

	record := struct {
		ID               *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		BourseCode       *string             `json:"bourseCode,omitempty" bson:"bourse_code,omitempty"`             // 交易所代码
		errorCode        bitmask.Bits        `json:"-" bson:"errorCode"`                                            // 错误码
		CreateAt         *time.Time          `json:"createAt" bson:"createAt"`                                      // 创建时间
		TransactionPrice *float64            `json:"transactionPrice,omitempty" bson:"transaction_price,omitempty"` // 成交价
		CurrentShare     *float64            `json:"currentShare,omitempty" bson:"current_share,omitempty"`         // 本次股份
		CurrentDividend  *float64            `json:"currentDividend,omitempty" bson:"current_dividend,omitempty"`   // 本次派息
	}{}

	err = json.Unmarshal(body, &record)
	if err != nil {
		util.RetFail(w, err)
		return
	}

}

func (app *App) ExchangeDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
