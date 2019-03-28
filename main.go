package main

//go:generate go-bindata out.tmpl

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/jessevdk/go-flags"
)

type generator struct {
	Package   string   `short:"p" long:"pkg" description:"name of the package in which the generate file will live" default:"main"`
	Output    string   `short:"o" long:"output" description:"name of the output file" default:"gopherplate.go"`
	Engine    string   `short:"e" long:"engine" description:"name of the sql engine" default:"mysql"`
	Source    string   `short:"s" long:"source" description:"pattern to use to get source files" default:"."`
	TableName []string `short:"t" long:"table-name" description:"overrides the default table name for a struct: StructName:table_name"`
	Named     bool     `long:"named" description:"use named arguments in the generated sql"`
}

func main() {
	var g generator
	if _, err := flag.Parse(&g); err != nil {
		log.Fatalf("could not parse flags %+v", err)
	}

	srcs, err := sourceFilePaths(g.Source)
	if err != nil {
		return
	}

	var sms []*structModel
	for _, src := range srcs {
		sm, err := g.structsInPath(src)
		if err != nil {
			log.Fatal(err)
		}

		if len(sm) != 0 {
			sms = append(sms, sm...)
		}
	}

	input := templateInput{
		Structs:     sms,
		Insert:      true,
		Update:      true,
		Join:        true,
		PackageName: g.Package,
	}

	if err := g.generate(&input); err != nil {
		log.Fatal(err)
	}
}

func sourceFilePaths(path string) ([]string, error) {
	var paths []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		paths = append(paths, path)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}
