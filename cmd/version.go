package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VERSION number: change manually
const VERSION = "1.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print current stargazer version",
	Long:  "print current stargazer version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
