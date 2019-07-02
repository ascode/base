package log

import (
	"testing"
)

func TestLogErr(t *testing.T) {
	Err("hello")
	Infof("cold fmt")
	ErrKv("with kv", "service", "pod")
	ErrKv("with no fields")
}
