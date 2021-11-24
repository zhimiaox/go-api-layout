/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package sdk

import (
	"testing"
)

func Test_ossImpl_GetPolicyToken(t *testing.T) {
	token, err := GetOSS().GetPolicyToken("other")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}
