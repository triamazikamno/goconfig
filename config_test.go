package goConfig

import "testing"

type testAux struct {
	a int    `config:"field a"`
	b string `config:"field b"`
}

func TestParseTags(t *testing.T) {
	aux := testAux{a: 1}
	parseTags(aux)
}
