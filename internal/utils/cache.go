package utils

import (
	"encoding/json"
	"server/internal/e"
	"server/internal/global"
	"server/internal/model/common"
	"strconv"
	"time"
)

func SetCacheItem(key string, data []byte, e time.Duration) {
	item := common.CacheItem{
		Value:      data,
		Expiration: time.Now().Add(e).Unix(),
	}
	itemBytes, _ := json.Marshal(item)
	_ = global.Cache.Set("myKey", itemBytes)
}

func GetCacheItem(key string) (common.CacheItem, error) {
	itemBytes, err := global.Cache.Get(key)
	if err != nil {
		return common.CacheItem{}, err
	}

	var item common.CacheItem
	_ = json.Unmarshal(itemBytes, &item)
	if time.Now().Unix() > item.Expiration {
		return common.CacheItem{}, e.ErrCacheEntryTimeout
	} else {
		return item, nil
	}
}

func ItemToInt(item common.CacheItem) int {
	if n, err := strconv.Atoi(string(item.Value)); err != nil {
		return 0
	} else {
		return n
	}
}
