package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNumber strings.Replacer
var strconvNumber strconv.NumError

var basicNumberTraitMapping = map[string]func(*schema.NumberTrait, []string){}
var customNumberTraitMapping = map[string]func(*schema.NumberTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Number", func(ctx jsonld.Context) (interface{}) {
		return NewNumberFromContext(ctx)
	})

}

func NewNumberFromContext(ctx jsonld.Context) (x schema.Number) {
	x.Type = "http://schema.org/Number"

	MapBasicToNumberTrait(ctx, &x.NumberTrait)
	MapCustomToNumberTrait(ctx, &x.NumberTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNumberTrait(ctx jsonld.Context, x *schema.NumberTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNumberTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNumberTrait(ctx jsonld.Context, x *schema.NumberTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNumberTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}