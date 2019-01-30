package jsonld

import (
	"github.com/21stio/go-rdfgen/pkg/parser"
	"errors"
)

var rdfTypeHandlers = map[string]func(Context) (interface{}){}

func RegisterRdfType(t string, f func(Context) (i interface{})) {
	rdfTypeHandlers[t] = f
}

type Context struct {
	Current  map[string][]string
	Subjects map[string]map[string][]string
	Decoded  map[string]interface{}
}

func NewContext() (ctx Context) {
	ctx.Decoded = map[string]interface{}{}

	return
}

func (ctx Context) SetCurrent(s string) (Context) {
	ctx.Current = ctx.Subjects[s]

	return ctx
}

func Unmarshal(data []byte) (v interface{}, err error) {
	p := parser.JsonLd{}

	triples, err := p.Parse(data)
	if err != nil {
		return
	}

	ctx := NewContext()
	ctx.Subjects = triples.MapBySubject()

	for k, _ := range ctx.Subjects {
		_, ok := ctx.Decoded[k]
		if ok {
			continue
		}

		ctx.Decoded[k], err = FillByRdfType(ctx.SetCurrent(k))
		if err != nil {
			return
		}

		delete(ctx.Subjects, k)
	}

	//s := spew.Sdump(ctx)
	//re := regexp.MustCompile("(?m)^.*" + `""` + ".*$[\r\n]+")
	//s = re.ReplaceAllString(s, "")
	//re = regexp.MustCompile("(?m)^.*" + `<nil>` + ".*$[\r\n]+")
	//s = re.ReplaceAllString(s, "")
	//
	//println(s)

	if len(ctx.Decoded) != 1 {
		err = errors.New("something went wrong")

		return
	}

	var k string
	for k = range ctx.Decoded {
	}

	v = ctx.Decoded[k]

	return
}

func FillByRdfType(ctx Context) (interface{}, error) {
	rdfType, ok := ctx.Current["http://www.w3.org/1999/02/22-rdf-syntax-ns#type"]
	if !ok {
		return nil, errors.New("no rdf type defined")
	}

	h, ok := rdfTypeHandlers[rdfType[0]]
	if !ok {
		return nil, errors.New("no rdf type handler defined")
	}

	x := h(ctx)

	return &x, nil
}
