package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSeries strings.Replacer
var strconvSeries strconv.NumError

var basicSeriesTraitMapping = map[string]func(*schema.SeriesTrait, []string){}
var customSeriesTraitMapping = map[string]func(*schema.SeriesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Series", func(ctx jsonld.Context) (interface{}) {
		return NewSeriesFromContext(ctx)
	})

}

func NewSeriesFromContext(ctx jsonld.Context) (x schema.Series) {
	x.Type = "http://schema.org/Series"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSeriesTrait(ctx, &x.SeriesTrait)
	MapCustomToSeriesTrait(ctx, &x.SeriesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSeriesTrait(ctx jsonld.Context, x *schema.SeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSeriesTrait(ctx jsonld.Context, x *schema.SeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}