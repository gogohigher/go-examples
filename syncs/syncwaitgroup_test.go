package syncs

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}
	wg.Wait()
	log.Println("main done.")
}

func worker(id int) {
	log.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	log.Printf("worker %d done\n", id)
}
