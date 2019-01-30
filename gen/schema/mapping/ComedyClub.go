package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsComedyClub strings.Replacer
var strconvComedyClub strconv.NumError

var basicComedyClubTraitMapping = map[string]func(*schema.ComedyClubTrait, []string){}
var customComedyClubTraitMapping = map[string]func(*schema.ComedyClubTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ComedyClub", func(ctx jsonld.Context) (interface{}) {
		return NewComedyClubFromContext(ctx)
	})

}

func NewComedyClubFromContext(ctx jsonld.Context) (x schema.ComedyClub) {
	x.Type = "http://schema.org/ComedyClub"
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


	MapBasicToComedyClubTrait(ctx, &x.ComedyClubTrait)
	MapCustomToComedyClubTrait(ctx, &x.ComedyClubTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToComedyClubTrait(ctx jsonld.Context, x *schema.ComedyClubTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicComedyClubTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToComedyClubTrait(ctx jsonld.Context, x *schema.ComedyClubTrait) () {
	for k, v := range ctx.Current {
		f, ok := customComedyClubTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}