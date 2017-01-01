package goConfig

import "testing"

type testAux struct {
	a int    `config:"field a"`
	b string `config:"field b"`
}

func TestParseTags(t *testing.T) {
	s := &testAux{a: 1}
	err := parseTags(s)
	if err != nil {
		t.Fatal(err)
	}
}
