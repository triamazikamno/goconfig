package goConfig

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/crgimenes/goConfig/goEnv"
)

// Settings default
type Settings struct {
	// Path sets default config path
	Path string
	// File name of default config file
	File string
	// FileRequired config file required
	FileRequired bool
}

// Setup Pointer to internal variables
var Setup *Settings

func init() {
	Setup = &Settings{
		Path:         "./",
		File:         "config.json",
		FileRequired: false,
	}
}

// LoadJSON config file
func LoadJSON(config interface{}) (err error) {
	configFile := Setup.Path + Setup.File
	file, err := os.Open(configFile)
	if os.IsNotExist(err) && !Setup.FileRequired {
		err = nil
		return
	} else if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return
	}

	return
}

// Load config file
func Load(config interface{}) (err error) {

	err = LoadJSON(config)
	if err != nil {
		return
	}

	err = goEnv.Parse(config)
	if err != nil {
		return
	}

	return
}

// Save config file
func Save(config interface{}) (err error) {
	_, err = os.Stat(Setup.Path)
	if os.IsNotExist(err) {
		os.Mkdir(Setup.Path, 0700)
	} else if err != nil {
		return
	}

	configFile := Setup.Path + Setup.File

	_, err = os.Stat(configFile)
	if err != nil {
		return
	}

	b, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return
	}

	err = ioutil.WriteFile(configFile, b, 0644)
	if err != nil {
		return
	}
	return
}
