package stack

// 实际使用中，key 必须为 string 类型
type Fx map[string]interface{}

func (f Fx) clone() Fx {
	fs := make(Fx, len(f))
	for k, v := range f {
		fs[k] = v
	}
	return fs
}
