package goConfig

import "testing"

type testSub struct {
	S1 int        `config:"field S1" cfgDefault:"1"`
	S2 int        `config:"field S2"`
	S3 string     `config:"field S3"`
	S4 testSubSub `config:"Sub Sub"`
}
type testSubSub struct {
	SS1 int    `config:"field SS1" cfgDefault:"2"`
	SS2 int    `config:"field SS2"`
	SS3 string `config:"field SS3"`
}
type testAux struct {
	A int     `config:"field a"`
	B string  `config:"field b"`
	S testSub `config:"Sub"`
}

func TestParseTags(t *testing.T) {
	s := &testAux{A: 1, S: testSub{S1: 1, S2: 2, S3: "test"}}
	err := parseTags(s)
	if err != nil {
		t.Fatal(err)
	}

	s1 := "test"
	err = parseTags(s1)
	if err == nil {
		t.Fatal("Error expected")
	}

	err = parseTags(&s1)
	if err == nil {
		t.Fatal("Error expected")
	}

}
