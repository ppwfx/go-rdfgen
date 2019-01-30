package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSportsClub strings.Replacer
var strconvSportsClub strconv.NumError

var basicSportsClubTraitMapping = map[string]func(*schema.SportsClubTrait, []string){}
var customSportsClubTraitMapping = map[string]func(*schema.SportsClubTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SportsClub", func(ctx jsonld.Context) (interface{}) {
		return NewSportsClubFromContext(ctx)
	})

}

func NewSportsClubFromContext(ctx jsonld.Context) (x schema.SportsClub) {
	x.Type = "http://schema.org/SportsClub"
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


	MapBasicToSportsClubTrait(ctx, &x.SportsClubTrait)
	MapCustomToSportsClubTrait(ctx, &x.SportsClubTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSportsClubTrait(ctx jsonld.Context, x *schema.SportsClubTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSportsClubTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSportsClubTrait(ctx jsonld.Context, x *schema.SportsClubTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSportsClubTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}