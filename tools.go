package utils

import (
	"math"
	"sort"
	"strconv"
	"time"
)

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatTime(sec int64) string {
	return time.Unix(sec, 0).Format("2006-01-02 15:04:05")
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

func ParseFloat(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		Errorf("Error parse float %s: %s", s, err.Error())
	}
	return v
}

func GetEarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := float64(6371000) // 6378137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	rt := dist * radius
	if math.IsNaN(rt) {
		rt = 0
	}
	return rt
}
