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
map i recolor

unmap f
map f toggle_fullscreen
map [fullscreen] f toggle_fullscreen

`
}

type ColourConfig struct {
	Page             string `json:"page" validate:"iscolor"`
	Text             string `json:"text" validate:"iscolor"`
	Background       string `json:"background" validate:"iscolor"`
	Highlight        string `json:"highlight" validate:"iscolor"`
	Highlight_active string `json:"highlight_active" validate:"iscolor"`
	Error            string `json:"error" validate:"iscolor"`
}

func isValidConfig(colour_config ColourConfig) bool {
	v := validator.New()
	v_err := v.Struct(colour_config)

	if v_err != nil {
		for _, v_err := range v_err.(validator.ValidationErrors) {
			fmt.Println(v_err)
		}
		return false
	} else {
		return true
	}
}

func getColourConfig(filename string) ColourConfig {
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

	colour_config := ColourConfig{}
	_ = json.Unmarshal([]byte(file), &colour_config)

	if !isValidConfig(colour_config) {
		os.Exit(1)
	}
	return colour_config
}

func writeConfig(colour_file string, rc_file string) {
	file, err := os.Create(rc_file)
	if err != nil {
		return
	}
	defer file.Close()

	colour_config := getColourConfig(colour_file)
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
	file.WriteString("set highlight-active-color \"" + colour_config.Highlight_active + "\"\n")
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
