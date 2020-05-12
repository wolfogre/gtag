package gtag

import (
	"bytes"
	"text/template"
)

type templateData struct {
	Package string
	Command string
	Types   []templateDataType
	Tags    []templateDataTag
}

type templateDataType struct {
	Name   string
	Fields []string
}

type templateDataTag struct {
	Name  string
	Value string
}

const templateLayout = `
// Code generated by gtag. DO NOT EDIT.
// See: https://github.com/wolfogre/gtag

//go:generate {{.Command}}
package {{.Package}}

import (
	"reflect"
	"strings"
)

{{$tags := .Tags}}
{{- range .Types}}

var (
	valueOf{{.Name}} = {{.Name}}{}
	typeOf{{.Name}}  = reflect.TypeOf(valueOf{{.Name}})

{{$type := .Name}}
{{- range .Fields}}
	_ = valueOf{{$type}}.{{.}}
	fieldOf{{$type}}{{.}}, _ = typeOf{{$type}}.FieldByName("{{.}}")
	tagOf{{$type}}{{.}} = fieldOf{{$type}}{{.}}.Tag
{{end}}
)

// {{$type}}Tags indicate tags of type {{$type}}
type {{$type}}Tags struct {
{{- range .Fields}}
	{{.}} string
{{- end}}
}

// Tags return specified tags of {{$type}}
func ({{$type}}) Tags(tag string, convert ...func(string) string) {{$type}}Tags {
	conv := func(in string) string { return strings.TrimSpace(strings.Split(in, ",")[0]) }
	if len(convert) > 0 && convert[0] != nil {
		conv = convert[0]
	}
	_ = conv
	return {{$type}}Tags{
{{- range .Fields}}
		{{.}}: conv(tagOf{{$type}}{{.}}.Get(tag)),
{{- end}}
	}
}

{{range $tags}}
// Tags{{.Name}} is alias of Tags("{{.Value}}")
func (v {{$type}}) Tags{{.Name}}() {{$type}}Tags {
	return v.Tags("{{.Value}}")
}

{{end}}

{{- end}}

`

func execute(data templateData) []byte {
	tp := template.Must(template.New("").Parse(templateLayout))

	out := &bytes.Buffer{}
	if err := tp.Execute(out, data); err != nil {
		panic(err)
	}

	return out.Bytes()
}
