package goConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Settings default
type Settings struct {
	// Path sets default config path
	Path string
	// File name of default config file
	File string
	// Tag set the main tag
	Tag string
	// TagDefault set tag default
	TagDefault string
	// TagDisabled used to not process an input
	TagDisabled string
	// EnvironmentVarSeparator separe names on environment variables
	EnvironmentVarSeparator string
}

// Setup Pointer to internal variables
var Setup *Settings

func init() {
	Setup = &Settings{
		Path:                    "./",
		File:                    "config.json",
		Tag:                     "cfg",
		TagDefault:              "cfgDefault",
		TagDisabled:             "-",
		EnvironmentVarSeparator: "_",
	}
}

// LoadJSON config file
func LoadJSON(config interface{}) (err error) {
	configFile := Setup.Path + Setup.File
	file, err := os.Open(configFile)
	if err != nil {
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

	err = parseTags(config, "")
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

func parseTags(s interface{}, superTag string) (err error) {

	st := reflect.TypeOf(s)

	if st.Kind() != reflect.Ptr {
		err = errors.New("Not a pointer")
		return
	}

	refField := st.Elem()
	if refField.Kind() != reflect.Struct {
		err = errors.New("Not a struct")
		return
	}

	//vt := reflect.ValueOf(s)
	refValue := reflect.ValueOf(s).Elem()
	for i := 0; i < refField.NumField(); i++ {
		field := refField.Field(i)
		value := refValue.Field(i)
		kind := field.Type.Kind()

		if field.PkgPath != "" {
			continue
		}

		t := field.Tag.Get(Setup.Tag)
		if t == Setup.TagDisabled {
			continue
		}

		if t == "" {
			t = strings.ToUpper(field.Name)
		}

		if superTag != "" {
			t = superTag + Setup.EnvironmentVarSeparator + t
		}

		env := os.Getenv(t)

		if env == "" && kind != reflect.Struct {
			continue
		}

		switch kind {
		case reflect.Struct:
			err = parseTags(value.Addr().Interface(), t)
			if err != nil {
				return
			}
		case reflect.String:
			//value.SetString("TEST")
			value.SetString(env)
		case reflect.Int:
			//value.SetInt(999)
			var intEnv int64
			intEnv, err = strconv.ParseInt(env, 10, 64)
			if err != nil {
				return
			}
			value.SetInt(intEnv)
		default:
			err = errors.New("Type not supported " + kind.String())
		}

		fmt.Println("name:", field.Name,
			"| cfg:", field.Tag.Get(Setup.Tag),
			"| cfgDefault:", field.Tag.Get(Setup.TagDefault),
			"| type:", field.Type)

	}
	return
}
