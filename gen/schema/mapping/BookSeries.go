package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBookSeries strings.Replacer
var strconvBookSeries strconv.NumError

var basicBookSeriesTraitMapping = map[string]func(*schema.BookSeriesTrait, []string){}
var customBookSeriesTraitMapping = map[string]func(*schema.BookSeriesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BookSeries", func(ctx jsonld.Context) (interface{}) {
		return NewBookSeriesFromContext(ctx)
	})

}

func NewBookSeriesFromContext(ctx jsonld.Context) (x schema.BookSeries) {
	x.Type = "http://schema.org/BookSeries"
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


	MapBasicToBookSeriesTrait(ctx, &x.BookSeriesTrait)
	MapCustomToBookSeriesTrait(ctx, &x.BookSeriesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBookSeriesTrait(ctx jsonld.Context, x *schema.BookSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBookSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBookSeriesTrait(ctx jsonld.Context, x *schema.BookSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBookSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}