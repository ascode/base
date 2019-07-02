package util

import (
	"bytes"
	"strconv"
)

func Int64s(args ...int64) string {
	var b bytes.Buffer
	for i, arg := range args {
		if i > 0 {
			b.WriteString(`,`)
		}
		b.WriteString(strconv.FormatInt(arg, 10))
	}
	return b.String()
}
