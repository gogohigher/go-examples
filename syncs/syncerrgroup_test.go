package syncs

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestErrGroup(t *testing.T) {
	case1()
}

/*
1. Go方法创建一个Goroutine，并在其中执行传入的函数
2. Wait方法等待所有Goroutine全部返回
  - 如果返回错误，这一组Goroutine最少返回一个错误
  - 如果返回空值，所有Goroutine都成功执行

3. 假设有多个Goroutine出现了错误，那么只会返回第一个出现的错误，后续的错误都会被舍弃
*/
func case1() {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	err := g.Wait()
	if err == nil {
		log.Println("Successfully fetched all URLs.")
	} else {
		log.Println("err: ", err)
	}
}

func case2() {
	var eg errgroup.Group

	eg.Go(func() error {
		fmt.Println("task 3")
		time.Sleep(2 * time.Second) // 如果加上这个，就只打印出task 2 error的错误
		return fmt.Errorf("task 3 error")
	})

	eg.Go(func() error {
		fmt.Println("task 2")
		return fmt.Errorf("task 2 error")
	})

	eg.Go(func() error {
		fmt.Println("task 1")
		return nil
	})

	// 网上都在说可以同时记录两个协程的错误，我这里测试只能记录一个
	err := eg.Wait()
	if err != nil {
		fmt.Println("main err: ", err.Error())
	}
}
