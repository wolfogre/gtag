package gtag

import (
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

type GenerateResult struct {
	Content []byte
	Output  string
}

func Generate(ctx context.Context, file, name string) (*GenerateResult, error) {
	f, err := loadFile(file)
	if err != nil {
		return nil, err
	}

	fields, err := parseStructField(f, name)
	if err != nil {
		return nil, err
	}

	if f.Name == nil {
		return nil, fmt.Errorf("can not find package name")
	}
	pkg := f.Name.Name

	data := templateData{
		Package: pkg,
		Type:    name,
		Fields:  fields,
	}

	src, err := format.Source(execute(data))
	if err != nil {
		return nil, err
	}

	ret := &GenerateResult{
		Content: src,
		Output:  strings.TrimSuffix(file, ".go") + "_tag.go",
	}

	if err := ret.Commit(); err != nil {
		return nil, err
	}

	return ret, nil
}

func (r *GenerateResult) Commit() error {
	if len(r.Content) == 0 {
		return nil
	}
	return ioutil.WriteFile(r.Output, r.Content, 0666)
}

func loadFile(name string) (*ast.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parser.ParseFile(token.NewFileSet(), name, f, 0)
}

func parseStructField(f *ast.File, name string) ([]string, error) {
	var fields []*ast.Field
	found := false

	ast.Inspect(f, func(n ast.Node) bool {
		if found {
			return false
		}
		switch t := n.(type) {
		case *ast.TypeSpec:
			if t.Name != nil && t.Name.Name == name {
				return true
			}
			return false
		case *ast.StructType:
			fields = t.Fields.List
			found = true
			return false
		}
		return true
	})

	if !found {
		return nil, fmt.Errorf("can not find struct %q", name)
	}

	var ret []string
	for _, field := range fields {
		for _, v := range field.Names {
			ret = append(ret, v.Name)
		}
	}

	return ret, nil
}
