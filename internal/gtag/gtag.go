package gtag

import (
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"

	"github.com/gochore/uniq"
	"golang.org/x/tools/go/packages"
)

type GenerateResult struct {
	Content []byte
	Source  string
	Output  string
}

func (r *GenerateResult) String() string {
	return fmt.Sprintf("%s\n%s", r.Output, r.Content)
}

func Generate(ctx context.Context, dir string, types []string, tags []string) ([]*GenerateResult, error) {
	types = types[:uniq.Strings(types)]
	tags = tags[:uniq.Strings(tags)]

	cmd := fmt.Sprintf("gtag -types %s -tags %s .", strings.Join(types, ","), strings.Join(tags, ","))

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
		result, err := generateFile(cmd, file, types, tags)
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

func generateFile(cmd, file string, types []string, tags []string) (*GenerateResult, error) {
	f, err := loadFile(file)
	if err != nil {
		return nil, err
	}

	if f.Name == nil {
		return nil, fmt.Errorf("can not find package name")
	}
	pkg := f.Name.Name

	fieldM := map[string][]*ast.Field{}
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
			var tmpFields []templateDataTypeField
			for _, field := range fields {
				tag := ""
				if field.Tag != nil {
					tag = field.Tag.Value
				}
				for _, name := range field.Names {
					tmpFields = append(tmpFields, templateDataTypeField{
						Name: name.Name,
						Tag:  tag,
					})
				}
			}

			data.Types = append(data.Types, templateDataType{
				Name:   typ,
				Fields: tmpFields,
			})
		}
	}

	for _, tag := range tags {
		data.Tags = append(data.Tags, templateDataTag{
			Name:  strings.ReplaceAll(cases.Title(language.English).String(strings.ReplaceAll(tag, "_", " ")), " ", ""),
			Value: tag,
		})
	}

	src, err := format.Source(execute(data))
	if err != nil {
		return nil, err
	}

	ret := &GenerateResult{
		Content: src,
		Source:  file,
		Output:  strings.TrimSuffix(file, ".go") + "_tag.go",
	}

	return ret, nil
}

func (r *GenerateResult) Commit() error {
	if len(r.Content) == 0 {
		return nil
	}
	return os.WriteFile(r.Output, r.Content, 0o644)
}

func loadFile(name string) (*ast.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parser.ParseFile(token.NewFileSet(), name, f, 0)
}

func parseStructField(f *ast.File, name string) ([]*ast.Field, bool) {
	var fields []*ast.Field
	found := false

	ast.Inspect(f, func(node ast.Node) bool {
		if found {
			return false
		}
		if t, ok := node.(*ast.TypeSpec); ok {
			if t.Name != nil && t.Name.Name == name {
				ast.Inspect(node, func(structNode ast.Node) bool {
					if found {
						return false
					}
					if t, ok := structNode.(*ast.StructType); ok {
						found = true
						fields = t.Fields.List
					}
					return !found
				})
			}
		}
		return !found
	})

	if !found {
		return nil, found
	}

	return fields, found
}
