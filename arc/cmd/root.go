package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var dockerPassword string

var rootCmd = &cobra.Command{
	Use:   "arc",
	Short: "A helpful CircleCI and GitHub tool.",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize()
}
