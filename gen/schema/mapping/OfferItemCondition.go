package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOfferItemCondition strings.Replacer
var strconvOfferItemCondition strconv.NumError

var basicOfferItemConditionTraitMapping = map[string]func(*schema.OfferItemConditionTrait, []string){}
var customOfferItemConditionTraitMapping = map[string]func(*schema.OfferItemConditionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OfferItemCondition", func(ctx jsonld.Context) (interface{}) {
		return NewOfferItemConditionFromContext(ctx)
	})

}

func NewOfferItemConditionFromContext(ctx jsonld.Context) (x schema.OfferItemCondition) {
	x.Type = "http://schema.org/OfferItemCondition"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOfferItemConditionTrait(ctx, &x.OfferItemConditionTrait)
	MapCustomToOfferItemConditionTrait(ctx, &x.OfferItemConditionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOfferItemConditionTrait(ctx jsonld.Context, x *schema.OfferItemConditionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOfferItemConditionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOfferItemConditionTrait(ctx jsonld.Context, x *schema.OfferItemConditionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOfferItemConditionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}