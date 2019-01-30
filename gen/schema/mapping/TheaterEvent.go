package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTheaterEvent strings.Replacer
var strconvTheaterEvent strconv.NumError

var basicTheaterEventTraitMapping = map[string]func(*schema.TheaterEventTrait, []string){}
var customTheaterEventTraitMapping = map[string]func(*schema.TheaterEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TheaterEvent", func(ctx jsonld.Context) (interface{}) {
		return NewTheaterEventFromContext(ctx)
	})

}

func NewTheaterEventFromContext(ctx jsonld.Context) (x schema.TheaterEvent) {
	x.Type = "http://schema.org/TheaterEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTheaterEventTrait(ctx, &x.TheaterEventTrait)
	MapCustomToTheaterEventTrait(ctx, &x.TheaterEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTheaterEventTrait(ctx jsonld.Context, x *schema.TheaterEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTheaterEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTheaterEventTrait(ctx jsonld.Context, x *schema.TheaterEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTheaterEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}