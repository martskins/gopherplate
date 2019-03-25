package main

//go:generate go-bindata out.tmpl

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	flag "github.com/jessevdk/go-flags"
	"golang.org/x/tools/imports"
)

type generator struct {
	Pkg       string   `short:"p" long:"pkg" description:"name of the package in which the generate file will live" default:"main"`
	Output    string   `short:"o" long:"output" description:"name of the output file" default:"gopherplate.go"`
	Engine    string   `short:"e" long:"engine" description:"name of the sql engine" default:"mysql"`
	Source    string   `short:"s" long:"source" description:"pattern to use to get source files" default:"."`
	TableName []string `short:"t" long:"table-name" description:"overrides the default table name for a struct: StructName:table_name"`
	Named     bool     `long:"named" description:"use named arguments in the generated sql"`
}

type updateParams struct {
	Setters string
}

type joinParams struct {
	Fields string
}

type insertParams struct {
	PlaceHolders string
	Fields       string
}

type field struct {
	Name       string
	Type       string
	ColumnName string
	IsRelation bool
	Select     bool
	Insert     bool
	Update     bool
	Tags       *string
}

type templateInput struct {
	Structs     []*structModel
	Insert      bool
	Update      bool
	Join        bool
	PackageName string
}

type structModel struct {
	TableName string
	Name      string
	Package   string
	Insert    *insertParams
	Update    *updateParams
	Join      *joinParams
	Fields    []field
}

func main() {
	var g generator
	if _, err := flag.Parse(&g); err != nil {
		log.Fatalf("could not parse flags %+v", err)
	}

	srcs, err := sourceFiles(g.Source)
	if err != nil {
		return
	}

	var sms []*structModel
	for _, src := range srcs {
		sm, err := g.parseAndGenerate(src)
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
		PackageName: g.Pkg,
	}

	if err := g.generate(&input); err != nil {
		log.Fatal(err)
	}
}

func sourceFiles(path string) ([]string, error) {
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

func (g *generator) parseAndGenerate(path string) ([]*structModel, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	ff, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bts, err := ioutil.ReadAll(ff)
	if err != nil {
		return nil, err
	}

	var sms []*structModel
	ast.Inspect(file, func(n ast.Node) bool {
		x, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		y, ok := x.Type.(*ast.StructType)
		if !ok {
			return false
		}

		source := string(bts)
		var name string
		var end = y.Pos() - 2
		for {
			end -= 1
			if string(source[end]) != " " {
				continue
			}
			name = strings.Trim(source[end:y.Pos()-2], " ")
			break
		}

		sm, err := g.maybeParse(y.Fields.List, string(bts), name)
		if err != nil {
			return false
		}

		if sm != nil {
			sms = append(sms, sm)
		}
		return false
	})

	return sms, nil
}

func (g *generator) getTokenBetween(source string, start, end token.Pos) string {
	typeInSource := source[start-1 : end-1]
	return typeInSource
}

func (g *generator) maybeParse(fields []*ast.Field, source string, structName string) (*structModel, error) {
	structFields := []field{}
	for _, f := range fields {
		if len(f.Names) == 0 || f.Tag == nil {
			continue
		}

		tag := reflect.StructTag(f.Tag.Value)
		t := tag.Get("gpl")
		if t == "" {
			t = tag.Get("`gpl")
		}
		if t == "-" || t == "" {
			continue
		}

		options := parseFieldOptions(f.Tag.Value)
		colName := parseColumnName(f.Tag.Value)
		var isRelation bool
		var insertable, updateable, selectable bool
		if len(options) == 0 {
			insertable = true
			updateable = true
			selectable = true
		} else {
			for _, o := range options {
				if o == optionRelation {
					isRelation = true
				}

				if o == optionSelect {
					selectable = true
				}

				if o == optionInsert {
					insertable = true
				}

				if o == optionUpdate {
					updateable = true
				}
			}
		}

		structFields = append(structFields, field{
			Name:       f.Names[0].String(),
			Type:       g.getTokenBetween(source, f.Type.Pos(), f.Type.End()),
			ColumnName: colName,
			IsRelation: isRelation,
			Insert:     insertable,
			Update:     updateable,
			Select:     selectable,
			Tags:       parseNonGplTags(f.Tag.Value),
		})
	}

	if len(structFields) == 0 {
		return nil, nil
	}

	tableName := strcase.ToSnake(structName)
	for _, tn := range g.TableName {
		parts := strings.Split(tn, ":")
		if parts[0] == structName {
			tableName = parts[1]
		}
	}

	var insPls, insFlds, joinFlds, updSetts []string
	for idx, sf := range structFields {
		if !sf.IsRelation {
			if sf.Insert {
				insPls = append(insPls, g.placeholder(sf.ColumnName, idx))
				insFlds = append(insFlds, sf.ColumnName)
			}

			if sf.Update {
				updSetts = append(updSetts, sf.ColumnName+" = "+g.placeholder(sf.ColumnName, idx))
			}

			if sf.Select {
				joinFlds = append(joinFlds, fmt.Sprintf("%[1]s.%[2]s as \"%[1]s.%[2]s\"", tableName, sf.ColumnName))
			}
		}
	}

	params := structModel{
		TableName: tableName,
		Name:      structName,
		Insert: &insertParams{
			PlaceHolders: strings.Join(insPls, ", "),
			Fields:       strings.Join(insFlds, ", "),
		},
		Update: &updateParams{
			Setters: strings.Join(updSetts, ", "),
		},
		Join: &joinParams{
			Fields: strings.Join(joinFlds, ", "),
		},
		Package: g.Pkg,
		Fields:  structFields,
	}

	return &params, nil
}

const (
	optionRelation string = "relation"
	optionSelect   string = "select"
	optionInsert   string = "insert"
	optionUpdate   string = "update"
)

func parseColumnName(tag string) string {
	re := regexp.MustCompile(`gpl:"(.*?)"`)
	tts := re.FindStringSubmatch(tag)
	if len(tts) == 0 {
		return ""
	}
	tts = strings.Split(tts[1], ",")

	return tts[0]
}

func parseFieldOptions(tag string) []string {
	re := regexp.MustCompile(`gpl:"(.*?)"`)
	tts := re.FindStringSubmatch(tag)
	if len(tts) <= 1 {
		return nil
	}
	tts = strings.Split(tts[1], ",")

	return tts[1:]
}

func parseNonGplTags(tag string) *string {
	re := regexp.MustCompile("gpl:\".*?\"")
	tt := string(re.ReplaceAll([]byte(tag), []byte("")))
	tt = strings.Trim(tt, " ")
	if tt == "``" {
		return nil
	}

	return &tt
}

func (g *generator) generate(params *templateInput) error {
	ass, err := Asset("out.tmpl")
	if err != nil {
		return err
	}

	funcs := template.FuncMap{
		"Contains": strings.Contains,
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
