package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBoolean strings.Replacer
var strconvBoolean strconv.NumError

var basicBooleanTraitMapping = map[string]func(*schema.BooleanTrait, []string){}
var customBooleanTraitMapping = map[string]func(*schema.BooleanTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Boolean", func(ctx jsonld.Context) (interface{}) {
		return NewBooleanFromContext(ctx)
	})

}

func NewBooleanFromContext(ctx jsonld.Context) (x schema.Boolean) {
	x.Type = "http://schema.org/Boolean"

	MapBasicToBooleanTrait(ctx, &x.BooleanTrait)
	MapCustomToBooleanTrait(ctx, &x.BooleanTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBooleanTrait(ctx jsonld.Context, x *schema.BooleanTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBooleanTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBooleanTrait(ctx jsonld.Context, x *schema.BooleanTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBooleanTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}