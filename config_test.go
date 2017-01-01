package goConfig

import "testing"

type testSub1 struct {
	S1 int
	S2 int
	S3 string
}
type testAux struct {
	A int      `config:"field a"`
	B string   `config:"field b"`
	S testSub1 `config:"Sub 1"`
}

func TestParseTags(t *testing.T) {
	s := &testAux{A: 1, S: testSub1{S1: 1, S2: 2, S3: "test"}}
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
