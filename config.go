package goConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/pkg/errors"
)

const (
	defaultPath       = "./"
	defaultConfigFile = "config.json"
)

// Configuration struct
type Configuration struct {
	Name  string
	Value interface{}
}

// Config instantiate the system settings.
var Config = Configuration{}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Load config file
func (c *Configuration) Load() (err error) {
	configFile := defaultPath + defaultConfigFile
	file, err := os.Open(configFile)
	if err != nil {
		log.Println("Load open config.json:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Println("Load Decode:", err)
		return
	}
	return
}

// Save config file
func (c *Configuration) Save() (err error) {
	_, err = os.Stat(defaultPath)

	if os.IsNotExist(err) {
		os.Mkdir(defaultPath, 0700)
	}

	configFile := defaultPath + defaultConfigFile

	_, err = os.Stat(configFile)
	if err != nil {
		log.Println(err)
		return
	}

	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(defaultConfigFile, b, 0644)
	if err != nil {
		log.Println(err)
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
			/*
				for i2 := 0; i2 < value.NumField(); i2++ {
					fmt.Printf("%v %#v\n",
						value.Field(i2).Kind().String(),
						value.Field(i2).Interface())
				}
			*/
			err = parseTags(value.Addr().Interface())
			if err != nil {
				return
			}

		}

		fmt.Println("name:", field.Name,
			"| cfg:", field.Tag.Get("config"),
			"| cfgDefault:", field.Tag.Get("cfgDefault"),
			"| type:", field.Type, kindStr)

	}
	return
}
