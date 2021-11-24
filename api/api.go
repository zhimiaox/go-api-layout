/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package api

import (
	"net/http"
	"time"

	v1 "github.com/zhimiaox/go-api-layout/api/v1"
	"github.com/zhimiaox/go-api-layout/middleware"
	"github.com/zhimiaox/go-api-layout/source"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/zhimiaox/go-api-layout/docs"
)

var Router *gin.Engine

func initRoute() {
	Router = gin.New()
	Router.Use(gin.Recovery(), middleware.LogMiddleware())
	// 状态监控
	ginpprof.Wrap(Router)
	/* ------ 文档模块 ------- */
	if source.Config.App.Debug {
		Router.Use(middleware.CorsMiddleware())
		Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	/* ------ sdk ------- */
	v1Route := Router.Group("/api/v1")
	{
		// 公共
		// v1.RegisterMember(v1Route)
		v1.RegisterHello(v1Route)
	}
}

func Start() {
	if source.Config.App.Debug {
		gin.SetMode(gin.DebugMode)
	}
	// 初始化route
	initRoute()
	httpServer := &http.Server{
		Addr:           source.Config.API.APIListen,
		Handler:        Router,
		ReadTimeout:    time.Duration(source.Config.API.ReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(source.Config.API.WriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logrus.Infof("Start HTTP Service Listening %s", source.Config.API.APIListen)
	if err := httpServer.ListenAndServe(); err != nil {
		logrus.Error("api server running error", err)
	}
}
