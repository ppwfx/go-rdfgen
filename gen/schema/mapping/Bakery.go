package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBakery strings.Replacer
var strconvBakery strconv.NumError

var basicBakeryTraitMapping = map[string]func(*schema.BakeryTrait, []string){}
var customBakeryTraitMapping = map[string]func(*schema.BakeryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Bakery", func(ctx jsonld.Context) (interface{}) {
		return NewBakeryFromContext(ctx)
	})

}

func NewBakeryFromContext(ctx jsonld.Context) (x schema.Bakery) {
	x.Type = "http://schema.org/Bakery"
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


	MapBasicToBakeryTrait(ctx, &x.BakeryTrait)
	MapCustomToBakeryTrait(ctx, &x.BakeryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBakeryTrait(ctx jsonld.Context, x *schema.BakeryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBakeryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBakeryTrait(ctx jsonld.Context, x *schema.BakeryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBakeryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}