package json

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/crgimenes/goConfig"
)

func init() {
	f := goConfig.Fileformat{
		Extension:   ".json",
		Save:        SaveJSON,
		Load:        LoadJSON,
		PrepareHelp: PrepareHelp,
	}
	goConfig.Formats = append(goConfig.Formats, f)
}

// LoadJSON config file
func LoadJSON(config interface{}) (err error) {
	configFile := goConfig.Path + goConfig.File
	file, err := os.Open(configFile)
	if os.IsNotExist(err) && !goConfig.FileRequired {
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

// SaveJSON config file
func SaveJSON(config interface{}) (err error) {
	_, err = os.Stat(goConfig.Path)
	if os.IsNotExist(err) {
		err = os.Mkdir(goConfig.Path, os.ModePerm)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	}

	configFile := goConfig.Path + goConfig.File

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

// PrepareHelp return help string for this file format.
func PrepareHelp(config interface{}) (help string, err error) {
	var helpAux []byte
	helpAux, err = json.MarshalIndent(&config, "", "    ")
	if err != nil {
		return
	}
	help = string(helpAux)
	return
}
