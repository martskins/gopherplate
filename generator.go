package main

import (
	"bytes"
	"go/format"
	"os"
	"strconv"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"
)

func (g *generator) generate(params *templateInput) error {
	ass, err := Asset("out.tmpl")
	if err != nil {
		return err
	}

	funcs := template.FuncMap{
		"Contains": strings.Contains,
		"Dec":      func(i int) int { return i - 1 },
		"SanitizeType": func(str string) string {
			return strings.Replace(str, "*", "", -1)
		},
	}

	tmp, err := template.New("auto-generated").Funcs(funcs).Parse(string(ass))
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer([]byte{})
	if err := tmp.Execute(buf, params); err != nil {
		return err
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	bts, err := imports.Process("", p, nil)
	if err != nil {
		return err
	}

	fl, err := os.Create(g.Output)
	if err != nil {
		return err
	}
	defer fl.Close()

	if _, err := fl.Write(bts); err != nil {
		return err
	}

	return nil
}

// MySQL uses the ? variant shown above
// PostgreSQL uses an enumerated $1, $2, etc bindvar syntax
// SQLite accepts both ? and $1 syntax
// Oracle uses a :name syntax
func (g *generator) placeholder(tagValue string, idx int) string {
	if g.Named {
		return ":" + tagValue
	}

	switch g.Engine {
	case "mysql":
		return "?"
	case "postgres":
		return "$" + strconv.Itoa(idx+1)
	case "sqlite":
		return "?"
	case "oracle":
		return ":" + tagValue
	default:
		return "?"
	}
}
