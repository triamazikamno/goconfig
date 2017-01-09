package goFlags

import (
	"fmt"
	"os"
	"testing"
)

type testStruct struct {
	A int    `flag:"A" flagDefault:"100"`
	B string `flag:"B" flagDefault:"200"`
	C string
	N string `flag:"-"`
	p string
	S testSub `flag:"S"`
}

type testSub struct {
	A int        `flag:"A" flagDefault:"300"`
	B string     `flag:"C" flagDefault:"400"`
	S testSubSub `flag:"S"`
}
type testSubSub struct {
	A int    `flag:"A" flagDefault:"500"`
	B string `flag:"S" flagDefault:"600"`
}

func TestParse(t *testing.T) {

	//os.Args = []string{"noop", "-flag1=val1", "arg1", "arg2"}
	//os.Args = []string{"program", "-h"}

	os.Args = []string{
		"program",
		"-a=8888",
		"-b=TEST",
		"-s_s_a=9999",
	}

	s := &testStruct{A: 1, S: testSub{A: 1, B: "2"}}

	err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("\n\nTestParseTags: %#v\n\n", s)

}
