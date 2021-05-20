package threading

import (
	"bytes"
	"fmt"
	"gin-api/app/utility/errorx"
	"runtime"
	"strconv"
)

// GoSafe 防止崩溃
func GoSafe(fn func() error, cleanups ...func()) {
	go RunSafe(fn, cleanups...)
}

// RoutineID Only for debug, never use it in production
func RoutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	// if error, just return 0
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return n
}

// RunSafe 防止函数 crash
func RunSafe(fn func() error, cleanups ...func()) (err error, crash error) {
	defer func() {
		for _, cleanup := range cleanups {
			cleanup()
		}

		if p := recover(); p != nil {
			if e, ok := p.(error); ok {
				crash = e
				return
			}
			crash = errorx.New(fmt.Sprintf("%v", p))
		}
	}()

	err = fn()

	return
}
