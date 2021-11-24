/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package consts

// 订单状态
const (
	OrderStatePay              int8 = 10  // 待支付
	OrderStateConfirmed        int8 = 20  // 待确认(已支付)
	OrderStateShipping         int8 = 30  // 待配送(已确认)
	OrderStateShippingComplete int8 = 40  // 待完成(已送达) | 确认收货
	OrderStateComplete         int8 = 50  // 完成(已确认收货)
	OrderStateCancel           int8 = -10 // 用户取消
	OrderStateRefused          int8 = -20 // 商家拒绝
)

// 库存类型
const (
	InventoryTypeStandard int8 = iota + 1 // 标准库存类型
	InventoryTypeDay                      // 按天库存类型
)

// 商品状态
const (
	GoodsStateUp   int8 = iota + 1 // 上架
	GoodsStateDown                 // 下架
)

// 删除标识
const (
	DeleteFlagBuyer  = 1
	DeleteFlagSeller = 2
)

// 订单列表分组
const (
	OrderGroup1 int8 = 1
	OrderGroup2 int8 = 2
	OrderGroup3 int8 = 3
)
