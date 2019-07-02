package cmd

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

var isFormatResp bool

var cfgPath string

func InitConf(cfg interface{}) error {
	var err error

	cobra.OnInitialize(func() {
		err = initConf(cfg)
	})

	return nil
}

func initConf(cfg interface{}) error {
	if cfgPath == "" { // check current project conf folder
		if p, err := filepath.Abs("."); err == nil {
			p = filepath.Join(p, "conf", "conf.toml")
			if fi, err := os.Stat(p); err == nil && !fi.IsDir() {
				cfgPath = p
			}
		}
	}
	if cfgPath != "" {
		if _, err := toml.DecodeFile(cfgPath, cfg); err != nil {
			return err
		}
	}
	return nil
}
