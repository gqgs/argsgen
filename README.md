# argsgen
CLI argument parser code generator


#### Usage:

```go
package options

//go:generate go run github.com/gqgs/argsgen

type options struct {
	i, input     string `arg:"input filename,+"`  // 1st positional argument
	o, output    string `arg:"output filename,+"` // 2nd positional argument
	db, database string `arg:"database name"`
	folder       string `arg:"target folder"`
	parallel     uint   `arg:"number of process in parallel"`
	profile      bool
}
```

:point_down:

```go
// Code generated by argsgen.
// DO NOT EDIT!
package options

import (
	"flag"
	"os"
)

func (o *options) Parse() error {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagSet.StringVar(&o.i, "i", o.i, "input filename")
	flagSet.StringVar(&o.i, "input", o.i, "input filename")
	flagSet.StringVar(&o.o, "o", o.o, "output filename")
	flagSet.StringVar(&o.o, "output", o.o, "output filename")
	flagSet.StringVar(&o.db, "db", o.db, "database name")
	flagSet.StringVar(&o.db, "database", o.db, "database name")
	flagSet.StringVar(&o.folder, "folder", o.folder, "target folder")
	flagSet.UintVar(&o.parallel, "parallel", o.parallel, "number of process in parallel")
	flagSet.BoolVar(&o.profile, "profile", o.profile, "")

	var positional []string
	args := os.Args[1:]
	for len(args) > 0 {
		if err := flagSet.Parse(args); err != nil {
			return err
		}

		if remaining := flagSet.NArg(); remaining > 0 {
			posIndex := len(args) - remaining
			positional = append(positional, args[posIndex])
			args = args[posIndex+1:]
			continue
		}
		break
	}

	if len(positional) == 0 {
		return nil
	}

	o.i = positional[0]
	o.input = positional[0]

	o.o = positional[1]
	o.output = positional[1]

	return nil
}

func (o *options) MustParse() {
	if err := o.Parse(); err != nil {
		panic(err)
	}
}
```
