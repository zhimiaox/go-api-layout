/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	// MissingParameter 参数缺失
	MissingParameter APIErrorCode = iota + 1
	// MaximumLimit 最大限制
	MaximumLimit
	// MinimumLimit 最小限制
	MinimumLimit
	// InvalidArguments 参数格式错误
	InvalidArguments
	// Service 服务错误
	Service
	// Models 模型
	Models
	// InternalServerError 服务器异常
	InternalServerError
	// Unauthorized 未授权
	Unauthorized
	// Forbidden 拒绝访问
	Forbidden
)

// APIErrorHTTPCodeRelation 自定义HTTP状态码
var APIErrorHTTPCodeRelation = map[APIErrorCode]int{
	MissingParameter:    http.StatusBadRequest,
	InternalServerError: http.StatusInternalServerError,
	Unauthorized:        http.StatusUnauthorized,
	Forbidden:           http.StatusForbidden,
}

type APIErrorCode int

type APIError interface {
	Error() string
	GetCode() APIErrorCode
}

type apiError struct {
	Msg         string
	Code        APIErrorCode
	MsgTemplate string
	Args        []interface{}
}

func (e *apiError) Error() string {
	return e.Msg
}

func (e apiError) GetCode() APIErrorCode {
	return e.Code
}

func NewAPIError(code APIErrorCode, msg string, args ...interface{}) APIError {
	e := &apiError{
		Code: code,
		Msg:  msg,
		Args: args,
	}
	if len(args) > 0 {
		e.MsgTemplate = msg
		for i, v := range args {
			e.Msg = strings.ReplaceAll(e.Msg, "{"+strconv.Itoa(i)+"}", fmt.Sprint(v))
		}
	}
	return e
}

func New(text string) error {
	return errors.New(text)
}
