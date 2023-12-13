package syncs

import (
	"log"
	"sync"
	"testing"
)

func TestSyncOnce(t *testing.T) {
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			// 只会执行一次
			log.Println("hello.")
		})
	}
	// 这个不会执行
	o.Do(func() {
		log.Println("hello...")
	})
}
