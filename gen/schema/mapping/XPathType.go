package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsXPathType strings.Replacer
var strconvXPathType strconv.NumError

var basicXPathTypeTraitMapping = map[string]func(*schema.XPathTypeTrait, []string){}
var customXPathTypeTraitMapping = map[string]func(*schema.XPathTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/XPathType", func(ctx jsonld.Context) (interface{}) {
		return NewXPathTypeFromContext(ctx)
	})

}

func NewXPathTypeFromContext(ctx jsonld.Context) (x schema.XPathType) {
	x.Type = "http://schema.org/XPathType"
	MapBasicToTextTrait(ctx, &x.TextTrait)
	MapCustomToTextTrait(ctx, &x.TextTrait)


	MapBasicToXPathTypeTrait(ctx, &x.XPathTypeTrait)
	MapCustomToXPathTypeTrait(ctx, &x.XPathTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToXPathTypeTrait(ctx jsonld.Context, x *schema.XPathTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicXPathTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToXPathTypeTrait(ctx jsonld.Context, x *schema.XPathTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customXPathTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}