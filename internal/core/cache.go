package core

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"time"
)

// 参考链接：https://liuqh.icu/2021/06/15/go/package/14-bigcache/

func InitializeCache() *bigcache.BigCache {
	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	return cache
}
