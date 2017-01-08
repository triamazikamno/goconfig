package goFlags

import (
	"fmt"
	"reflect"
	"strings"

	"flag"

	"github.com/crgimenes/goConfig/structTag"
)

var parametersStringMap map[*reflect.Value]*string
var parametersIntMap map[*reflect.Value]*int

func init() {
	parametersStringMap = make(map[*reflect.Value]*string)
	parametersIntMap = make(map[*reflect.Value]*int)

	SetTag("flag")
	SetTagDefault("flagDefault")

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
	if err != nil {
		return
	}

	flag.Parse()

	for k, v := range parametersStringMap {
		k.SetString(*v)
	}

	for k, v := range parametersIntMap {
		k.SetInt(int64(*v))
	}
	return
}

func reflectInt(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	fmt.Println(tag)

	var aux int

	flag.IntVar(&aux, strings.ToLower(tag), 0, "")
	parametersIntMap[value] = &aux

	return
}

func reflectString(field *reflect.StructField, value *reflect.Value, tag string) (err error) {

	// get value from default settings
	//ret = field.Tag.Get(structTag.TagDefault)
	fmt.Println(tag)

	var aux string

	flag.StringVar(&aux, strings.ToLower(tag), "", "")
	parametersStringMap[value] = &aux

	return
}
