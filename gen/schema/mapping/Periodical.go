package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPeriodical strings.Replacer
var strconvPeriodical strconv.NumError

var basicPeriodicalTraitMapping = map[string]func(*schema.PeriodicalTrait, []string){}
var customPeriodicalTraitMapping = map[string]func(*schema.PeriodicalTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Periodical", func(ctx jsonld.Context) (interface{}) {
		return NewPeriodicalFromContext(ctx)
	})

}

func NewPeriodicalFromContext(ctx jsonld.Context) (x schema.Periodical) {
	x.Type = "http://schema.org/Periodical"
	MapBasicToCreativeWorkSeriesTrait(ctx, &x.CreativeWorkSeriesTrait)
	MapCustomToCreativeWorkSeriesTrait(ctx, &x.CreativeWorkSeriesTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToSeriesTrait(ctx, &x.SeriesTrait)
	MapCustomToSeriesTrait(ctx, &x.SeriesTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)


	MapBasicToPeriodicalTrait(ctx, &x.PeriodicalTrait)
	MapCustomToPeriodicalTrait(ctx, &x.PeriodicalTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPeriodicalTrait(ctx jsonld.Context, x *schema.PeriodicalTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPeriodicalTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPeriodicalTrait(ctx jsonld.Context, x *schema.PeriodicalTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPeriodicalTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}