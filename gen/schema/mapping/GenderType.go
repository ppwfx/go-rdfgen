package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGenderType strings.Replacer
var strconvGenderType strconv.NumError

var basicGenderTypeTraitMapping = map[string]func(*schema.GenderTypeTrait, []string){}
var customGenderTypeTraitMapping = map[string]func(*schema.GenderTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GenderType", func(ctx jsonld.Context) (interface{}) {
		return NewGenderTypeFromContext(ctx)
	})

}

func NewGenderTypeFromContext(ctx jsonld.Context) (x schema.GenderType) {
	x.Type = "http://schema.org/GenderType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGenderTypeTrait(ctx, &x.GenderTypeTrait)
	MapCustomToGenderTypeTrait(ctx, &x.GenderTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGenderTypeTrait(ctx jsonld.Context, x *schema.GenderTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGenderTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGenderTypeTrait(ctx jsonld.Context, x *schema.GenderTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGenderTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}