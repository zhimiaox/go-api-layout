/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package utils

import (
	"fmt"
	"testing"
)

func TestSuperConvert(t *testing.T) {
	type b struct {
		A1 string
		B1 int
		G1 bool
	}
	type a struct {
		A1 string
	}
	as := a{
		A1: "123",
	}
	bs := b{}
	SuperConvert(&as, &bs)
	fmt.Printf("%#v", bs)
}
