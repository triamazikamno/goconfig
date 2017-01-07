package getEnv

import (
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/crgimenes/goConfig/structTag"
)

func init() {
	SetTag("env")
	SetTagDefault("envDefault")

	structTag.ParseMap[reflect.Int] = reflectInt
	structTag.ParseMap[reflect.String] = reflectString
}

// SetTag set a new tag
func SetTag(tag string) {
	structTag.Tag = tag
}

// SetTagDefault set a new TagDefault to retorn default values
func SetTagDefault(tag string) {
	structTag.TagDefault = tag
}

// Parse config file
func Parse(config interface{}) (err error) {
	err = structTag.Parse(config, "")
	return
}

func getNewValue(field *reflect.StructField, value *reflect.Value, tag string) (ret string) {

	// get value from environment variable
	ret = os.Getenv(strings.ToUpper(tag))
	if ret != "" {
		return
	}

	// get value from default settings
	ret = field.Tag.Get(structTag.TagDefault)

	return
}

func reflectInt(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	newValue := getNewValue(field, value, tag)
	if newValue == "" {
		return
	}

	var intNewValue int64
	intNewValue, err = strconv.ParseInt(newValue, 10, 64)
	if err != nil {
		return
	}

	value.SetInt(intNewValue)

	return
}

func reflectString(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	newValue := getNewValue(field, value, tag)

	value.SetString(newValue)

	return
}
