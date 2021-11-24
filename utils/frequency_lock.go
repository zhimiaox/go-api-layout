/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package utils

import (
	"sync"
	"time"
)

type frequencyLockItem struct {
	Key      string
	LifeSpan time.Duration // 生命周期
	CreateOn time.Time     // 创建时间
}

type frequencyLock struct {
	sync.RWMutex
	CleanerDuration time.Duration                // 触发定时清理器的时间
	Cleaner         *time.Timer                  // 定时清理器
	Items           map[string]frequencyLockItem // 子集
}

// NewLockTable 新建
func NewLockTable() *frequencyLock {
	return &frequencyLock{
		Items: make(map[string]frequencyLockItem),
	}
}

// IsLock 是否锁
func (l *frequencyLock) IsLock(key string, lockTime time.Duration) bool {
	l.Lock()
	if item, ok := l.Items[key]; ok {
		l.Unlock()
		if time.Since(item.CreateOn) > item.LifeSpan {
			l.cleanerCheck()
			return false
		}
		return true
	}
	l.Items[key] = frequencyLockItem{
		Key:      key,
		LifeSpan: lockTime,
		CreateOn: time.Now(),
	}
	cleannerDuraction := l.CleanerDuration
	l.Unlock()
	if cleannerDuraction == 0 {
		l.cleanerCheck()
	}
	return false
}

func (l *frequencyLock) cleanerCheck() {
	l.Lock()
	defer l.Unlock()
	if l.Cleaner != nil {
		l.Cleaner.Stop()
	}
	// 遍历当前限制的key, 遇到过期的将其删掉
	// 其余的则从中找到最近一个将要过期的key并且将它还有多少时间过期作为下一次清理任务的定时时间
	now := time.Now()
	smallestDuracton := 0 * time.Second
	for key, item := range l.Items {
		lifeSpan := item.LifeSpan
		createOn := item.CreateOn
		if now.Sub(createOn) >= lifeSpan {
			delete(l.Items, key)
		} else {
			if smallestDuracton == 0 || lifeSpan-now.Sub(createOn) < smallestDuracton {
				smallestDuracton = lifeSpan - now.Sub(createOn)
			}
		}
	}
	l.CleanerDuration = smallestDuracton
	// 将最近一个将要过期的key距离现在的时间作为启动清理任务的定时时间
	if l.CleanerDuration > 0 {
		l.Cleaner = time.AfterFunc(l.CleanerDuration, func() {
			go l.cleanerCheck()
		})
	}
}
