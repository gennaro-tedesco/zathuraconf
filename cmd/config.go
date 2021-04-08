package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func getDefault() string {
	return `
# window frame configurations
set window-height 3000
set window-width 1000
set statusbar-h-padding 0
set statusbar-v-padding 0
set page-padding 10
set window-title-page true
set guioptions vc

# set scrolling options
set scroll-step 40

# font size
set font "Fira Code 15"

# clipboard configurations
set selection-clipboard clipboard

# mappings
map i recolor

unmap f
map f toggle_fullscreen
map [fullscreen] f toggle_fullscreen
`
}

type ColourConfig struct {
	Page             string `json: "page"`
	Text             string `json: "text"`
	Background       string `json: "background"`
	Highlight        string `json: "highlight"`
	Highlight_active string `json: "highlight_active"`
	Error            string `json: "error"`
}

func getColourConfig() ColourConfig {
	file, err := ioutil.ReadFile("examples/solarized_dark.json")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file doesn't exist")
			os.Exit(1)
		}
	}

	colour_config := ColourConfig{}
	_ = json.Unmarshal([]byte(file), &colour_config)
	return colour_config
}

func writeConfig() {
	file, err := os.Create("text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(getDefault())
}