package errors

import (
	"bytes"

	"git.bgenius.cn/universe/base/stack"
)

type StackErr struct {
	stack []*Frame
	error
}

func (s *StackErr) copy() *StackErr {
	frames := make([]*Frame, len(s.stack))
	copy(frames, s.stack)
	return &StackErr{
		stack: frames,
		error: s.error,
	}
}

func (s *StackErr) RStackStr() string {
	b := bytes.Buffer{}
	for i := len(s.stack) - 1; i >= 0; i-- {
		if i != len(s.stack)-1 {
			b.WriteString(";")
		}
		b.WriteString(s.stack[i].Name)
	}
	return b.String()
}

func (s *StackErr) FxStack() stack.Fx {
	return stack.Fx{
		"_stack": s.RStackStr(),
	}
}

func (s *StackErr) Cause() error {
	if s.error == nil {
		return nil
	}
	if err, ok := IsStackErr(s.error); ok {
		return err.Cause()
	}
	return s.error
}

func WrapEx(depth int, err error) error {
	if err == nil {
		return nil
	}
	frame := GetFrame(depth)

	e, ok := IsStackErr(err)
	if !ok {
		return &StackErr{
			stack: []*Frame{frame},
			error: err,
		}
	}

	e = e.copy()
	e.stack = append(e.stack, frame)
	return e
}

func IsStackErr(err error) (*StackErr, bool) {
	if e, ok := err.(*StackErr); ok {
		return e, true
	}
	return nil, false
}
