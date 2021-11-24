/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package models

import (
	errors2 "errors"
	"sync"

	"github.com/zhimiaox/go-api-layout/models/po"
	"github.com/zhimiaox/go-api-layout/utils"

	"gorm.io/gorm"
)

var (
	memberInstance Member
	memberOnce     sync.Once
)

type Member interface {
	GetUIDByOpenID(db *gorm.DB, param *po.Member) (id int, err error)
	UpdateUserInfo(db *gorm.DB, id int, param *po.Member) error
	Get(db *gorm.DB, uid int) (*po.Member, error)
	GetNickname(db *gorm.DB, ids []int) (map[int]string, error)
	GetShortInfo(db *gorm.DB, ids []int) (map[int]po.Member, error)
}

type memberImpl struct{}

func (m *memberImpl) GetShortInfo(db *gorm.DB, ids []int) (map[int]po.Member, error) {
	ids = utils.IdsUniqueFitter(ids)
	data := make([]po.Member, 0)
	resp := make(map[int]po.Member)
	tx := db.Model(&po.Member{}).
		Select("id, nickname, avatar_url").
		Where("id IN (?)", ids).Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	for i := range data {
		resp[data[i].ID] = data[i]
	}
	return resp, nil
}

func (m *memberImpl) GetNickname(db *gorm.DB, ids []int) (map[int]string, error) {
	ids = utils.IdsUniqueFitter(ids)
	data := make([]po.Member, 0)
	resp := make(map[int]string)
	tx := db.Model(&po.Member{}).Select("id, nickname").Where("id IN (?)", ids).Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	for i := range data {
		resp[data[i].ID] = data[i].Nickname
	}
	return resp, nil
}

func (m *memberImpl) Get(db *gorm.DB, uid int) (*po.Member, error) {
	resp := &po.Member{}
	err := db.Where("id=?", uid).First(resp).Error
	return resp, err
}

func (m *memberImpl) GetUIDByOpenID(db *gorm.DB, param *po.Member) (id int, err error) {
	user := &po.Member{}
	tx := db.Model(user).Where(&po.Member{AppID: param.AppID, OpenID: param.OpenID}).Select("id").First(&user)
	if tx.Error != nil {
		if !errors2.Is(tx.Error, gorm.ErrRecordNotFound) {
			return 0, tx.Error
		}
		tx = db.Model(user).Create(&param)
		if tx.Error != nil {
			return 0, tx.Error
		}
		id = param.ID
	} else {
		id = user.ID
	}
	return id, nil
}

func (m *memberImpl) UpdateUserInfo(db *gorm.DB, id int, param *po.Member) error {
	param.ID = 0
	param.AppID = ""
	param.OpenID = ""
	param.UnionID = ""
	return db.Model(&param).Where(po.BaseModel{ID: id}).Updates(&param).Error
}

func GetMember() Member {
	memberOnce.Do(func() {
		memberInstance = &memberImpl{}
	})
	return memberInstance
}
