package cmd

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestValidConfig(t *testing.T) {
	valid_json_file := "../examples/solarized_dark.json"

	valid_file, _ := ioutil.ReadFile(valid_json_file)

	valid_colour_config := ColourConfig{}
	_ = json.Unmarshal([]byte(valid_file), &valid_colour_config)
	assert.Equal(t, true, isValidConfig(valid_colour_config))

}

func TestInValidConfig(t *testing.T) {
	invalid_json_file := "../examples/invalid.json"

	invalid_file, _ := ioutil.ReadFile(invalid_json_file)

	invalid_colour_config := ColourConfig{}
	_ = json.Unmarshal([]byte(invalid_file), &invalid_colour_config)
	assert.Equal(t, false, isValidConfig(invalid_colour_config))
}
