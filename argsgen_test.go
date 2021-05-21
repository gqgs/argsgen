package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func Test_parse(t *testing.T) {
	builder := new(strings.Builder)
	err := parse("testdata/options.go", "options", builder)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := ioutil.ReadFile("testdata/options_gen.go")
	if err != nil {
		t.Fatal(err)
	}
	output := builder.String()
	if cmp := strings.Compare(string(expected), output); cmp != 0 {
		t.Errorf("parse():\nwant: %q\ngot: %q", expected, output)
	}
}
