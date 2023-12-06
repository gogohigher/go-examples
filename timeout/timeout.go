package timeout

import (
	"context"
	"time"
)

// 正确的超时控制
func handle(parentCtx context.Context, fn func() error) error {
	ctx, cancel := context.WithTimeout(parentCtx, time.Second*1)
	defer cancel()
	// 必须是有缓冲的channel，如果是无缓冲的，则会导致goroutine的泄露
	done := make(chan error, 1)
	go func() {
		done <- fn()
	}()
	select {
	// 当cancel()执行的的时候，会调用这个
	case <-ctx.Done():
		return ctx.Err()

	case err := <-done:
		return err
	}
}
