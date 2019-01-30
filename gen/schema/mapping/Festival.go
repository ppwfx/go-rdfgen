package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFestival strings.Replacer
var strconvFestival strconv.NumError

var basicFestivalTraitMapping = map[string]func(*schema.FestivalTrait, []string){}
var customFestivalTraitMapping = map[string]func(*schema.FestivalTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Festival", func(ctx jsonld.Context) (interface{}) {
		return NewFestivalFromContext(ctx)
	})

}

func NewFestivalFromContext(ctx jsonld.Context) (x schema.Festival) {
	x.Type = "http://schema.org/Festival"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFestivalTrait(ctx, &x.FestivalTrait)
	MapCustomToFestivalTrait(ctx, &x.FestivalTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFestivalTrait(ctx jsonld.Context, x *schema.FestivalTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFestivalTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFestivalTrait(ctx jsonld.Context, x *schema.FestivalTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFestivalTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}