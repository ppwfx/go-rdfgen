package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsQuantity strings.Replacer
var strconvQuantity strconv.NumError

var basicQuantityTraitMapping = map[string]func(*schema.QuantityTrait, []string){}
var customQuantityTraitMapping = map[string]func(*schema.QuantityTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Quantity", func(ctx jsonld.Context) (interface{}) {
		return NewQuantityFromContext(ctx)
	})

}

func NewQuantityFromContext(ctx jsonld.Context) (x schema.Quantity) {
	x.Type = "http://schema.org/Quantity"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToQuantityTrait(ctx, &x.QuantityTrait)
	MapCustomToQuantityTrait(ctx, &x.QuantityTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToQuantityTrait(ctx jsonld.Context, x *schema.QuantityTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicQuantityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToQuantityTrait(ctx jsonld.Context, x *schema.QuantityTrait) () {
	for k, v := range ctx.Current {
		f, ok := customQuantityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}