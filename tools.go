package utils

import (
	"sort"
	"time"
)

func Now() string {
	return time.Now().Format("2006-01-02T15:04:05")
}

/**
* 判断一个字符串数组是否包含某字符串元素
 */
func ContainsString(arr []string, keyword string) bool {
	i := sort.SearchStrings(arr, keyword)
	return i != len(arr) && arr[i] == keyword
}
