package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDistillery strings.Replacer
var strconvDistillery strconv.NumError

var basicDistilleryTraitMapping = map[string]func(*schema.DistilleryTrait, []string){}
var customDistilleryTraitMapping = map[string]func(*schema.DistilleryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Distillery", func(ctx jsonld.Context) (interface{}) {
		return NewDistilleryFromContext(ctx)
	})

}

func NewDistilleryFromContext(ctx jsonld.Context) (x schema.Distillery) {
	x.Type = "http://schema.org/Distillery"
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


	MapBasicToDistilleryTrait(ctx, &x.DistilleryTrait)
	MapCustomToDistilleryTrait(ctx, &x.DistilleryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDistilleryTrait(ctx jsonld.Context, x *schema.DistilleryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDistilleryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDistilleryTrait(ctx jsonld.Context, x *schema.DistilleryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDistilleryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}