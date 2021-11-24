/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package po

import (
	"time"
)

type BaseModel struct {
	ID        int       `gorm:"column:id;primarykey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// Member 用户表
type Member struct {
	BaseModel
	AppID     string `gorm:"column:app_id;uniqueIndex:member_app_id_open_id_uindex;type:varchar(50);not null"`
	OpenID    string `gorm:"column:open_id;uniqueIndex:member_app_id_open_id_uindex;type:varchar(50);not null"`
	UnionID   string `gorm:"unique;column:union_id;type:varchar(50);not null"`
	Nickname  string `gorm:"column:nickname;type:varchar(50);not null"`    // 用户昵称
	AvatarURL string `gorm:"column:avatar_url;type:varchar(300);not null"` // 头像
	Phone     string `gorm:"column:phone;type:varchar(15);not null"`       // 手机号
	Gender    int    `gorm:"column:gender;type:int;not null"`              // 1-男性 2-女性
	Country   string `gorm:"column:country;type:varchar(20);not null"`     // 国家
	Province  string `gorm:"column:province;type:varchar(20);not null"`    // 省
	City      string `gorm:"column:city;type:varchar(20);not null"`        // 市
	State     int    `gorm:"column:state;type:int"`                        // 用户状态 -1锁定 1正常
}

func (Member) TableName() string {
	return "member"
}
