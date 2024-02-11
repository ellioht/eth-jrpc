package config

import "github.com/spf13/cobra"

func init() {
	cobra.OnInitialize(initConfig)
}
