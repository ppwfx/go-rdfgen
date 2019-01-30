package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDaySpa strings.Replacer
var strconvDaySpa strconv.NumError

var basicDaySpaTraitMapping = map[string]func(*schema.DaySpaTrait, []string){}
var customDaySpaTraitMapping = map[string]func(*schema.DaySpaTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DaySpa", func(ctx jsonld.Context) (interface{}) {
		return NewDaySpaFromContext(ctx)
	})

}

func NewDaySpaFromContext(ctx jsonld.Context) (x schema.DaySpa) {
	x.Type = "http://schema.org/DaySpa"
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


	MapBasicToDaySpaTrait(ctx, &x.DaySpaTrait)
	MapCustomToDaySpaTrait(ctx, &x.DaySpaTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDaySpaTrait(ctx jsonld.Context, x *schema.DaySpaTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDaySpaTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDaySpaTrait(ctx jsonld.Context, x *schema.DaySpaTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDaySpaTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}