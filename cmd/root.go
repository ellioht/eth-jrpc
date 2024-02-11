package cmd

import (
	"fmt"
	"github.com/ellioht/eth-jrpc/pkg/version"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "root",
	Short:   "root is the root command",
	Version: version.GetVersion(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("eth-jrpc %s", version.GetVersion()))
	},
}

func Root() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
