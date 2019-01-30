package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWinery strings.Replacer
var strconvWinery strconv.NumError

var basicWineryTraitMapping = map[string]func(*schema.WineryTrait, []string){}
var customWineryTraitMapping = map[string]func(*schema.WineryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Winery", func(ctx jsonld.Context) (interface{}) {
		return NewWineryFromContext(ctx)
	})

}

func NewWineryFromContext(ctx jsonld.Context) (x schema.Winery) {
	x.Type = "http://schema.org/Winery"
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


	MapBasicToWineryTrait(ctx, &x.WineryTrait)
	MapCustomToWineryTrait(ctx, &x.WineryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWineryTrait(ctx jsonld.Context, x *schema.WineryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWineryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWineryTrait(ctx jsonld.Context, x *schema.WineryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWineryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}