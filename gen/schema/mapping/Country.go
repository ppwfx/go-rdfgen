package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCountry strings.Replacer
var strconvCountry strconv.NumError

var basicCountryTraitMapping = map[string]func(*schema.CountryTrait, []string){}
var customCountryTraitMapping = map[string]func(*schema.CountryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Country", func(ctx jsonld.Context) (interface{}) {
		return NewCountryFromContext(ctx)
	})

}

func NewCountryFromContext(ctx jsonld.Context) (x schema.Country) {
	x.Type = "http://schema.org/Country"
	MapBasicToAdministrativeAreaTrait(ctx, &x.AdministrativeAreaTrait)
	MapCustomToAdministrativeAreaTrait(ctx, &x.AdministrativeAreaTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCountryTrait(ctx, &x.CountryTrait)
	MapCustomToCountryTrait(ctx, &x.CountryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCountryTrait(ctx jsonld.Context, x *schema.CountryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCountryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCountryTrait(ctx jsonld.Context, x *schema.CountryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCountryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}