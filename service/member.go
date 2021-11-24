/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package service

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/zhimiaox/go-api-layout/errors"
	"github.com/zhimiaox/go-api-layout/models"
	"github.com/zhimiaox/go-api-layout/models/po"
	"github.com/zhimiaox/go-api-layout/models/vo"
	wechat "github.com/zhimiaox/go-api-layout/sdk"
	"github.com/zhimiaox/go-api-layout/source"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	"gorm.io/gorm"
)

var (
	memberInstance Member
	memberOnce     sync.Once
)

type Member interface {
	Login(req *vo.LoginReq) (*vo.LoginResp, errors.APIError)
	UpdateUserInfo(uid int, req *vo.UserInfoUpdateReq) errors.APIError
	Info(uid int) (*vo.SellerInfo, errors.APIError)
}

type memberImpl struct {
	db                *gorm.DB
	redis             *redis.Client
	repo              models.Member
	sellerMiniprogram *miniprogram.MiniProgram
	buyerMiniprogram  *miniprogram.MiniProgram
}

func (m *memberImpl) Info(uid int) (*vo.SellerInfo, errors.APIError) {
	info, err := m.repo.Get(m.db, uid)
	if err != nil {
		return nil, errors.NewAPIError(errors.Models, err.Error())
	}
	return &vo.SellerInfo{
		UID:           uid,
		AvatarURL:     info.AvatarURL,
		Nickname:      info.Nickname,
		SellerAgeYear: time.Now().Year() - info.CreatedAt.Year(),
	}, nil
}

func (m *memberImpl) Login(req *vo.LoginReq) (*vo.LoginResp, errors.APIError) {
	var (
		session auth.ResCode2Session
		err     error
	)
	switch req.AppID {
	case source.Config.Wechat.Seller.AppID:
		session, err = m.sellerMiniprogram.GetAuth().Code2Session(req.JsCode)
	case source.Config.Wechat.Buyer.AppID:
		session, err = m.buyerMiniprogram.GetAuth().Code2Session(req.JsCode)
	default:
		return nil, errors.NewAPIError(errors.Service, "非法的APPID")
	}
	if err != nil {
		return nil, errors.NewAPIError(errors.Service, "微信接口调用失败:{0}", err.Error())
	}
	uid, err := m.repo.GetUIDByOpenID(m.db, &po.Member{
		AppID:   req.AppID,
		OpenID:  session.OpenID,
		UnionID: session.UnionID,
	})
	if err != nil {
		return nil, errors.NewAPIError(errors.Models, err.Error())
	}
	hostname, _ := os.Hostname()
	now := time.Now()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  req.AppID,
		ExpiresAt: now.Add(time.Duration(source.Config.Auth.JwtExpire)).Unix(),
		Id:        fmt.Sprintf("%d", uid),
		IssuedAt:  now.Unix(),
		Issuer:    hostname,
		NotBefore: now.Unix(),
		Subject:   "",
	}).SignedString([]byte(source.Config.Auth.JwtSecret))
	if err != nil {
		return nil, errors.NewAPIError(errors.Service, err.Error())
	}
	return &vo.LoginResp{
		OpenID: session.OpenID,
		UID:    uid,
		Token:  token,
	}, nil
}

func (m *memberImpl) UpdateUserInfo(uid int, req *vo.UserInfoUpdateReq) errors.APIError {
	err := m.repo.UpdateUserInfo(m.db, uid, req.ToModel())
	if err != nil {
		return errors.NewAPIError(errors.Models, err.Error())
	}
	return nil
}

func GetMember() Member {
	memberOnce.Do(func() {
		memberInstance = &memberImpl{
			db:                source.GetMysql(),
			redis:             source.GetRedis(),
			sellerMiniprogram: wechat.NewSellerMiniProgramClient(),
			buyerMiniprogram:  wechat.NewBuyerMiniProgramClient(),
			repo:              models.GetMember(),
		}
	})
	return memberInstance
}
