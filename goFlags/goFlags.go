package goFlags

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"flag"

	"github.com/crgimenes/goConfig/structTag"
)

var parametersStringMap map[*reflect.Value]*string
var parametersIntMap map[*reflect.Value]*int

// Preserve disable default values and get only visited parameters thus preserving the values passed in the structure, default false
var Preserve bool

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

// Parse configuration
func Parse(config interface{}) (err error) {

	err = structTag.Parse(config, "")
	if err != nil {
		return
	}

	flag.Parse()

	//fmt.Printf("%v#f", flag.CommandLine)
	flag.Visit(visitTest)

	for k, v := range parametersStringMap {
		fmt.Printf("- \"%v\"\n", *v)
		k.SetString(*v)
	}

	for k, v := range parametersIntMap {
		fmt.Printf("- \"%v\"\n", int64(*v))
		k.SetInt(int64(*v))

	}
	return
}

func visitTest(f *flag.Flag) {
	fmt.Printf("name \"%v\"\n", f.Name)
}

func reflectInt(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	var aux int
	var defaltValue string
	var defaltValueInt int

	defaltValue = field.Tag.Get(structTag.TagDefault)

	if defaltValue == "" || defaltValue == "0" {
		defaltValueInt = 0
	} else {
		defaltValueInt, err = strconv.Atoi(defaltValue)
		if err != nil {
			return
		}
	}

	flag.IntVar(&aux, strings.ToLower(tag), defaltValueInt, "")
	parametersIntMap[value] = &aux

	fmt.Println(tag, defaltValue)

	return
}

func reflectString(field *reflect.StructField, value *reflect.Value, tag string) (err error) {

	var aux string
	var defaltValue string
	defaltValue = field.Tag.Get(structTag.TagDefault)

	flag.StringVar(&aux, strings.ToLower(tag), defaltValue, "")
	parametersStringMap[value] = &aux

	fmt.Println(tag, defaltValue)

	return
}
