package grace

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	isClosed   int32
	processing sync.WaitGroup
	ErrClosing = errors.New("server is closing")
)

func ExitMark() {
	atomic.StoreInt32(&isClosed, 1)
}

func WaitTaskDone() {
	processing.Wait()
}

func CountInTask() bool {
	if atomic.LoadInt32(&isClosed) == 1 {
		return false
	}
	processing.Add(1)
	return true
}

func DoneTask() {
	processing.Done()
}
