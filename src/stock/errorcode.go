/*
 * @Author: fuRan NgeKaworu@gmail.com
 * @Date: 2021-08-30 13:24:40
 * @LastEditors: fuRan NgeKaworu@gmail.com
 * @LastEditTime: 2023-01-30 15:25:28
 * @FilePath: /stock/stock-go/src/stock/errorcode.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package stock

import "github.com/NgeKaworu/stock/src/bitmask"

const (
	YEAR_ERR bitmask.Bits = 1 << iota
	CUR_ERR
	EXC_ERR
	POS_ERR
)
