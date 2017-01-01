package goConfig

import "testing"

type testAux struct {
	a int `config:"AUX"`
}

func TestParseTags(t *testing.T) {
	aux := testAux{a: 1}
	parseTags(aux)
}
