package conf

type Redis struct {
	Addr        string `toml:"addr"`
	Pwd         string `toml:"pwd"`
	ReadTimeout int    `toml:"read_timeout"` // second
}
