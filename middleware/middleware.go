/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package middleware

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/zhimiaox/go-api-layout/models/vo"
	"github.com/zhimiaox/go-api-layout/source"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

// CorsMiddleware 跨域
func CorsMiddleware() gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowHeaders = append(cfg.AllowHeaders, vo.Authorization)
	return cors.New(cfg)
}

// LogMiddleware 日志中间件
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqURL := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求ip
		clientIP := c.ClientIP()
		// 日志格式
		logrus.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqURL,
		}).Info()
	}
}

// LoginAuthMiddleware jwt鉴权
func LoginAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bearer
		tokens := strings.Split(c.GetHeader(vo.Authorization), " ")
		if len(tokens) != 2 || tokens[0] != "Bearer" || tokens[1] == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token, err := jwt.ParseWithClaims(tokens[1], &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(source.Config.Auth.JwtSecret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 过渡期返回新token
		now := time.Now()
		if now.Add(time.Duration(source.Config.Auth.JwtBufferTime)).Unix() > claims.ExpiresAt {
			claims.Issuer, _ = os.Hostname()
			claims.IssuedAt = now.Unix()
			claims.ExpiresAt = now.Add(time.Duration(source.Config.Auth.JwtExpire)).Unix()
			claims.NotBefore = now.Unix()
			if newToken, err2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
				SignedString([]byte(source.Config.Auth.JwtSecret)); err2 == nil {
				c.Header(vo.Authorization, newToken)
			}
		}
		uid, err := strconv.Atoi(claims.Id)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set(vo.UID, uid)
		c.Set(vo.APPID, claims.Audience)
		c.Set(vo.Authorization, claims)
		c.Next()
	}
}
