/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package vo

import (
	"github.com/zhimiaox/go-api-layout/models/po"
)

type LoginReq struct {
	AppID  string `json:"app_id"`
	JsCode string `json:"js_code"`
}

type LoginResp struct {
	OpenID string `json:"open_id"`
	UID    int    `json:"uid"`
	Token  string `json:"token"`
}

type UserInfoUpdateReq struct {
	Nickname  string `json:"nickname"`   // 用户昵称
	AvatarURL string `json:"avatar_url"` // 头像
	Phone     string `json:"phone"`      // 手机号
	Gender    int    `json:"gender"`     // 1-男性 2-女性
	Country   string `json:"country"`    // 国家
	Province  string `json:"province"`   // 省
	City      string `json:"city"`       // 市
}

func (r *UserInfoUpdateReq) ToModel() *po.Member {
	return &po.Member{
		Nickname:  r.Nickname,
		AvatarURL: r.AvatarURL,
		Phone:     r.Phone,
		Gender:    r.Gender,
		Country:   r.Country,
		Province:  r.Province,
		City:      r.City,
	}
}

type SellerInfo struct {
	UID           int    `json:"uid"`
	AvatarURL     string `json:"avatar_url"`      // 头像
	Nickname      string `json:"nickname"`        // 昵称
	FavoriteNum   int    `json:"favorite_num"`    // 收藏数量
	GoodsNum      int    `json:"goods_num"`       // 在售商品数量
	SellerAgeYear int    `json:"seller_age_year"` // 喵龄
}
