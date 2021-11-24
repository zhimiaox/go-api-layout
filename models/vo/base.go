/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package vo

import (
	"net/http"

	"github.com/zhimiaox/go-api-layout/errors"
	"github.com/zhimiaox/go-api-layout/models/dto"

	"github.com/gin-gonic/gin"
)

const (
	UID           = "UID"
	APPID         = "APPID"
	Authorization = "Authorization"
)

// TimeRangeReq 时间区间筛选
type TimeRangeReq dto.TimeRange

func (p *TimeRangeReq) DTO() *dto.TimeRange {
	return (*dto.TimeRange)(p)
}

// IDPageReq ID分页
type IDPageReq dto.IDPage

func (p *IDPageReq) DTO() *dto.IDPage {
	return (*dto.IDPage)(p)
}

// PageReq page分页
type PageReq dto.Page

func (p *PageReq) DTO() *dto.Page {
	return (*dto.Page)(p)
}

// PageInfo 分页返回标准结构
type PageInfo struct {
	Page      int         `json:"page"` // id分页情况下返回0
	PageSize  int         `json:"page_size"`
	TotalSize int64       `json:"total_size"`
	Data      interface{} `json:"data"`
}

// ID 返回ID
type ID struct {
	ID int `json:"id"`
}

// APIErrorResult api默认错误返回结构
type APIErrorResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func APISuccess(c *gin.Context, e interface{}) {
	c.JSON(http.StatusOK, e)
}

func APIError(c *gin.Context, e ...interface{}) {
	httpStatus := http.StatusBadRequest
	result := &APIErrorResult{
		Code: 0,
		Msg:  "未知错误",
	}
	argLen := len(e)
	if argLen == 0 {
		c.JSON(httpStatus, result)
	}
	if argLen > 1 {
		if arg2, ok := e[1].(string); ok {
			result.Msg = arg2
		}
	}
	switch arg1 := e[0].(type) {
	case int:
		result.Code = arg1
	case errors.APIErrorCode:
		result.Code = int(arg1)
	case string:
		result.Msg = arg1
	case errors.APIError:
		result.Code = int(arg1.GetCode())
		result.Msg = arg1.Error()
	case error:
		result.Code = 0
		result.Msg = arg1.Error()
	}
	if s, ok := errors.APIErrorHTTPCodeRelation[errors.APIErrorCode(result.Code)]; ok {
		httpStatus = s
	}
	c.JSON(httpStatus, result)
}
