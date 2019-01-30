package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPublicSwimmingPool strings.Replacer
var strconvPublicSwimmingPool strconv.NumError

var basicPublicSwimmingPoolTraitMapping = map[string]func(*schema.PublicSwimmingPoolTrait, []string){}
var customPublicSwimmingPoolTraitMapping = map[string]func(*schema.PublicSwimmingPoolTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PublicSwimmingPool", func(ctx jsonld.Context) (interface{}) {
		return NewPublicSwimmingPoolFromContext(ctx)
	})

}

func NewPublicSwimmingPoolFromContext(ctx jsonld.Context) (x schema.PublicSwimmingPool) {
	x.Type = "http://schema.org/PublicSwimmingPool"
	MapBasicToSportsActivityLocationTrait(ctx, &x.SportsActivityLocationTrait)
	MapCustomToSportsActivityLocationTrait(ctx, &x.SportsActivityLocationTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToPublicSwimmingPoolTrait(ctx, &x.PublicSwimmingPoolTrait)
	MapCustomToPublicSwimmingPoolTrait(ctx, &x.PublicSwimmingPoolTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPublicSwimmingPoolTrait(ctx jsonld.Context, x *schema.PublicSwimmingPoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPublicSwimmingPoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPublicSwimmingPoolTrait(ctx jsonld.Context, x *schema.PublicSwimmingPoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPublicSwimmingPoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}