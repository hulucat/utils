package utils

import (
	"sort"
	"strconv"
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

func ParseInt(s string) int {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		Errorf("Error parse int %s: %s", s, err.Error())
	}
	return int(v)
}

func ParseInt32(s string) int32 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		Errorf("Error parse int32 %s: %s", s, err.Error())
	}
	return int32(v)
}

func ParseInt64(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		Errorf("Error parse int64 %s: %s", s, err.Error())
	}
	return v
}
