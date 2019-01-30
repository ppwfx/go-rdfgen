package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDanceEvent strings.Replacer
var strconvDanceEvent strconv.NumError

var basicDanceEventTraitMapping = map[string]func(*schema.DanceEventTrait, []string){}
var customDanceEventTraitMapping = map[string]func(*schema.DanceEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DanceEvent", func(ctx jsonld.Context) (interface{}) {
		return NewDanceEventFromContext(ctx)
	})

}

func NewDanceEventFromContext(ctx jsonld.Context) (x schema.DanceEvent) {
	x.Type = "http://schema.org/DanceEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDanceEventTrait(ctx, &x.DanceEventTrait)
	MapCustomToDanceEventTrait(ctx, &x.DanceEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDanceEventTrait(ctx jsonld.Context, x *schema.DanceEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDanceEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDanceEventTrait(ctx jsonld.Context, x *schema.DanceEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDanceEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}