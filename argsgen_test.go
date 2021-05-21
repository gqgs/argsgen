package main

import (
	"io/ioutil"
	"regexp"
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
	spaceRemover := regexp.MustCompile("\\s+")
	output := spaceRemover.ReplaceAllString(builder.String(), "\n")
	want := spaceRemover.ReplaceAllString(string(expected), "\n")

	if cmp := strings.Compare(want, output); cmp != 0 {
		t.Errorf("parse():\n%q\n\n\n%q", want, output)
	}
}
