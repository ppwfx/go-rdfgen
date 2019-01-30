package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAdultEntertainment strings.Replacer
var strconvAdultEntertainment strconv.NumError

var basicAdultEntertainmentTraitMapping = map[string]func(*schema.AdultEntertainmentTrait, []string){}
var customAdultEntertainmentTraitMapping = map[string]func(*schema.AdultEntertainmentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AdultEntertainment", func(ctx jsonld.Context) (interface{}) {
		return NewAdultEntertainmentFromContext(ctx)
	})

}

func NewAdultEntertainmentFromContext(ctx jsonld.Context) (x schema.AdultEntertainment) {
	x.Type = "http://schema.org/AdultEntertainment"
	MapBasicToEntertainmentBusinessTrait(ctx, &x.EntertainmentBusinessTrait)
	MapCustomToEntertainmentBusinessTrait(ctx, &x.EntertainmentBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToAdultEntertainmentTrait(ctx, &x.AdultEntertainmentTrait)
	MapCustomToAdultEntertainmentTrait(ctx, &x.AdultEntertainmentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAdultEntertainmentTrait(ctx jsonld.Context, x *schema.AdultEntertainmentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAdultEntertainmentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAdultEntertainmentTrait(ctx jsonld.Context, x *schema.AdultEntertainmentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAdultEntertainmentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}