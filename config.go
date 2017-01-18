package goConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/crgimenes/goConfig/goEnv"
	"github.com/crgimenes/goConfig/goFlags"
)

const tag = "cfg"
const tagDefault = "cfgDefault"

// Path sets default config path
var Path string

// File name of default config file
var File string

// FileRequired config file required
var FileRequired bool

var HelpString string

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

	goEnv.Setup(tag, tagDefault)
	err = goEnv.Parse(config)
	if err != nil {
		return
	}

	prepareHelp(config)

	goFlags.Setup(tag, tagDefault)
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
		os.Mkdir(Path, 0700)
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
