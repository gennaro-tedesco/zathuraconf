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

	if colourConfig, ok := getColourConfig(colourFile); ok {
		file.WriteString(getDefault())
		file.WriteString("# ----- colour config ----- \n")
		file.WriteString("set notification-error-bg \"" + colourConfig.Background + "\"\n")
		file.WriteString("set notification-error-fg \"" + colourConfig.Error + "\"\n")
		file.WriteString("set notification-warning-bg \"" + colourConfig.Background + "\"\n")
		file.WriteString("set notification-warning-fg \"" + colourConfig.Error + "\"\n")
		file.WriteString("set notification-bg \"" + colourConfig.Background + "\"\n")
		file.WriteString("set notification-fg \"" + colourConfig.Error + "\"\n")
		file.WriteString("set completion-group-bg \"" + colourConfig.Background + "\"\n")
		file.WriteString("set completion-group-fg \"" + colourConfig.Highlight + "\"\n")
		file.WriteString("set completion-bg \"" + colourConfig.Page + "\"\n")
		file.WriteString("set completion-fg \"" + colourConfig.Text + "\"\n")
		file.WriteString("set completion-highlight-bg \"" + colourConfig.Highlight + "\"\n")
		file.WriteString("set completion-highlight-fg \"" + colourConfig.Background + "\"\n")
		file.WriteString("set inputbar-bg \"" + colourConfig.Background + "\"\n")
		file.WriteString("set inputbar-fg \"" + colourConfig.Text + "\"\n")
		file.WriteString("set statusbar-bg \"" + colourConfig.Background + "\"\n")
		file.WriteString("set statusbar-fg \"" + colourConfig.Text + "\"\n")
		file.WriteString("set highlight-color \"" + colourConfig.Highlight + "\"\n")
		file.WriteString("set highlight-active-color \"" + colourConfig.HighlightActive + "\"\n")
		file.WriteString("set default-bg \"" + colourConfig.Background + "\"\n")
		file.WriteString("set default-fg " + colourConfig.Text + "\n")
		file.WriteString("set recolor-lightcolor \"" + colourConfig.Page + "\"\n")
		file.WriteString("set recolor-darkcolor \"" + colourConfig.Text + "\"\n")
		file.WriteString("set index-bg \"" + colourConfig.Page + "\"\n")
		file.WriteString("set index-fg \"" + colourConfig.Text + "\"\n")
		file.WriteString("set index-active-bg \"" + colourConfig.Highlight + "\"\n")
		file.WriteString("set index-active-fg \"" + colourConfig.Background + "\"\n")
		file.WriteString("set recolor " + "true" + "\n")
		fmt.Println("zathura configuration written to", rcFile)
	}
}
