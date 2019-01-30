package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBusinessFunction strings.Replacer
var strconvBusinessFunction strconv.NumError

var basicBusinessFunctionTraitMapping = map[string]func(*schema.BusinessFunctionTrait, []string){}
var customBusinessFunctionTraitMapping = map[string]func(*schema.BusinessFunctionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BusinessFunction", func(ctx jsonld.Context) (interface{}) {
		return NewBusinessFunctionFromContext(ctx)
	})

}

func NewBusinessFunctionFromContext(ctx jsonld.Context) (x schema.BusinessFunction) {
	x.Type = "http://schema.org/BusinessFunction"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBusinessFunctionTrait(ctx, &x.BusinessFunctionTrait)
	MapCustomToBusinessFunctionTrait(ctx, &x.BusinessFunctionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBusinessFunctionTrait(ctx jsonld.Context, x *schema.BusinessFunctionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBusinessFunctionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBusinessFunctionTrait(ctx jsonld.Context, x *schema.BusinessFunctionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBusinessFunctionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}