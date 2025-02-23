package main

import (
	"fmt"
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
				t.Fatalf("\nparse:%d():\n\twant:%q\n\tgot: %q", i, want[i], got[i])
			}
		}
	}
}

func Test_parse(t *testing.T) {
	testcases := []string{
		"simple",
		"positional",
		"required",
		"required_and_positional",
		"positional_and_required",
	}

	for _, tt := range testcases {
		t.Run(tt, func(t *testing.T) {
			compare(t, fmt.Sprintf("testdata/%s.go", tt), fmt.Sprintf("testdata/%s_gen.go", tt))
		})
	}
}
