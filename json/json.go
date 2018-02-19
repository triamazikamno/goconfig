package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/crgimenes/goconfig"
)

func init() {
	f := goconfig.Fileformat{
		Extension:   ".json",
		Save:        SaveJSON,
		Load:        LoadJSON,
		PrepareHelp: PrepareHelp,
	}
	goconfig.Formats = append(goconfig.Formats, f)
}

// LoadJSON config file
func LoadJSON(config interface{}) (err error) {
	configFile := filepath.Join(goconfig.Path, goconfig.File)
	file, err := os.Open(configFile)
	if os.IsNotExist(err) && !goconfig.FileRequired {
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
	_, err = os.Stat(goconfig.Path)
	if os.IsNotExist(err) {
		err = os.Mkdir(goconfig.Path, os.ModePerm)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	}

	configFile := filepath.Join(goconfig.Path, goconfig.File)

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
