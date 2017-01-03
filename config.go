package goConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/pkg/errors"
)

const (
	defaultPath       = "./"
	defaultConfigFile = "config.json"
)

// Load config file
func Load(config interface{}) (err error) {
	configFile := defaultPath + defaultConfigFile
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

// Save config file
func Save(config interface{}) (err error) {
	_, err = os.Stat(defaultPath)

	if os.IsNotExist(err) {
		os.Mkdir(defaultPath, 0700)
	}

	configFile := defaultPath + defaultConfigFile

	_, err = os.Stat(configFile)
	if err != nil {
		return
	}

	b, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return
	}

	err = ioutil.WriteFile(defaultConfigFile, b, 0644)
	if err != nil {
		return
	}
	return
}

// Getenv get enviroment variable
func Getenv(env string) (r string) {
	r = os.Getenv(env)
	return
}

func parseTags(s interface{}) (err error) {

	st := reflect.TypeOf(s)
	vt := reflect.ValueOf(s)

	if st.Kind() != reflect.Ptr {
		err = errors.New("Not a pointer")
		return
	}

	ref := st.Elem()
	if ref.Kind() != reflect.Struct {
		err = errors.New("Not a struct")
		return
	}

	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)

		kindStr := ""
		if field.Type.Kind() == reflect.Struct {
			kindStr = "Struct"

			value := vt.Elem().Field(i)
			err = parseTags(value.Addr().Interface())
			if err != nil {
				return
			}

		}

		// change value POC
		if field.Type.Kind() == reflect.String {
			value := vt.Elem().Field(i)
			value.SetString("TEST")
		}

		fmt.Println("name:", field.Name,
			"| cfg:", field.Tag.Get("config"),
			"| cfgDefault:", field.Tag.Get("cfgDefault"),
			"| type:", field.Type, kindStr)

	}
	return
}
