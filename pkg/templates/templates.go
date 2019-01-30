package templates

import "text/template"

var Type = template.Must(template.New("").Parse(
`package {{ .Package }}

type {{ .Name }}Trait struct {
{{ range $i, $f := .Fields }}
	{{ if ne $f.Comment "" }}// {{ $f.Comment }}{{ end }}
	{{ $f.Name }} {{ if not $f.IsBasicType }}*{{ end }}{{ $f.Type }} ` + "`" + `json:"{{ $f.JsonTag }},omitempty" jsonld:"{{ $f.JsonLdTag }}"` + "`" + `
{{ end }}
}

type {{ .Name }} struct {
	MetaTrait
	{{ .Name }}Trait
{{- range $i, $e := .Extends }}
	{{ $e }}Trait
{{- end }}
	AdditionalTrait
}

func New{{ .Name }}() (x {{ .Name }}) {
	x.Type = "{{ .Type }}"

	return
}
`))

var Map = template.Must(template.New("").Parse(
	`package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/{{ .Package }}"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var strings{{ .Name }} strings.Replacer
var strconv{{ .Name }} strconv.NumError

var basic{{ .Name }}TraitMapping = map[string]func(*{{ .Package }}.{{ .Name }}Trait, []string){}
var custom{{ .Name }}TraitMapping = map[string]func(*{{ .Package }}.{{ .Name }}Trait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("{{ .Type }}", func(ctx jsonld.Context) (interface{}) {
		return New{{$.Name}}FromContext(ctx)
	})

{{ range $i, $f := .Fields }}
{{- if eq $f.Type "string" }}
	basic{{ $.Name }}TraitMapping["{{ $f.JsonLdTag }}"] = func(x *{{ $.Package }}.{{ $.Name }}Trait, s []string) {
		x.{{ $f.Name }} = s[0]
	}
{{ end -}}

{{- if eq $f.Type "float64" }}
	basic{{ $.Name }}TraitMapping["{{ $f.JsonLdTag }}"] = func(x *{{ $.Package }}.{{ $.Name }}Trait, s []string) {
		var err error
		x.{{ $f.Name }}, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}
{{ end -}}

{{- if eq $f.Type "bool" }}
	basic{{ $.Name }}TraitMapping["{{ $f.JsonLdTag }}"] = func(x *{{ $.Package }}.{{ $.Name }}Trait, s []string) {
		var err error
		x.{{ $f.Name }}, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}
{{ end }}
{{ end -}}

{{ range $i, $f := .Fields }}
{{- if eq $f.Type "interface{}" }}
	custom{{ $.Name }}TraitMapping["{{ $f.JsonLdTag }}"] = func(x *{{ $.Package }}.{{ $.Name }}Trait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.{{ $f.Name }}, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.{{ $f.Name }} = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.{{ $f.Name }} = s
		}
	}
{{ end -}}
{{- if not $f.IsBasicType }}
	custom{{ $.Name }}TraitMapping["{{ $f.JsonLdTag }}"] = func(x *{{ $.Package }}.{{ $.Name }}Trait, ctx jsonld.Context, s string){
		var y schema.{{ $f.Type }}
		if strings.HasPrefix(s, "_:") {
			y = New{{ $f.Type }}FromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.New{{ $f.Type }}()
			y.Id = s
		}

		x.{{ $f.Name }} = &y

		return
	}
{{ end -}}
{{ end -}}
}

func New{{ .Name }}FromContext(ctx jsonld.Context) (x {{ .Package }}.{{ .Name }}) {
	x.Type = "{{ .Type }}"

{{- range $i, $e := .Extends }}
	MapBasicTo{{ $e }}Trait(ctx, &x.{{ $e }}Trait)
	MapCustomTo{{ $e }}Trait(ctx, &x.{{ $e }}Trait)
{{ end }}

	MapBasicTo{{ $.Name }}Trait(ctx, &x.{{ $.Name }}Trait)
	MapCustomTo{{ $.Name }}Trait(ctx, &x.{{ $.Name }}Trait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicTo{{ .Name }}Trait(ctx jsonld.Context, x *{{ .Package }}.{{ .Name }}Trait) () {
	for k, v := range ctx.Current {
		f, ok := basic{{ .Name }}TraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomTo{{ .Name }}Trait(ctx jsonld.Context, x *{{ .Package }}.{{ .Name }}Trait) () {
	for k, v := range ctx.Current {
		f, ok := custom{{ .Name }}TraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}`))
