package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEventSeries strings.Replacer
var strconvEventSeries strconv.NumError

var basicEventSeriesTraitMapping = map[string]func(*schema.EventSeriesTrait, []string){}
var customEventSeriesTraitMapping = map[string]func(*schema.EventSeriesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EventSeries", func(ctx jsonld.Context) (interface{}) {
		return NewEventSeriesFromContext(ctx)
	})

}

func NewEventSeriesFromContext(ctx jsonld.Context) (x schema.EventSeries) {
	x.Type = "http://schema.org/EventSeries"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToSeriesTrait(ctx, &x.SeriesTrait)
	MapCustomToSeriesTrait(ctx, &x.SeriesTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)


	MapBasicToEventSeriesTrait(ctx, &x.EventSeriesTrait)
	MapCustomToEventSeriesTrait(ctx, &x.EventSeriesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEventSeriesTrait(ctx jsonld.Context, x *schema.EventSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEventSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEventSeriesTrait(ctx jsonld.Context, x *schema.EventSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEventSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}