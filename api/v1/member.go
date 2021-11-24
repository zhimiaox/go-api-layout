/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package v1

import (
	middleware2 "github.com/zhimiaox/go-api-layout/middleware"
	"github.com/zhimiaox/go-api-layout/models/vo"
	"github.com/zhimiaox/go-api-layout/service"

	"github.com/gin-gonic/gin"
)

type member struct {
	srv service.Member
}

func RegisterMember(r *gin.RouterGroup) {
	m := &member{
		srv: service.GetMember(),
	}
	group := r.Group("/member")
	group.POST("/login", m.Login)
	loginGroup := group.Use(middleware2.LoginAuthMiddleware())
	loginGroup.PUT("/userinfo", m.UpdateUserInfo)
}

// Login
// @Summary 登录
// @Tags 账户
// @Param default body vo.LoginReq false "入参"
// @Success 200 {object}  vo.LoginResp
// @Router /v1/member/login [post]
func (m *member) Login(c *gin.Context) {
	req := &vo.LoginReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		vo.APIError(c, err)
		return
	}
	loginInfo, apiError := m.srv.Login(req)
	if apiError != nil {
		vo.APIError(c, apiError)
		return
	}
	c.Header(vo.Authorization, "Bearer "+loginInfo.Token)
	vo.APISuccess(c, loginInfo)
}

// UpdateUserInfo
// @Summary 更新用户信息
// @Description 更新用户信息
// @Tags 账户
// @Param default body vo.UserInfoUpdateReq false "入参"
// @Success 200
// @Security ApiAuth
// @Router /v1/member/userinfo [put]
func (m *member) UpdateUserInfo(c *gin.Context) {
	req := &vo.UserInfoUpdateReq{}
	if err := c.ShouldBind(req); err != nil {
		vo.APIError(c, err)
		return
	}
	err := m.srv.UpdateUserInfo(c.GetInt(vo.UID), req)
	if err != nil {
		vo.APIError(c, err)
		return
	}
}
