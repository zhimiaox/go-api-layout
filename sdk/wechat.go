/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zhimiaox/go-api-layout/consts"
	"github.com/zhimiaox/go-api-layout/source"

	"github.com/go-redis/redis/v8"
	wechatSdk "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

func NewSellerMiniProgramClient() *miniprogram.MiniProgram {
	wc := wechatSdk.NewWechat()
	return wc.GetMiniProgram(&miniConfig.Config{
		AppID:     source.Config.Wechat.Seller.AppID,
		AppSecret: source.Config.Wechat.Seller.AppSecret,
		Cache:     NewCache(consts.MPCacheSellerPrefix),
	})
}

func NewBuyerMiniProgramClient() *miniprogram.MiniProgram {
	wc := wechatSdk.NewWechat()
	return wc.GetMiniProgram(&miniConfig.Config{
		AppID:     source.Config.Wechat.Buyer.AppID,
		AppSecret: source.Config.Wechat.Buyer.AppSecret,
		Cache:     NewCache(consts.MPCacheBuyerPrefix),
	})
}

// Cache redis cache
type Cache struct {
	redis  *redis.Client
	prefix string
}

// NewCache 实例化
func NewCache(prefix string) *Cache {
	client := &Cache{
		prefix: prefix,
		redis:  source.GetRedis(),
	}
	return client
}

// Get 获取一个值
func (c *Cache) Get(key string) interface{} {
	var data []byte
	var err error
	if data, err = c.redis.Get(context.Background(), fmt.Sprintf("%s:%s", c.prefix, key)).Bytes(); err != nil {
		return nil
	}
	var reply interface{}
	if err = json.Unmarshal(data, &reply); err != nil {
		return nil
	}
	return reply
}

// Set 设置一个值
func (c *Cache) Set(key string, val interface{}, timeout time.Duration) (err error) {
	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return
	}
	err = c.redis.Set(context.Background(), fmt.Sprintf("%s:%s", c.prefix, key), data, timeout).Err()
	return
}

// IsExist 判断key是否存在
func (c *Cache) IsExist(key string) bool {
	i, err := c.redis.Exists(context.Background(), fmt.Sprintf("%s:%s", c.prefix, key)).Result()
	if err != nil {
		return false
	}
	return i > 0
}

// Delete 删除
func (c *Cache) Delete(key string) error {
	err := c.redis.Del(context.Background(), fmt.Sprintf("%s:%s", c.prefix, key)).Err()
	if err != nil {
		return err
	}
	return nil
}
