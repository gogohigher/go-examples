package syncs

import (
	"fmt"
	"sync"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("creating new user.")
			return &User{}
		},
	}

	// 从池中获取对象
	user := pool.Get().(*User)
	user.Name = "zhangsan"
	user.Age = 18
	fmt.Printf("user: %+v\n", user)

	// 将对象放回池中
	pool.Put(user)

	// 再次从池中获取对象
	user2 := pool.Get().(*User)
	fmt.Printf("user: %+v\n", user2)
}
