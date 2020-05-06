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

	"github.com/gochore/uniq"
	"golang.org/x/tools/go/packages"
)

type GenerateResult struct {
	Content []byte
	Output  string
}

func (r *GenerateResult) String() string {
	return fmt.Sprintf("%s\n%s", r.Output, r.Content)
}

func Generate(ctx context.Context, dir string, types []string) ([]*GenerateResult, error) {
	cmd := fmt.Sprintf("gtag -types %s .", strings.Join(types, ","))

	types = types[:uniq.Strings(types)]

	pkgs, err := packages.Load(&packages.Config{
		Mode:    packages.NeedFiles,
		Context: ctx,
		Dir:     dir,
		Env:     os.Environ(),
		Fset:    token.NewFileSet(),
	})
	if err != nil {
		return nil, err
	}

	var files []string
	for _, pkg := range pkgs {
		files = append(files, pkg.GoFiles...)
	}

	var ret []*GenerateResult
	for _, file := range files {
		result, err := generateFile(ctx, cmd, file, types)
		if err != nil {
			return nil, err
		}
		if result != nil {
			ret = append(ret, result)
		}
	}
	for _, v := range ret {
		if err := v.Commit(); err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func generateFile(ctx context.Context, cmd, file string, types []string) (*GenerateResult, error) {
	f, err := loadFile(file)
	if err != nil {
		return nil, err
	}

	if f.Name == nil {
		return nil, fmt.Errorf("can not find package name")
	}
	pkg := f.Name.Name

	fieldM := map[string][]string{}
	for _, typ := range types {
		fields, ok := parseStructField(f, typ)
		if ok {
			fieldM[typ] = fields
		}
	}

	if len(fieldM) == 0 {
		return nil, nil
	}

	data := templateData{
		Package: pkg,
		Command: cmd,
	}

	for _, typ := range types {
		if fields, ok := fieldM[typ]; ok {
			data.Types = append(data.Types, templateDataType{
				Name:   typ,
				Fields: fields,
			})
		}
	}

	src, err := format.Source(execute(data))
	if err != nil {
		return nil, err
	}

	ret := &GenerateResult{
		Content: src,
		Output:  strings.TrimSuffix(file, ".go") + "_tag.go",
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

func parseStructField(f *ast.File, name string) ([]string, bool) {
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
		return nil, found
	}

	var ret []string
	for _, field := range fields {
		for _, v := range field.Names {
			ret = append(ret, v.Name)
		}
	}

	return ret, found
}
