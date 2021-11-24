/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package source

import (
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlInstance *gorm.DB
	mysqlOnce     sync.Once
)

func GetMysql() *gorm.DB {
	mysqlOnce.Do(func() {
		var err error
		mysqlInstance, err = gorm.Open(
			mysql.New(mysql.Config{
				DSN: fmt.Sprintf(
					"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
					Config.Mysql.User,
					Config.Mysql.Password,
					Config.Mysql.Host,
					Config.Mysql.Database,
				),
				DefaultStringSize:         256,   // string 类型字段的默认长度
				DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
				DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
				DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
				SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
			}),
			&gorm.Config{},
		)
		if err != nil {
			logrus.Fatalf("models.Setup err: %v", err)
			return
		}
		if Config.App.Debug {
			mysqlInstance = mysqlInstance.Debug()
		}
		sqlDB, err := mysqlInstance.DB()
		if err != nil {
			logrus.Fatalf("models.Setup err: %v", err)
			return
		}
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
	return mysqlInstance
}
