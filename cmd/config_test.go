package cmd

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestValidConfig(t *testing.T) {
	validJsonFile := "../examples/solarized_dark.json"

	validFile, _ := ioutil.ReadFile(validJsonFile)

	validColourConfig := ColourConfig{}
	_ = json.Unmarshal([]byte(validFile), &validColourConfig)
	assert.Equal(t, true, isValidConfig(validColourConfig))

}

func TestInValidConfig(t *testing.T) {
	invalidJsonFile := "../examples/invalid.json"

	invalidFile, _ := ioutil.ReadFile(invalidJsonFile)

	invalidColourConfig := ColourConfig{}
	_ = json.Unmarshal([]byte(invalidFile), &invalidColourConfig)
	assert.Equal(t, false, isValidConfig(invalidColourConfig))
}
