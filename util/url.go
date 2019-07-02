package util

import (
	"strings"
)

func GetDomain(host string) string {
	host = strings.TrimPrefix(host, "http://")
	host = strings.TrimPrefix(host, "https://")
	return strings.Split(host, ":")[0]
}
