package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEventVenue strings.Replacer
var strconvEventVenue strconv.NumError

var basicEventVenueTraitMapping = map[string]func(*schema.EventVenueTrait, []string){}
var customEventVenueTraitMapping = map[string]func(*schema.EventVenueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EventVenue", func(ctx jsonld.Context) (interface{}) {
		return NewEventVenueFromContext(ctx)
	})

}

func NewEventVenueFromContext(ctx jsonld.Context) (x schema.EventVenue) {
	x.Type = "http://schema.org/EventVenue"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEventVenueTrait(ctx, &x.EventVenueTrait)
	MapCustomToEventVenueTrait(ctx, &x.EventVenueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEventVenueTrait(ctx jsonld.Context, x *schema.EventVenueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEventVenueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEventVenueTrait(ctx jsonld.Context, x *schema.EventVenueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEventVenueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}