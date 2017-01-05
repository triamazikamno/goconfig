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
	// FileRequired config file required
	FileRequired bool
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

var parseMap map[reflect.Kind]func(
	field *reflect.StructField,
	value *reflect.Value,
	tag string) (err error)

func init() {
	Setup = &Settings{
		Path:                    "./",
		File:                    "config.json",
		Tag:                     "cfg",
		TagDefault:              "cfgDefault",
		TagDisabled:             "-",
		EnvironmentVarSeparator: "_",
		FileRequired:            false,
	}

	parseMap = make(map[reflect.Kind]func(
		field *reflect.StructField,
		value *reflect.Value, tag string) (err error))

	parseMap[reflect.Struct] = reflectStruct
	parseMap[reflect.Int] = reflectInt
	parseMap[reflect.String] = reflectString

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

		t := updateTag(&field, superTag)
		if t == "" {
			continue
		}

		if f, ok := parseMap[kind]; ok {
			err = f(&field, &value, t)
			if err != nil {
				return
			}
		} else {
			err = errors.New("Type not supported " + kind.String())
			return
		}

		fmt.Println("name:", field.Name,
			"| cfg:", field.Tag.Get(Setup.Tag),
			"| cfgDefault:", field.Tag.Get(Setup.TagDefault),
			"| type:", field.Type)

	}
	return
}

func updateTag(field *reflect.StructField, superTag string) (ret string) {
	ret = field.Tag.Get(Setup.Tag)
	if ret == Setup.TagDisabled {
		return
	}

	if ret == "" {
		ret = strings.ToUpper(field.Name)
	}

	if superTag != "" {
		ret = superTag + Setup.EnvironmentVarSeparator + ret
	}
	return
}

func getNewValue(field *reflect.StructField, tag string) (ret string) {

	ret = os.Getenv(tag)
	if ret != "" {
		return
	}

	ret = field.Tag.Get(Setup.TagDefault)
	if ret != "" {
		return
	}

	return

}

func reflectStruct(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	err = parseTags(value.Addr().Interface(), tag)
	return
}

func reflectInt(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	//value.SetInt(999)

	newValue := getNewValue(field, tag)

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
	newValue := getNewValue(field, tag)

	value.SetString(newValue)

	return
}
