/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package dto

import (
	"time"
)

// TimeRange 时间区间筛选
type TimeRange struct {
	BeginTime time.Time `json:"begin_time" form:"begin_time"`
	EndTime   time.Time `json:"end_time" form:"end_time"`
}

type Search struct {
	SearchType int8   `json:"search_type" form:"search_type"`
	Keywords   string `json:"keywords" form:"keywords"` // 模糊搜索
}

type IDPage struct {
	Search
	LastID   int `json:"last_id" form:"last_id"`                // id分页首页0
	PageSize int `json:"page_size" form:"page_size" binding:""` // 每页大小
}

type Page struct {
	Search
	Page     int `json:"page" form:"page"`                      // 常规分页首页1
	PageSize int `json:"page_size" form:"page_size" binding:""` // 每页大小
}

func (p *Page) Offset() int {
	if p.Page < 2 { //nolint:gomnd
		return 0
	}
	return (p.Page - 1) * p.PageSize
}

type SellerOrderList struct {
	UID          int
	OrderState   []int8
	ShippingTime *TimeRange
	Page         Page
}

type BuyerOrderList struct {
	UID        int
	OrderState []int8
	Page       Page
}

type InventoryDo struct {
	GoodsID       int
	InventoryType int8
	ShippingTime  time.Time
	Number        int
}
