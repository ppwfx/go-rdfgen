package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSportsActivityLocation strings.Replacer
var strconvSportsActivityLocation strconv.NumError

var basicSportsActivityLocationTraitMapping = map[string]func(*schema.SportsActivityLocationTrait, []string){}
var customSportsActivityLocationTraitMapping = map[string]func(*schema.SportsActivityLocationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SportsActivityLocation", func(ctx jsonld.Context) (interface{}) {
		return NewSportsActivityLocationFromContext(ctx)
	})

}

func NewSportsActivityLocationFromContext(ctx jsonld.Context) (x schema.SportsActivityLocation) {
	x.Type = "http://schema.org/SportsActivityLocation"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToSportsActivityLocationTrait(ctx, &x.SportsActivityLocationTrait)
	MapCustomToSportsActivityLocationTrait(ctx, &x.SportsActivityLocationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSportsActivityLocationTrait(ctx jsonld.Context, x *schema.SportsActivityLocationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSportsActivityLocationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSportsActivityLocationTrait(ctx jsonld.Context, x *schema.SportsActivityLocationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSportsActivityLocationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}