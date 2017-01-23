package goConfig

import (
	"fmt"
	"os"
	"testing"

	"github.com/crgimenes/goConfig/goFlags"
	"github.com/crgimenes/goConfig/structTag"
)

type testStruct struct {
	A int    `cfg:"A" cfgDefault:"100"`
	B string `cfg:"B" cfgDefault:"200"`
	C string
	N string `cfg:"-"`
	p string
	S testSub `cfg:"S"`
}

type testSub struct {
	A int        `cfg:"A" cfgDefault:"300"`
	B string     `cfg:"C" cfgDefault:"400"`
	S testSubSub `cfg:"S"`
}
type testSubSub struct {
	A int    `cfg:"A" cfgDefault:"500"`
	B string `cfg:"S" cfgDefault:"600"`
}

func TestFindFileFormat(t *testing.T) {
	_, err := findFileFormat(".json")
	if err != ErrFileFormatNotDefined {
		t.Fatal(err)
	}
	Formats = []Fileformat{Fileformat{Extension: ".json"}}
	_, err = findFileFormat(".json")
	if err != nil {
		t.Fatal(err)
	}
}

// -=-=-=-=-=-=-=-=-=

func mLoad(config interface{}) (err error) {
	return
}

func mSave(config interface{}) (err error) {
	return
}

func mPrepareHelp(config interface{}) (help string, err error) {
	return
}

// -=-=-=-=-=-=-=-=-

func TestParse(t *testing.T) {

	s := &testStruct{A: 1, S: testSub{A: 1, B: "2"}}
	File = "config.txt"

	Formats = []Fileformat{Fileformat{Extension: ".json", Save: mSave, Load: mLoad, PrepareHelp: mPrepareHelp}}

	err := Parse(s)
	if err != ErrFileFormatNotDefined {
		t.Fatal("Error ErrFileFormatNotDefined expected")
	}

	File = "config.json"

	os.Setenv("A", "900")
	os.Setenv("B", "TEST")

	Tag = ""
	err = Parse(s)
	if err != structTag.ErrUndefinedTag {
		t.Fatal("Error structTag.ErrUndefinedTag expected")
	}

	Tag = "cfg"
	err = Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("\n\nTestParseTags: %#v\n\n", s)

	os.Setenv("A", "900ERROR")

	goFlags.Reset()
	err = Parse(s)
	if err == nil {
		t.Fatal("Error expected")
	}

	os.Setenv("A", "")

	goFlags.Reset()
	err = Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	s1 := "test"
	goFlags.Reset()
	err = Parse(s1)
	if err == nil {
		t.Fatal("Error expected")
	}

	goFlags.Reset()
	err = Parse(&s1)
	if err == nil {
		t.Fatal("Error expected")
	}

}
