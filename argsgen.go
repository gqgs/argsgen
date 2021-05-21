package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//go:generate go run $GOFILE

type options struct {
	i, input     string `arg:"input filename,+"`
	o, output    string `arg:"output filename,+"`
	db, database string `arg:"database name"`
	folder       string `arg:"target folder"`
	parallel     uint   `arg:"number of process in parallel"`
	profile      bool
}

type flagSet struct {
	VarFunc     string
	Field       string
	Description string
	Default     string
}

type vars struct {
	Package    string
	StructName string
	Flags      []flagSet
	Positional map[int][]string
}

func parse(filename, pkg string, writer io.Writer) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseGlob("./templates/*.tmpl")
	if err != nil {
		return err
	}

	vars := vars{
		Package: pkg,
	}

	if err = tmpl.ExecuteTemplate(writer, "header.tmpl", vars); err != nil {
		return err
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			vars.StructName = x.Name.Name
		case *ast.StructType:
			vars.Flags = []flagSet{}
			vars.Positional = make(map[int][]string)

			for _, field := range x.Fields.List {
				var varFunc string
				t, ok := field.Type.(*ast.Ident)
				if !ok {
					continue
				}

				switch name := t.Name; name {
				case "string", "bool", "uint", "int":
					varFunc = fmt.Sprintf("%sVar", strings.Title(name))
				default:
					panic("type not supported " + name)
				}
				var tag string
				if field.Tag != nil {
					tagString, _ := strconv.Unquote(field.Tag.Value)
					tag = reflect.StructTag(tagString).Get(("arg"))
					options := strings.Split(tag, ",")
					if len(options) > 1 && options[1] == "+" {
						tag = options[0]
						n := len(vars.Positional)
						for _, name := range field.Names {
							vars.Positional[n] = append(vars.Positional[n], name.String())
						}
					}
				}
				for _, name := range field.Names {
					vars.Flags = append(vars.Flags, flagSet{
						VarFunc:     varFunc,
						Field:       name.String(),
						Description: tag,
						Default:     field.Names[0].String(),
					})
				}
			}

			if err = tmpl.ExecuteTemplate(writer, "parse.tmpl", vars); err != nil {
				panic(err)
			}

			return false
		}
		return true
	})

	return nil
}

func main() {
	input := os.Getenv("GOFILE")
	pkg := os.Getenv("GOPACKAGE")
	writer := new(bytes.Buffer)

	if err := parse(input, pkg, writer); err != nil {
		log.Fatal(err)
	}

	output := strings.TrimSuffix(input, ".go") + "_gen.go"
	file, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	out := writer.String()
	file.WriteString(out)
}
