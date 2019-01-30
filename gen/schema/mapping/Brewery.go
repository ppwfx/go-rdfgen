package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBrewery strings.Replacer
var strconvBrewery strconv.NumError

var basicBreweryTraitMapping = map[string]func(*schema.BreweryTrait, []string){}
var customBreweryTraitMapping = map[string]func(*schema.BreweryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Brewery", func(ctx jsonld.Context) (interface{}) {
		return NewBreweryFromContext(ctx)
	})

}

func NewBreweryFromContext(ctx jsonld.Context) (x schema.Brewery) {
	x.Type = "http://schema.org/Brewery"
	MapBasicToFoodEstablishmentTrait(ctx, &x.FoodEstablishmentTrait)
	MapCustomToFoodEstablishmentTrait(ctx, &x.FoodEstablishmentTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToBreweryTrait(ctx, &x.BreweryTrait)
	MapCustomToBreweryTrait(ctx, &x.BreweryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBreweryTrait(ctx jsonld.Context, x *schema.BreweryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBreweryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBreweryTrait(ctx jsonld.Context, x *schema.BreweryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBreweryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}