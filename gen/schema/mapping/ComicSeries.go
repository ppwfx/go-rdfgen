package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsComicSeries strings.Replacer
var strconvComicSeries strconv.NumError

var basicComicSeriesTraitMapping = map[string]func(*schema.ComicSeriesTrait, []string){}
var customComicSeriesTraitMapping = map[string]func(*schema.ComicSeriesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ComicSeries", func(ctx jsonld.Context) (interface{}) {
		return NewComicSeriesFromContext(ctx)
	})

}

func NewComicSeriesFromContext(ctx jsonld.Context) (x schema.ComicSeries) {
	x.Type = "http://schema.org/ComicSeries"
	MapBasicToPeriodicalTrait(ctx, &x.PeriodicalTrait)
	MapCustomToPeriodicalTrait(ctx, &x.PeriodicalTrait)

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


	MapBasicToComicSeriesTrait(ctx, &x.ComicSeriesTrait)
	MapCustomToComicSeriesTrait(ctx, &x.ComicSeriesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToComicSeriesTrait(ctx jsonld.Context, x *schema.ComicSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicComicSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToComicSeriesTrait(ctx jsonld.Context, x *schema.ComicSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customComicSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}