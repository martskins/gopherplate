package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

const (
	optionRelation string = "relation"
	optionSelect   string = "select"
	optionInsert   string = "insert"
	optionUpdate   string = "update"
)

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
	Structs      []*structModel
	Insert       bool
	Update       bool
	Join         bool
	PackageName  string
	ShouldExport bool
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

func (g *generator) structsInPath(path string) ([]*structModel, error) {
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

		sm, err := g.structModelFromFields(y.Fields.List, string(bts), name)
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

func getTokenBetween(source string, start, end token.Pos) string {
	typeInSource := source[start-1 : end-1]
	return typeInSource
}

func (g *generator) structModelFromFields(fields []*ast.Field, source string, structName string) (*structModel, error) {
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
			Type:       getTokenBetween(source, f.Type.Pos(), f.Type.End()),
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
		Package: g.Package,
		Fields:  structFields,
	}

	return &params, nil
}
