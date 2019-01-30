package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsStadiumOrArena strings.Replacer
var strconvStadiumOrArena strconv.NumError

var basicStadiumOrArenaTraitMapping = map[string]func(*schema.StadiumOrArenaTrait, []string){}
var customStadiumOrArenaTraitMapping = map[string]func(*schema.StadiumOrArenaTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/StadiumOrArena", func(ctx jsonld.Context) (interface{}) {
		return NewStadiumOrArenaFromContext(ctx)
	})

}

func NewStadiumOrArenaFromContext(ctx jsonld.Context) (x schema.StadiumOrArena) {
	x.Type = "http://schema.org/StadiumOrArena"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToSportsActivityLocationTrait(ctx, &x.SportsActivityLocationTrait)
	MapCustomToSportsActivityLocationTrait(ctx, &x.SportsActivityLocationTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToStadiumOrArenaTrait(ctx, &x.StadiumOrArenaTrait)
	MapCustomToStadiumOrArenaTrait(ctx, &x.StadiumOrArenaTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToStadiumOrArenaTrait(ctx jsonld.Context, x *schema.StadiumOrArenaTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicStadiumOrArenaTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToStadiumOrArenaTrait(ctx jsonld.Context, x *schema.StadiumOrArenaTrait) () {
	for k, v := range ctx.Current {
		f, ok := customStadiumOrArenaTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}