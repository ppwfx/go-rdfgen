package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInteger strings.Replacer
var strconvInteger strconv.NumError

var basicIntegerTraitMapping = map[string]func(*schema.IntegerTrait, []string){}
var customIntegerTraitMapping = map[string]func(*schema.IntegerTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Integer", func(ctx jsonld.Context) (interface{}) {
		return NewIntegerFromContext(ctx)
	})

}

func NewIntegerFromContext(ctx jsonld.Context) (x schema.Integer) {
	x.Type = "http://schema.org/Integer"
	MapBasicToNumberTrait(ctx, &x.NumberTrait)
	MapCustomToNumberTrait(ctx, &x.NumberTrait)


	MapBasicToIntegerTrait(ctx, &x.IntegerTrait)
	MapCustomToIntegerTrait(ctx, &x.IntegerTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToIntegerTrait(ctx jsonld.Context, x *schema.IntegerTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicIntegerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToIntegerTrait(ctx jsonld.Context, x *schema.IntegerTrait) () {
	for k, v := range ctx.Current {
		f, ok := customIntegerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}