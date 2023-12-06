package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestFixedWindowRateLimit(t *testing.T) {
	for i := 0; i < 20; i++ {
		go func() {
			ok := FixedWindowRateLimit("t1", 1*time.Second, 5)
			fmt.Println("rateLimit is:", ok)
		}()
	}
	time.Sleep(time.Second * 3)
}
