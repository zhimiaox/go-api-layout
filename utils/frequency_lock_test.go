/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package utils

import (
	"testing"
	"time"
)

func TestNewLockTable(t *testing.T) {
	lock := NewLockTable()
	isLock := lock.IsLock("a", 2*time.Second)
	if isLock {
		t.Fatal("首次上锁失败")
	}
	isLock = lock.IsLock("a", 2*time.Second)
	if !isLock {
		t.Fatal("二次上锁穿透")
	}
	time.Sleep(2 * time.Second)
	isLock = lock.IsLock("a", 2*time.Second)
	if isLock {
		t.Fatal("解锁失败")
	}
}
