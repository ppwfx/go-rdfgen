package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsApartmentComplex strings.Replacer
var strconvApartmentComplex strconv.NumError

var basicApartmentComplexTraitMapping = map[string]func(*schema.ApartmentComplexTrait, []string){}
var customApartmentComplexTraitMapping = map[string]func(*schema.ApartmentComplexTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ApartmentComplex", func(ctx jsonld.Context) (interface{}) {
		return NewApartmentComplexFromContext(ctx)
	})

}

func NewApartmentComplexFromContext(ctx jsonld.Context) (x schema.ApartmentComplex) {
	x.Type = "http://schema.org/ApartmentComplex"
	MapBasicToResidenceTrait(ctx, &x.ResidenceTrait)
	MapCustomToResidenceTrait(ctx, &x.ResidenceTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToApartmentComplexTrait(ctx, &x.ApartmentComplexTrait)
	MapCustomToApartmentComplexTrait(ctx, &x.ApartmentComplexTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToApartmentComplexTrait(ctx jsonld.Context, x *schema.ApartmentComplexTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicApartmentComplexTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToApartmentComplexTrait(ctx jsonld.Context, x *schema.ApartmentComplexTrait) () {
	for k, v := range ctx.Current {
		f, ok := customApartmentComplexTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}