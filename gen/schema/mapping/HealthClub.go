package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHealthClub strings.Replacer
var strconvHealthClub strconv.NumError

var basicHealthClubTraitMapping = map[string]func(*schema.HealthClubTrait, []string){}
var customHealthClubTraitMapping = map[string]func(*schema.HealthClubTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HealthClub", func(ctx jsonld.Context) (interface{}) {
		return NewHealthClubFromContext(ctx)
	})

}

func NewHealthClubFromContext(ctx jsonld.Context) (x schema.HealthClub) {
	x.Type = "http://schema.org/HealthClub"
	MapBasicToHealthAndBeautyBusinessTrait(ctx, &x.HealthAndBeautyBusinessTrait)
	MapCustomToHealthAndBeautyBusinessTrait(ctx, &x.HealthAndBeautyBusinessTrait)

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


	MapBasicToHealthClubTrait(ctx, &x.HealthClubTrait)
	MapCustomToHealthClubTrait(ctx, &x.HealthClubTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHealthClubTrait(ctx jsonld.Context, x *schema.HealthClubTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHealthClubTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHealthClubTrait(ctx jsonld.Context, x *schema.HealthClubTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHealthClubTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}