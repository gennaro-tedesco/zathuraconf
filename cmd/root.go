package cmd

import (
	"os/user"

	"github.com/spf13/cobra"
)

var usr, _ = user.Current()
var dir = usr.HomeDir
var zathurarcfile = dir + "/.config/zathura/zathurarc"

var rootCmd = &cobra.Command{
	Use:   "zathuraconf",
	Args:  cobra.ExactArgs(1),
	Short: "change zathura colour scheme",
	Long:  `change zathura colour scheme from command line`,
	Run: func(cmd *cobra.Command, args []string) {
		rcFile, _ := cmd.Flags().GetString("path")
		writeConfig(args[0], rcFile)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.SetHelpTemplate(getRootHelp())
	rootCmd.Flags().StringP("path", "p", zathurarcfile, "path to zathurarc location")
}

func getRootHelp() string {
	return `
zathuraconf: change zathura colour scheme

Arguments:
  config.json    configuration file containing colour scheme

  {
	"page": "#073642",
	"text": "#93a1a1",
	"background": "#002b36",
	"highlight": "#859900",
	"highlight_active": "#268bd2",
	"error": "#cb4b16"
  }

Usage:
  zathuraconf config.json [flag]

Help commands:
  version     print current version

Flags:
  -h, --help   help for zathuraconf
  -p, --path   path to zathuraconf location: defaults to ~/.config/zathura/zathurarc
`
}
