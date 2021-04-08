package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const ZATHURARC = "~/.config/zathura/zathurarc"

var rootCmd = &cobra.Command{
	Use:   "zathuraconf",
	Short: "change zathura colour scheme",
	Long:  `change zathura colour scheme from command line`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ZATHURARC)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.SetHelpTemplate(GetRootHelp())
	rootCmd.Flags().StringP("path", "p", ZATHURARC, "path to zathurarc location")
}

func GetRootHelp() string {
	return `
zathuraconf: change zathura colour scheme

Arguments:
  config.json    configuration file containing colour scheme

Usage:
  Provide a config.json file of the form
  ...
  and simply run
      zathuraconf config.json

  If your zathuraconf is in a different path than ~/.config/zathura/zathurarc
	  zathuraconf config.json -p /path/to/zathuraconf

Help commands:
  version     print current version

Flags:
  -h, --help   help for zathuraconf
  -p, --path   path to zathuraconf location: defaults to ~/.config/zathura/zathurarc
`
}
