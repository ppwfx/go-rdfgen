package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCasino strings.Replacer
var strconvCasino strconv.NumError

var basicCasinoTraitMapping = map[string]func(*schema.CasinoTrait, []string){}
var customCasinoTraitMapping = map[string]func(*schema.CasinoTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Casino", func(ctx jsonld.Context) (interface{}) {
		return NewCasinoFromContext(ctx)
	})

}

func NewCasinoFromContext(ctx jsonld.Context) (x schema.Casino) {
	x.Type = "http://schema.org/Casino"
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


	MapBasicToCasinoTrait(ctx, &x.CasinoTrait)
	MapCustomToCasinoTrait(ctx, &x.CasinoTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCasinoTrait(ctx jsonld.Context, x *schema.CasinoTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCasinoTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCasinoTrait(ctx jsonld.Context, x *schema.CasinoTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCasinoTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}