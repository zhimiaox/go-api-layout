/*
 * 纸喵软件
 * Copyright (c) 2017~2021 http://zhimiaox.cn All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2021/11/24 下午10:00
 */

package utils

import "sort"

// IdsUniqueFitter ids去重
func IdsUniqueFitter(ids []int) []int {
	sort.Ints(ids)
	var newIds []int
	var lastID int
	for i, id := range ids {
		if i == 0 {
			lastID = id
			newIds = append(newIds, id)
		} else if id != lastID {
			lastID = id
			newIds = append(newIds, id)
		}
	}
	return newIds
}
