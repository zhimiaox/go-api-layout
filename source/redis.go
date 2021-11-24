/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package source

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	redisInstance *redis.Client
	redisOnce     sync.Once
)

func GetRedis() *redis.Client {
	redisOnce.Do(func() {
		redisInstance = redis.NewClient(&redis.Options{
			Addr:     Config.Redis.Host,
			Password: Config.Redis.Auth, // no password set
			DB:       Config.Redis.DB,   // use default DB
		})
	})
	return redisInstance
}
