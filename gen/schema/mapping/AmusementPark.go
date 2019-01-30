package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAmusementPark strings.Replacer
var strconvAmusementPark strconv.NumError

var basicAmusementParkTraitMapping = map[string]func(*schema.AmusementParkTrait, []string){}
var customAmusementParkTraitMapping = map[string]func(*schema.AmusementParkTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AmusementPark", func(ctx jsonld.Context) (interface{}) {
		return NewAmusementParkFromContext(ctx)
	})

}

func NewAmusementParkFromContext(ctx jsonld.Context) (x schema.AmusementPark) {
	x.Type = "http://schema.org/AmusementPark"
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


	MapBasicToAmusementParkTrait(ctx, &x.AmusementParkTrait)
	MapCustomToAmusementParkTrait(ctx, &x.AmusementParkTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAmusementParkTrait(ctx jsonld.Context, x *schema.AmusementParkTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAmusementParkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAmusementParkTrait(ctx jsonld.Context, x *schema.AmusementParkTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAmusementParkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}