package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "root",
}

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "path of config file")
	rootCmd.PersistentFlags().BoolVarP(&isFormatResp, "format-resp", "f", true, "format the respose of restful")
}

func Exec() error {
	return rootCmd.Execute()
}
