package goConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/crgimenes/goConfig/goEnv"
	"github.com/crgimenes/goConfig/goFlags"
)

// Tag to set main name of field
var Tag = "cfg"

// TagDefault to set default value
var TagDefault = "cfgDefault"

// Path sets default config path
var Path string

// File name of default config file
var File string

// FileRequired config file required
var FileRequired bool

// HelpString temporarily saves help
var HelpString string

//Usage is a function to show the help, can be replaced by your own version.
var Usage func()

func init() {
	Usage = DefaultUsage
	Path = "./"
	File = "config.json"
	FileRequired = false
}

// Parse configuration
func Parse(config interface{}) (err error) {

	err = LoadJSON(config)
	if err != nil {
		return
	}

	goEnv.Setup(Tag, TagDefault)
	err = goEnv.Parse(config)
	if err != nil {
		return
	}

	prepareHelp(config)

	goFlags.Setup(Tag, TagDefault)
	goFlags.Usage = Usage
	goFlags.Preserve = true
	err = goFlags.Parse(config)
	if err != nil {
		return
	}

	return
}

// LoadJSON config file
func LoadJSON(config interface{}) (err error) {
	configFile := Path + File
	file, err := os.Open(configFile)
	if os.IsNotExist(err) && !FileRequired {
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

// Save config file
func Save(config interface{}) (err error) {
	_, err = os.Stat(Path)
	if os.IsNotExist(err) {
		err = os.Mkdir(Path, os.ModePerm)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	}

	configFile := Path + File

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

func prepareHelp(config interface{}) (err error) {
	var helpAux []byte
	helpAux, err = json.MarshalIndent(config, "", "    ")
	if err != nil {
		return
	}
	HelpString = string(helpAux)
	return
}

func PrintDefaults() {
	fmt.Printf("Config file %q:\n", Path+File)
	fmt.Println(HelpString)
}

func DefaultUsage() {
	fmt.Println("Usage")
	goFlags.PrintDefaults()
	goEnv.PrintDefaults()
	PrintDefaults()
}
