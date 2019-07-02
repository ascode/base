package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(in string) string {
	c := md5.New()
	c.Write([]byte(in))
	return hex.EncodeToString(c.Sum(nil))
}
