package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNewspaper strings.Replacer
var strconvNewspaper strconv.NumError

var basicNewspaperTraitMapping = map[string]func(*schema.NewspaperTrait, []string){}
var customNewspaperTraitMapping = map[string]func(*schema.NewspaperTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Newspaper", func(ctx jsonld.Context) (interface{}) {
		return NewNewspaperFromContext(ctx)
	})

}

func NewNewspaperFromContext(ctx jsonld.Context) (x schema.Newspaper) {
	x.Type = "http://schema.org/Newspaper"
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


	MapBasicToNewspaperTrait(ctx, &x.NewspaperTrait)
	MapCustomToNewspaperTrait(ctx, &x.NewspaperTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNewspaperTrait(ctx jsonld.Context, x *schema.NewspaperTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNewspaperTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNewspaperTrait(ctx jsonld.Context, x *schema.NewspaperTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNewspaperTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}