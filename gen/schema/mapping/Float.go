package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFloat strings.Replacer
var strconvFloat strconv.NumError

var basicFloatTraitMapping = map[string]func(*schema.FloatTrait, []string){}
var customFloatTraitMapping = map[string]func(*schema.FloatTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Float", func(ctx jsonld.Context) (interface{}) {
		return NewFloatFromContext(ctx)
	})

}

func NewFloatFromContext(ctx jsonld.Context) (x schema.Float) {
	x.Type = "http://schema.org/Float"
	MapBasicToNumberTrait(ctx, &x.NumberTrait)
	MapCustomToNumberTrait(ctx, &x.NumberTrait)


	MapBasicToFloatTrait(ctx, &x.FloatTrait)
	MapCustomToFloatTrait(ctx, &x.FloatTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFloatTrait(ctx jsonld.Context, x *schema.FloatTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFloatTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFloatTrait(ctx jsonld.Context, x *schema.FloatTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFloatTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}