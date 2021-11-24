/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:10
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhimiaox/go-api-layout/models/vo"
)

type hello struct{}

func RegisterHello(r *gin.RouterGroup) {
	m := &hello{}
	group := r.Group("/hello")
	group.Any("", m.Hi)
}

// Hi
// @Summary 测试页面
// @Tags default
// @Success 200
// @Router /v1/hello [get]
func (h hello) Hi(c *gin.Context) {
	vo.APISuccess(c, "(#^.^#)")
}
