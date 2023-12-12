package goroutines

import "log"

type TaskRunner struct {
	limitChan chan struct{}
}

func NewTaskRunner(size int) *TaskRunner {
	return &TaskRunner{
		limitChan: make(chan struct{}, size),
	}
}

// Run 内部开启go
func (t *TaskRunner) Run(fn func()) {
	t.limitChan <- struct{}{}

	go func() {
		defer func() {
			<-t.limitChan
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		fn()
	}()
}
