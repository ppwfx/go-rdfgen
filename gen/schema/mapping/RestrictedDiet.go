package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRestrictedDiet strings.Replacer
var strconvRestrictedDiet strconv.NumError

var basicRestrictedDietTraitMapping = map[string]func(*schema.RestrictedDietTrait, []string){}
var customRestrictedDietTraitMapping = map[string]func(*schema.RestrictedDietTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RestrictedDiet", func(ctx jsonld.Context) (interface{}) {
		return NewRestrictedDietFromContext(ctx)
	})

}

func NewRestrictedDietFromContext(ctx jsonld.Context) (x schema.RestrictedDiet) {
	x.Type = "http://schema.org/RestrictedDiet"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRestrictedDietTrait(ctx, &x.RestrictedDietTrait)
	MapCustomToRestrictedDietTrait(ctx, &x.RestrictedDietTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRestrictedDietTrait(ctx jsonld.Context, x *schema.RestrictedDietTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRestrictedDietTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRestrictedDietTrait(ctx jsonld.Context, x *schema.RestrictedDietTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRestrictedDietTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}