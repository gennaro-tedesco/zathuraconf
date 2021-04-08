package cmd

import (
	"os/user"

	"github.com/spf13/cobra"
)

var usr, _ = user.Current()
var dir = usr.HomeDir
var ZATHURARC = dir + "/.config/zathura/zathurarc"

var rootCmd = &cobra.Command{
	Use:   "zathuraconf",
	Args:  cobra.ExactArgs(1),
	Short: "change zathura colour scheme",
	Long:  `change zathura colour scheme from command line`,
	Run: func(cmd *cobra.Command, args []string) {
		rc_file, _ := cmd.Flags().GetString("path")
		writeConfig(args[0], rc_file)
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
  Provide a config.json file of the form (example)

  config.json
  {
	"page": "#073642",
	"text": "#93a1a1",
	"background": "#002b36",
	"highlight": "#859900",
	"highlight_active": "#268bd2",
	"error": "#cb4b16"
  }

  and simply run
      zathuraconf config.json

  If your zathurarc is in a different path than ~/.config/zathura/zathurarc
	  zathurarc config.json -p /path/to/zathuraconf

Help commands:
  version     print current version

Flags:
  -h, --help   help for zathuraconf
  -p, --path   path to zathuraconf location: defaults to ~/.config/zathura/zathurarc
`
}
