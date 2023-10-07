package main

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

var (
	spaceRemover = regexp.MustCompile(`\s+`)
)

func compare(t *testing.T, inputFile, outputFile string) {
	builder := new(strings.Builder)
	err := parse(inputFile, "options", builder)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatal(err)
	}

	got := spaceRemover.ReplaceAllString(builder.String(), "\n")
	want := spaceRemover.ReplaceAllString(string(expected), "\n")

	if cmp := strings.Compare(want, got); cmp != 0 {
		want := strings.Split(want, "\n")
		got := strings.Split(got, "\n")

		for i := 0; i < len(want) && i < len(got); i++ {
			if strings.Compare(want[i], got[i]) != 0 {
				t.Fatalf("\nparse.%d():\n\twant:%q\n\tgot: %q", i, want[i], got[i])
			}
		}
	}
}

func Test_parse_options(t *testing.T) {
	compare(t, "testdata/options.go", "testdata/options_gen.go")
}

func Test_parse_required(t *testing.T) {
	compare(t, "testdata/required.go", "testdata/required_gen.go")
}
