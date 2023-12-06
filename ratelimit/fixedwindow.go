package ratelimit

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdsCli *redis.Client

func init() {
	rdsCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// FixedWindowRateLimit 固定窗口计数算法
// 利用redis.Incr是原子操作
func FixedWindowRateLimit(key string, interval time.Duration, limitNum int64) bool {
	// 只要不超过1s，window值都是一样的
	window := time.Now().Unix() / int64(interval.Seconds())
	newKey := fmt.Sprintf("%s:%d:%d:%d", key, interval, limitNum, window)

	startCount := 0
	// 如果已经存在key，不会执行，也不会报错
	if _, err := rdsCli.SetNX(newKey, startCount, time.Second*1).Result(); err != nil {
		panic(err)
	}

	// 原子操作
	curNum, err := rdsCli.Incr(newKey).Result()
	if err != nil {
		panic(err)
	}

	return curNum <= limitNum
}
