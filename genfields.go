//go:build ignore

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"strings"
)

type fieldType struct {
	Name   string
	Parent string
	Values []string
}

var (
	types []fieldType

	fieldTypeName = flag.String(
		"type",
		"Field",
		"The type to generate for",
	)
)

const fieldConvertionTemplate = `{{range $type := .}}
// String returns the english name for {{ $type.Name | short}}.
// String implements [fmt.Stringer] interface.
func ({{$type.Name | short}} {{$type.Name}}) String() string {
	return {{$type.Name | short}}.Localize(language.English)
}

// Localize returns the localized name for {{$type.Name | short}} in the language langTag.
func ({{$type.Name | short}} {{$type.Name}}) Localize(langTag language.Tag) string {
	{{- range $childType := $}}{{if eq $childType.Parent $type.Name}}
	if {{$childType.Name | short}}, ok := {{$type.Name |short}}.{{$childType.Name}}(); ok {
		return {{$childType.Name | short}}.Localize(langTag)
	}
{{end}}{{end}}
	switch {{$type.Name | short}} {
	{{- range $value := $type.Values }}
	case {{ $value }}:
		return lang.MustLocalize("monopoly.field.{{ $value | lower}}", langTag){{end}}
	default:
		return "UNKNOWN"
	}
}

// GoString implements [fmt.GoStringer] interface.
func ({{$type.Name | short}} {{$type.Name}}) GoString() string {
	{{- range $childType := $}}{{if eq $childType.Parent $type.Name}}
	if {{$childType.Name | short}}, ok := {{$type.Name |short}}.{{$childType.Name}}(); ok {
		return {{$childType.Name | short}}.GoString()
	}
{{end}}{{end}}
	switch {{$type.Name | short}} {
	{{- range $value := $type.Values }}
	case {{$value}}:
		return "{{$value}}"{{end}}
	default:
		return "UNKNOWN"
	}
}
{{range $childType := $}}{{if eq $childType.Parent $type.Name}}
// {{$childType.Name}} converts a [{{$type.Name}}] into a [{{$childType.Name}}] and reports weather {{$type.Name | short}} is a {{$childType.Name}}.
func ({{$type.Name | short}} {{$type.Name}}) {{$childType.Name}}() ({{$childType.Name}}, bool) {
	switch {{$childType.Name | short}} := {{$childType.Name}}({{$type.Name | short}}); {{$childType.Name | short}} {
	case {{range $index, $value := $childType.Values}}{{$value}}{{if ne $index (sub1 (len $childType.Values))}},
		{{end}}{{end}}{{ range $otherType := $}}{{if eq $otherType.Parent $childType.Name}}{{range $value := $otherType.Values}},
		{{$otherType.Parent}}({{$value}}){{end}}{{end}}{{end}}:
		return {{$childType.Name | short}}, true
	default:
		return -1, false
	}
}
{{end}}{{end}}{{end}}`

func init() {
	flag.Parse()
}

func main() {
	fieldConvTmpl, err := template.New("fields").
		Funcs(template.FuncMap{
			"lower": func(s string) string { return strings.ToLower(s) },
			"short": func(s string) string { return strings.ToLower(s[0:1]) },
			"sub1":  func(n int) int { return n - 1 },
		}).
		Parse(fieldConvertionTemplate)
	if err != nil {
		log.Fatalf("Failed to parse field conversion template: %v", err)
	}

	goFile, err := parser.ParseFile(token.NewFileSet(), "types.go", nil, 0)
	if err != nil {
		log.Fatalf("Failed to parse go file: %v", err)
	}
	parseFileDeclarations(goFile)

	outFile, err := os.Create("fields.go")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer outFile.Close()

	outFile.WriteString(fmt.Sprintf(`package %s

import (
	"github.com/%s/%s/lang"
	"golang.org/x/text/language"
)
`,
		goFile.Name, "Kesuaheli", goFile.Name))

	err = fieldConvTmpl.Execute(outFile, types)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("pkg %#v", goFile)
}

func parseFileDeclarations(goFile *ast.File) {
	for _, decl := range goFile.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			return
		}
		var valType ast.Expr
		for _, spec := range genDecl.Specs {
			switch s := spec.(type) {
			case *ast.TypeSpec:
				parseTypeSpecification(s)
			case *ast.ValueSpec:
				parseValueSpecification(s, valType)
			default:
				log.Printf("unknown spec: %#v", s)
				continue
			}
		}
	}
}

func parseTypeSpecification(s *ast.TypeSpec) {
	topLevelType := getTopLevelType(s)
	if topLevelType == nil {
		log.Printf("Failed to get top level type for type %s", s.Name)
		return
	}
	if topLevelType.Name.String() != *fieldTypeName {
		return
	}

	var parentName string
	if s != topLevelType {
		tType, ok := s.Type.(*ast.Ident)
		if !ok {
			return
		}
		parentName = tType.String()
	}
	addFieldType(s.Name.String(), parentName)
}

func parseValueSpecification(s *ast.ValueSpec, valType ast.Expr) {
	if s.Type != nil {
		valType = s.Type
	}
	topLevelType := getTopLevelType(valType)
	if topLevelType == nil {
		log.Printf("Failed to get top level type for value %v (%v) ", s, s.Type)
		return
	}
	if topLevelType.Name.String() != *fieldTypeName {
		return
	}

	vType, ok := valType.(*ast.Ident)
	if !ok {
		return
	}

	for _, name := range s.Names {
		addFieldValue(vType.String(), name.String())
	}
}

func getTopLevelType(t ast.Node) *ast.TypeSpec {
	var typeIdent *ast.Ident
	switch t := t.(type) {
	case *ast.TypeSpec:
		var ok bool
		typeIdent, ok = t.Type.(*ast.Ident)
		if !ok {
			return nil
		}
		if typeIdent.Obj == nil {
			return t
		}
	case *ast.Ident:
		if t.Obj == nil {
			return nil
		}
		typeIdent = t
	default:
		return nil
	}
	if typeIdent.Obj.Decl == nil {
		return nil
	}

	parentSpec, ok := typeIdent.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return nil
	}
	return getTopLevelType(parentSpec)
}

func addFieldType(typeName, parentName string) {
	for i, v := range types {
		if v.Name != typeName {
			continue
		}
		v.Parent = parentName
		types[i] = v
		return
	}

	types = append(types, fieldType{
		Name:   typeName,
		Parent: parentName,
	})
}

func addFieldValue(typeName, valueName string) {
	for i, v := range types {
		if v.Name != typeName {
			continue
		}
		v.Values = append(v.Values, valueName)
		types[i] = v
		return
	}

	types = append(types, fieldType{
		Name:   typeName,
		Values: []string{valueName},
	})
}
