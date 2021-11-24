/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package source

import "time"

type duration time.Duration

func (d *duration) UnmarshalText(text []byte) error {
	du, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = duration(du)
	return nil
}

var Config struct {
	App struct {
		Debug bool
	}
	Auth struct {
		JwtSecret     string
		JwtExpire     duration
		JwtBufferTime duration
	}
	API struct {
		APIHost      string
		APIListen    string
		ClientListen string
		ReadTimeOut  int
		WriteTimeOut int
	}
	Mysql struct {
		Host        string
		User        string
		Password    string
		Database    string
		TablePrefix string
	}
	Redis struct {
		Host string
		Auth string
		DB   int
	}
	Aliyun struct {
		AccessKey    string
		AccessSecret string
		RegionID     string
		OSS          struct {
			BucketName  string
			Domain      string
			Endpoint    string
			UploadDIR   string
			CallBackURL string
		}
	}
	Wechat struct {
		Seller struct {
			AppID          string
			AppSecret      string
			Token          string
			EncodingAESKey string
		}
		Buyer struct {
			AppID          string
			AppSecret      string
			Token          string
			EncodingAESKey string
		}
	}
}
