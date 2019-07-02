package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Frame struct {
	Name string
}

func GetFrame(skip int) *Frame {
	pc, _, ln, ok := runtime.Caller(1 + skip)
	if !ok {
		return nil
	}
	return &Frame{
		Name: fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), ln),
	}
}
