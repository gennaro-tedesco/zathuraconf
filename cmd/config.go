package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
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

unmap f
map f toggle_fullscreen
map [fullscreen] f toggle_fullscreen

`
}

type ColourConfig struct {
	Page            string `json:"page" validate:"iscolor"`
	Text            string `json:"text" validate:"iscolor"`
	Background      string `json:"background" validate:"iscolor"`
	Highlight       string `json:"highlight" validate:"iscolor"`
	HighlightActive string `json:"highlight_active" validate:"iscolor"`
	Error           string `json:"error" validate:"iscolor"`
}

func isValidConfig(colourConfig ColourConfig) bool {
	v := validator.New()
	vErr := v.Struct(colourConfig)

	if vErr != nil {
		for _, vErr := range vErr.(validator.ValidationErrors) {
			fmt.Println(vErr)
		}
		fmt.Println("please provide valid json")
		return false
	}
	return true
}

func getColourConfig(filename string) (ColourConfig, bool) {
	if filepath.Ext(filename) != ".json" {
		fmt.Printf("please provide a valid json file")
		os.Exit(1)
	}
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("file %s doesn't exist", filename)
			os.Exit(1)
		}
	}

	colourConfig := ColourConfig{}
	_ = json.Unmarshal([]byte(file), &colourConfig)

	if !isValidConfig(colourConfig) {
		return colourConfig, false
	}
	return colourConfig, true
}

func writeConfig(colourFile string, rcFile string) {
	file, err := os.Create(rcFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	if colour_config, ok := getColourConfig(colourFile); ok {
		file.WriteString(getDefault())
		file.WriteString("# ----- colour config ----- \n")
		file.WriteString("set notification-error-bg \"" + colour_config.Background + "\"\n")
		file.WriteString("set notification-error-fg \"" + colour_config.Error + "\"\n")
		file.WriteString("set notification-warning-bg \"" + colour_config.Background + "\"\n")
		file.WriteString("set notification-warning-fg \"" + colour_config.Error + "\"\n")
		file.WriteString("set notification-bg \"" + colour_config.Background + "\"\n")
		file.WriteString("set notification-fg \"" + colour_config.Error + "\"\n")
		file.WriteString("set completion-group-bg \"" + colour_config.Background + "\"\n")
		file.WriteString("set completion-group-fg \"" + colour_config.Highlight + "\"\n")
		file.WriteString("set completion-bg \"" + colour_config.Page + "\"\n")
		file.WriteString("set completion-fg \"" + colour_config.Text + "\"\n")
		file.WriteString("set completion-highlight-bg \"" + colour_config.Highlight + "\"\n")
		file.WriteString("set completion-highlight-fg \"" + colour_config.Background + "\"\n")
		file.WriteString("set inputbar-bg \"" + colour_config.Background + "\"\n")
		file.WriteString("set inputbar-fg \"" + colour_config.Text + "\"\n")
		file.WriteString("set statusbar-bg \"" + colour_config.Background + "\"\n")
		file.WriteString("set statusbar-fg \"" + colour_config.Text + "\"\n")
		file.WriteString("set highlight-color \"" + colour_config.Highlight + "\"\n")
		file.WriteString("set highlight-active-color \"" + colour_config.HighlightActive + "\"\n")
		file.WriteString("set default-bg \"" + colour_config.Background + "\"\n")
		file.WriteString("set default-fg " + colour_config.Text + "\n")
		file.WriteString("set recolor-lightcolor \"" + colour_config.Page + "\"\n")
		file.WriteString("set recolor-darkcolor \"" + colour_config.Text + "\"\n")
		file.WriteString("set index-bg \"" + colour_config.Page + "\"\n")
		file.WriteString("set index-fg \"" + colour_config.Text + "\"\n")
		file.WriteString("set index-active-bg \"" + colour_config.Highlight + "\"\n")
		file.WriteString("set index-active-fg \"" + colour_config.Background + "\"\n")
		file.WriteString("set recolor " + "true" + "\n")
	}
}
