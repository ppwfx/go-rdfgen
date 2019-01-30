package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNightClub strings.Replacer
var strconvNightClub strconv.NumError

var basicNightClubTraitMapping = map[string]func(*schema.NightClubTrait, []string){}
var customNightClubTraitMapping = map[string]func(*schema.NightClubTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/NightClub", func(ctx jsonld.Context) (interface{}) {
		return NewNightClubFromContext(ctx)
	})

}

func NewNightClubFromContext(ctx jsonld.Context) (x schema.NightClub) {
	x.Type = "http://schema.org/NightClub"
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


	MapBasicToNightClubTrait(ctx, &x.NightClubTrait)
	MapCustomToNightClubTrait(ctx, &x.NightClubTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNightClubTrait(ctx jsonld.Context, x *schema.NightClubTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNightClubTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNightClubTrait(ctx jsonld.Context, x *schema.NightClubTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNightClubTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}