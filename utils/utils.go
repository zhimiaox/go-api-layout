/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package utils

import (
	"encoding/json"
	"errors"
)

// DecodeJSONPayload 解码载荷
func DecodeJSONPayload(payload []byte, res interface{}, secret string) error {
	raw := AESDecrypt(string(payload), secret)
	if raw == "" {
		return errors.New("解码失败")
	}
	err := json.Unmarshal([]byte(raw), res)
	if err != nil {
		return err
	}
	return nil
}

// EncodeJSONPayload 加密载荷
func EncodeJSONPayload(payload interface{}, secret string) ([]byte, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	aes := AESEncrypt(string(payloadJSON), secret)
	if aes == "" {
		return nil, errors.New("加密失败")
	}
	return []byte(aes), nil
}
