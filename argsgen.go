package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/template"
)

//go:embed templates/*
var fs embed.FS

type flagSet struct {
	VarFunc     string
	Field       string
	Description string
	Default     string
}

type required struct {
	Field string
	Zero  string
}

type vars struct {
	StructName string
	Flags      []flagSet
	Positional map[int][]string
	Required   []required
	Aliases    map[string]string
}

type header struct {
	Package  string
	Packages map[string]struct{}
}

func parse(filename, pkg string, writer io.Writer) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFS(fs, "templates/*.tmpl")
	if err != nil {
		return err
	}

	var header header
	header.Packages = make(map[string]struct{})
	header.Packages["flag"] = struct{}{}
	header.Packages["os"] = struct{}{}
	header.Packages["fmt"] = struct{}{}
	header.Package = pkg

	var varsList []vars
	var vs vars

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			vs.StructName = x.Name.Name
		case *ast.StructType:
			// skip anonymous structs
			if vs.StructName == "" {
				return true
			}

			vs.Flags = []flagSet{}
			vs.Positional = make(map[int][]string)
			vs.Required = []required{}
			vs.Aliases = make(map[string]string)

			for _, field := range x.Fields.List {
				var varFunc string
				var ident *ast.Ident
				switch t := field.Type.(type) {
				case *ast.Ident:
					ident = t
				case *ast.SelectorExpr:
					pkg, ok := t.X.(*ast.Ident)
					if ok && pkg.Name == "time" {
						ident = t.Sel
						break
					}
					continue
				default:
					continue
				}

				switch name := ident.Name; name {
				case "uint", "uint64", "int", "int64", "float64", "string", "bool", "Duration":
					varFunc = fmt.Sprintf("%sVar", strings.Title(name))
				default:
					panic("type not supported " + name)
				}
				var tag string
				if field.Tag != nil {
					tagString, _ := strconv.Unquote(field.Tag.Value)
					tag = reflect.StructTag(tagString).Get(("arg"))
					options := strings.Split(tag, ",")
					for _, option := range options[1:] {
						switch option {
						case "+", "positional":
							tag = options[0]
							n := len(vs.Positional)
							for _, name := range field.Names {
								vs.Positional[n] = append(vs.Positional[n], name.String())
							}
						case "!", "required":
							tag = options[0]
							for _, name := range field.Names {
								var zero string
								switch ident.Name {
								case "uint", "uint64", "int", "int64", "float64":
									zero = "0"
								case "string":
									zero = `""`
								default:
									panic("type cannot be a required field: " + ident.Name)
								}
								vs.Required = append(vs.Required, required{
									Field: name.String(),
									Zero:  zero,
								})
								header.Packages["errors"] = struct{}{}
							}
						}
					}
				}
				for _, name := range field.Names {
					vs.Flags = append(vs.Flags, flagSet{
						VarFunc:     varFunc,
						Field:       name.String(),
						Description: tag,
						Default:     field.Names[0].String(),
					})
				}

				for _, alias := range field.Names[1:] {
					vs.Aliases[field.Names[0].String()] = alias.String()
				}
			}

			varsList = append(varsList, vs)
			vs = vars{}

			return false
		}
		return true
	})

	if err = tmpl.ExecuteTemplate(writer, "header.tmpl", header); err != nil {
		return err
	}

	for _, vs := range varsList {
		if err = tmpl.ExecuteTemplate(writer, "parse.tmpl", vs); err != nil {
			return err
		}
	}

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
