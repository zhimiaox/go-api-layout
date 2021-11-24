/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package sdk

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io"
	"path"
	"sync"
	"time"

	"github.com/zhimiaox/go-api-layout/source"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	ossInstance OSS
	ossOnce     sync.Once
)

type OSS interface {
	GetBucket() (*oss.Bucket, error)
	GetPolicyToken(rootPath string) (*PolicyToken, error)
}

type ossImpl struct{}

func (o *ossImpl) GetBucket() (*oss.Bucket, error) {
	conf := source.Config
	client, err := oss.New(
		conf.Aliyun.OSS.Endpoint,
		conf.Aliyun.AccessKey,
		conf.Aliyun.AccessSecret,
	)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(conf.Aliyun.OSS.BucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func (o *ossImpl) GetPolicyToken(rootPath string) (*PolicyToken, error) {
	conf := source.Config.Aliyun
	rootPath = path.Join(conf.OSS.UploadDIR, rootPath)
	expireTime := time.Now().Add(60 * time.Second)
	var tokenExpire = expireTime.UTC().Format("2006-01-02T15:04:05Z")
	result, err := json.Marshal(ConfigStruct{
		Expiration: tokenExpire,
		Conditions: [][]string{
			{"starts-with", "$key", rootPath},
			// {"content-length-range", 0, 1048576000},
		},
	})
	if err != nil {
		return nil, err
	}
	resultB64Str := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(sha1.New, []byte(conf.AccessSecret))
	_, err = io.WriteString(h, resultB64Str)
	if err != nil {
		return nil, err
	}
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	callback, err := json.Marshal(CallbackParam{
		CallbackURL:      conf.OSS.CallBackURL,
		CallbackBody:     "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}",
		CallbackBodyType: "application/x-www-form-urlencoded",
	})
	if err != nil {
		return nil, err
	}
	return &PolicyToken{
		AccessKeyID: conf.AccessKey,
		Host:        conf.OSS.Domain,
		Expire:      expireTime.Unix(),
		Signature:   signedStr,
		Policy:      resultB64Str,
		Directory:   rootPath,
		Callback:    base64.StdEncoding.EncodeToString(callback),
	}, nil
}

func GetOSS() OSS {
	ossOnce.Do(func() {
		ossInstance = &ossImpl{}
	})
	return ossInstance
}

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type CallbackParam struct {
	CallbackURL      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}

type PolicyToken struct {
	AccessKeyID string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
}
