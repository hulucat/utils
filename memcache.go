package utils

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var Mclient *memcache.Client

func GetCache(key string) (value string) {
	it, err := Mclient.Get(key)
	if err != nil {
		return ""
	}
	return string(it.Value)
}

/**
* expiration: 多少秒以后超时
 */
func SetCache(key, value string, expiration int32) {
	if err := Mclient.Set(&memcache.Item{Key: key, Value: []byte(value), Expiration: expiration}); err != nil {
		Errorf("Error set memcache: %s - %s", key, value)
	}
}
