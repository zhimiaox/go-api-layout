/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zhimiaox/go-api-layout/api"
	"github.com/zhimiaox/go-api-layout/source"

	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

// @title 纸喵 API
// @version 1.0
// @description 纸喵软件系列
// @termsOfService http://zhimiao.org

// @contact.name API Support
// @contact.url http://tools.zhimiao.org
// @contact.email mail@xiaoliu.org

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http https
// @BasePath /api

// @securityDefinitions.apikey ApiAuth
// @in header
// @name Authorization
func main() {
	if err := configor.Load(&source.Config, "config.toml"); err != nil {
		panic(err)
	}
	http.DefaultClient.Timeout = time.Minute
	go api.Start()
	logrus.Infoln("signal received, server closed. ", waitForSignal())
}

func waitForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
