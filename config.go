package goConfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/crgimenes/goConfig/structTag"
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

	structTag.ParseMap[reflect.Int] = reflectInt
	structTag.ParseMap[reflect.String] = reflectString

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

	err = structTag.Parse(config, "")
	if err != nil {
		return
	}

	postProc()

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

func getNewValue(field *reflect.StructField, value *reflect.Value, tag string) (ret string) {

	//TODO: get value from parameter.

	// get value from environment variable
	ret = os.Getenv(strings.ToUpper(tag))
	if ret != "" {
		return
	}

	// get value from config file
	switch value.Kind() {
	case reflect.String:
		ret = value.String()
		return
	case reflect.Int:
		ret = strconv.FormatInt(value.Int(), 10)
		return
	}

	// get value from default settings
	ret = field.Tag.Get(structTag.TagDefault)

	return
}

func postProc() {
}

func reflectInt(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	//value.SetInt(999)

	newValue := getNewValue(field, value, tag)

	var intNewValue int64
	intNewValue, err = strconv.ParseInt(newValue, 10, 64)
	if err != nil {
		return
	}
	value.SetInt(intNewValue)

	return
}

func reflectString(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	//value.SetString("TEST")

	newValue := getNewValue(field, value, tag)

	value.SetString(newValue)

	return
}
