package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBedType strings.Replacer
var strconvBedType strconv.NumError

var basicBedTypeTraitMapping = map[string]func(*schema.BedTypeTrait, []string){}
var customBedTypeTraitMapping = map[string]func(*schema.BedTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BedType", func(ctx jsonld.Context) (interface{}) {
		return NewBedTypeFromContext(ctx)
	})

}

func NewBedTypeFromContext(ctx jsonld.Context) (x schema.BedType) {
	x.Type = "http://schema.org/BedType"
	MapBasicToQualitativeValueTrait(ctx, &x.QualitativeValueTrait)
	MapCustomToQualitativeValueTrait(ctx, &x.QualitativeValueTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBedTypeTrait(ctx, &x.BedTypeTrait)
	MapCustomToBedTypeTrait(ctx, &x.BedTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBedTypeTrait(ctx jsonld.Context, x *schema.BedTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBedTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBedTypeTrait(ctx jsonld.Context, x *schema.BedTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBedTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}