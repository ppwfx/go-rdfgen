package main

import (
	"log"
	"github.com/knakk/rdf"
	"regexp"
	"strings"
	"github.com/davecgh/go-spew/spew"
	"github.com/21stio/go-rdfgen/pkg/types"
	"github.com/21stio/go-rdfgen/pkg/helper"
	"github.com/21stio/go-rdfgen/pkg/types/gen"
	"errors"
	"github.com/21stio/go-rdfgen/pkg/templates"
	"os"
	"github.com/21stio/go-rdfgen/pkg/resolve"
	"github.com/piprate/json-gold/ld"
)

// definedBy -> package
// rdfs:subClassOf -> inherits

//type Package struct {
//	structs
//}
//
//type Struct struct {
//	Subject string
//	SubClassOf
//}

var definedBy = map[string]bool{}
var subjects = map[rdf.Subject][]rdf.Triple{}

var prefixRegex = regexp.MustCompile(`@prefix [a-z]*: <([htps]*:\/\/[a-zA-Z0-9.\/-]*#?)> .`)

var baseTypes = map[string]bool{}

func init() {
	baseTypes["bool"] = true
	baseTypes["string"] = true
	baseTypes["int"] = true
	baseTypes["int8"] = true
	baseTypes["int16"] = true
	baseTypes["int32"] = true
	baseTypes["int64"] = true
	baseTypes["uint"] = true
	baseTypes["uint8"] = true
	baseTypes["uint16"] = true
	baseTypes["uint32"] = true
	baseTypes["uint64"] = true
	baseTypes["uintptr"] = true
	baseTypes["byte"] = true
	baseTypes["rune"] = true
	baseTypes["float32"] = true
	baseTypes["float64"] = true
	baseTypes["complex64"] = true
	baseTypes["complex128"] = true
	baseTypes["interface{}"] = true
}

func main() {
	err := func() (err error) {
		err = compact()
		if err != nil {
			return
		}

		return
	}()
	if err != nil {
		log.Fatal(err)
	}
}

func mapType(t string) (string) {
	typeMappers := []types.TypeMapper{helper.SchemaOrg{}}

	for _, tm := range typeMappers {
		if tm.CanMapType(t) {
			t = tm.MapType(t)
		}
	}

	return t
}

func IsBasicType(t string) (bool) {
	_, ok := baseTypes[t]

	return ok
}

func resolveSubClasses(subject string, m map[string]map[string][]string) (subclasses []string) {
	for _, v := range m[subject]["http://www.w3.org/2000/01/rdf-schema#subClassOf"] {
		if strings.Contains(v, "w3.org") {
			continue
		}

		subclasses = append(subclasses, v)

		subclasses = append(subclasses, resolveSubClasses(v, m)...)
	}

	return
}

func TrimPrefix(s []string, prefix string) ([]string) {
	for i, v := range s {
		s[i] = strings.TrimPrefix(v, prefix)
	}

	return s
}

func UniqueStrings(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func generateSchemaOrg() (err error) {
	t, err := resolve.ResolveTriples("http://schema.org/OpeningHoursSpecification.jsonld")
	if err != nil {
		return
	}

	t = t.RemoveDuplicates()

	classes := getRdfClasses(t)
	subClasses := getRdfSubClasses(t)

	includes := t.FilterByPredicate("http://schema.org/domainIncludes").GroupByObject()

	m := t.MapBySubject()
	subMm := subClasses.MapBySubject()

	structs := map[string]gen.Struct{}
	for _, c := range classes {
		s := gen.Struct{
			Package: "schema",
			Name:    strings.Replace(c.Subject, "http://schema.org/", "", -1),
			Type:    c.Subject,
		}

		extends := resolveSubClasses(c.Subject, subMm)
		extends = TrimPrefix(extends, "http://schema.org/")
		s.Extends = UniqueStrings(extends)

		fields := []gen.Field{}
		for _, i := range includes[c.Subject] {
			label := m[i.Subject]["http://www.w3.org/2000/01/rdf-schema#label"][0]

			comment := m[i.Subject]["http://www.w3.org/2000/01/rdf-schema#comment"][0]

			comment += "\n\t//\n\t// RangeIncludes:"
			for _, v := range m[i.Subject]["http://schema.org/rangeIncludes"] {
				comment += "\n\t// - " + v
			}
			comment += "\n\t//"

			rangeIncludes := m[i.Subject]["http://schema.org/rangeIncludes"]

			f := gen.Field{}
			f.Comment = comment
			f.Name = strings.Title(label)
			f.JsonTag = label
			f.JsonLdTag = i.Subject

			if len(rangeIncludes) == 0 {
				err = errors.New("len(rangeIncludes) == 0")

				return
			}

			if len(rangeIncludes) == 1 {
				f.Type = mapType(rangeIncludes[0])
			}

			if len(rangeIncludes) > 1 {
				f.Type = "interface{}"
			}

			f.IsBasicType = IsBasicType(f.Type)

			fields = append(fields, f)
		}
		s.Fields = fields

		structs[c.Subject] = s
	}

	for _, s := range structs {
		p := "./gen/schema/" + s.Name + ".go"
		err = os.Remove(p)
		if err != nil {
			if !os.IsNotExist(err) {
				return
			}
		}

		var f *os.File
		f, err = os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return
		}

		err = templates.Type.Execute(f, s)
		if err != nil {
			return
		}

		p = "./gen/schema/mapping/" + s.Name + ".go"
		err = os.Remove(p)
		if err != nil {
			if !os.IsNotExist(err) {
				return
			}
		}

		f, err = os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return
		}

		err = templates.Map.Execute(f, s)
		if err != nil {
			return
		}

		err = f.Close()
		if err != nil {
			return
		}
	}

	spew.Dump(structs)

	return
}

func getRdfClasses(triples types.Triples) ([]types.Triple) {
	return triples.
		FilterBySubjectContains("schema.org").
		FilterByPredicate("http://www.w3.org/1999/02/22-rdf-syntax-ns#type").
		FilterByObject("http://www.w3.org/2000/01/rdf-schema#Class")
}

func getRdfSubClasses(triples types.Triples) (types.Triples) {
	return triples.
		FilterBySubjectContains("schema.org").
		FilterByPredicate("http://www.w3.org/2000/01/rdf-schema#subClassOf")
}
