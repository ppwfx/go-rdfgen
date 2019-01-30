package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCarUsageType strings.Replacer
var strconvCarUsageType strconv.NumError

var basicCarUsageTypeTraitMapping = map[string]func(*schema.CarUsageTypeTrait, []string){}
var customCarUsageTypeTraitMapping = map[string]func(*schema.CarUsageTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CarUsageType", func(ctx jsonld.Context) (interface{}) {
		return NewCarUsageTypeFromContext(ctx)
	})

}

func NewCarUsageTypeFromContext(ctx jsonld.Context) (x schema.CarUsageType) {
	x.Type = "http://schema.org/CarUsageType"
	MapBasicToQualitativeValueTrait(ctx, &x.QualitativeValueTrait)
	MapCustomToQualitativeValueTrait(ctx, &x.QualitativeValueTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCarUsageTypeTrait(ctx, &x.CarUsageTypeTrait)
	MapCustomToCarUsageTypeTrait(ctx, &x.CarUsageTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCarUsageTypeTrait(ctx jsonld.Context, x *schema.CarUsageTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCarUsageTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCarUsageTypeTrait(ctx jsonld.Context, x *schema.CarUsageTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCarUsageTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}