package goFlags

import (
	"fmt"
	"os"
	"testing"
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

func TestParse(t *testing.T) {

	//os.Args = []string{"noop", "-flag1=val1", "arg1", "arg2"}
	//os.Args = []string{"program", "-h"}

	os.Args = []string{
		"program",
		"-a=100",
		"-b=TEST",
		"-s_s_a=600",
	}

	s := &testStruct{A: 1, S: testSub{A: 1, B: "2"}}
	err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("\n\nTestParseTags: %#v\n\n", s)

}
