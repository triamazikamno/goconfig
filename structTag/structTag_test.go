package structTag

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type testStruct struct {
	A int     `cfg:"A" cfgDefault:"100"`
	B string  `cfg:"B" cfgDefault:"200"`
	S testSub `cfg:"S"`
}

type testSub struct {
	A int        `cfg:"A" cfgDefault:"300"`
	B int        `cfg:"B" cfgDefault:"400"`
	C string     `cfg:"C" cfgDefault:"500"`
	S testSubSub `cfg:"S"`
}
type testSubSub struct {
	A int    `cfg:"A" cfgDefault:"600"`
	B int    `cfg:"B" cfgDefault:"700"`
	C string `cfg:"S" cfgDefault:"900"`
}

func reflectIntTestFunc(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	newValue := field.Tag.Get(TagDefault)

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

func reflectStringTestFunc(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	newValue := field.Tag.Get(TagDefault)
	value.SetString(newValue)

	return
}

func TestParse(t *testing.T) {
	Tag = "cfg"
	TagDefault = "cfgDefault"

	s := &testStruct{A: 1, S: testSub{A: 1, B: 2, C: "test"}}

	err := Parse(s, "")
	if err != ErrTypeNotSupported {
		t.Fatal("ErrTypeNotSupported error expected")
	}

	ParseMap[reflect.Int] = reflectIntTestFunc
	ParseMap[reflect.String] = reflectStringTestFunc
	err = Parse(s, "")
	if err != nil {
		t.Fatal("teste", err)
	}

	fmt.Printf("\n\nTestParseTags: %#v\n\n", s)

	s1 := "test"
	err = Parse(s1, "")
	if err != ErrNotAPointer {
		t.Fatal("ErrNotAPointer error expected")
	}

	err = Parse(&s1, "")
	if err != ErrNotAStruct {
		t.Fatal("ErrNotAStruct error expected")
	}
}
