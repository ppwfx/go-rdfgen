package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsStructuredValue strings.Replacer
var strconvStructuredValue strconv.NumError

var basicStructuredValueTraitMapping = map[string]func(*schema.StructuredValueTrait, []string){}
var customStructuredValueTraitMapping = map[string]func(*schema.StructuredValueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/StructuredValue", func(ctx jsonld.Context) (interface{}) {
		return NewStructuredValueFromContext(ctx)
	})

}

func NewStructuredValueFromContext(ctx jsonld.Context) (x schema.StructuredValue) {
	x.Type = "http://schema.org/StructuredValue"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToStructuredValueTrait(ctx jsonld.Context, x *schema.StructuredValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicStructuredValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToStructuredValueTrait(ctx jsonld.Context, x *schema.StructuredValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customStructuredValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}