package goFlags

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"flag"

	"github.com/crgimenes/goConfig/structTag"
)

type parameterMeta struct {
	Kind  reflect.Kind
	Value interface{}
	Tag   string
}

var parametersMetaMap map[*reflect.Value]parameterMeta
var visitedMap map[string]*flag.Flag

// Preserve disable default values and get only visited parameters thus preserving the values passed in the structure, default false
var Preserve bool

var Usage func()

// Setup maps and variables
func Setup(tag string, tagDefault string) {
	Usage = DefaultUsage

	parametersMetaMap = make(map[*reflect.Value]parameterMeta)
	visitedMap = make(map[string]*flag.Flag)

	structTag.Setup()
	SetTag(tag)
	SetTagDefault(tagDefault)

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

	flag.Visit(loadVisit)

	for k, v := range parametersMetaMap {
		if _, ok := visitedMap[v.Tag]; !ok && Preserve {
			continue
		}
		switch v.Kind {
		case reflect.String:
			value := *v.Value.(*string)
			//fmt.Printf("Parse %v = \"%v\"\n", v.Tag, value)
			k.SetString(value)
		case reflect.Int:
			value := *v.Value.(*int)
			//fmt.Printf("Parse %v = \"%v\"\n", v.Tag, value)
			k.SetInt(int64(value))

		}
	}

	return
}

// Reset maps caling setup function
func Reset() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.Usage = nil

	structTag.Reset()
	Setup(structTag.Tag, structTag.TagDefault)
}

func loadVisit(f *flag.Flag) {
	visitedMap[f.Name] = f
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

	meta := parameterMeta{}
	meta.Value = &aux
	meta.Tag = strings.ToLower(tag)
	meta.Kind = reflect.Int
	parametersMetaMap[value] = meta

	flag.IntVar(&aux, meta.Tag, defaltValueInt, "")

	//fmt.Println(tag, defaltValue)

	return
}

func reflectString(field *reflect.StructField, value *reflect.Value, tag string) (err error) {

	var aux string
	var defaltValue string
	defaltValue = field.Tag.Get(structTag.TagDefault)

	meta := parameterMeta{}
	meta.Value = &aux
	meta.Tag = strings.ToLower(tag)
	meta.Kind = reflect.String
	parametersMetaMap[value] = meta

	flag.StringVar(&aux, meta.Tag, defaltValue, "")

	//fmt.Println(tag, defaltValue)

	return
}

func PrintDefaults() {
	flag.PrintDefaults()

}

func DefaultUsage() {
	fmt.Println("Usage")
	PrintDefaults()
}
