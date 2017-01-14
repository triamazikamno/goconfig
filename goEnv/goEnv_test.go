package goEnv

import (
	"os"
	"testing"
)

type testStruct struct {
	A int    `cfg:"A" cfgDefault:"100"`
	B string `cfg:"B" cfgDefault:"200"`
	C string
	N string `cfg:"-"`
	M int
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

	Setup("cfg", "cfgDefault")

	os.Setenv("A", "900")
	os.Setenv("B", "TEST")

	s := &testStruct{A: 1, S: testSub{A: 1, B: "2"}}
	err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	if s.A != 900 {
		t.Fatal("s.A != 900, s.A:", s.A)
	}

	if s.B != "TEST" {
		t.Fatal("s.B != \"TEST\", s.B:", s.B)
	}

	if s.S.S.B != "600" {
		t.Fatal("s.S.S.B != \"600\", s.S.S.B:", s.S.S.B)
	}

	os.Setenv("A", "900ERROR")

	err = Parse(s)
	if err == nil {
		t.Fatal("Error expected")
	}

	os.Setenv("A", "")

	err = Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	if s.A != 100 {
		t.Fatal("s.A != 100, s.A:", s.A)
	}

	s1 := "test"
	err = Parse(s1)
	if err == nil {
		t.Fatal("Error expected")
	}

	err = Parse(&s1)
	if err == nil {
		t.Fatal("Error expected")
	}
}
