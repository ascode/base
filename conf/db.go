package conf

import (
	"fmt"
)

type Mysql struct {
	User     string `toml:"user"`
	Pwd      string `toml:"pwd"`
	Addr     string `toml:"addr"`
	DB       string `toml:"db"`
	RTimeout string `toml:"r_timeout"`
	WTimeout string `toml:"w_timeout"`
	IdleConn int    `toml:"idle_conn"`
	OpenConn int    `toml:"open_conn"`
}

func Dsn(cfg interface{}) string {
	switch v := cfg.(type) {
	case *Mysql:
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&readTimeout=%s&writeTimeout=%s&parseTime=true&loc=Asia%%2FShanghai", v.User,
			v.Pwd, v.Addr, v.DB, v.RTimeout, v.WTimeout)
	default:
		panic("unknown conf type")
	}
}
