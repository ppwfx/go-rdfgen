package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicVenue strings.Replacer
var strconvMusicVenue strconv.NumError

var basicMusicVenueTraitMapping = map[string]func(*schema.MusicVenueTrait, []string){}
var customMusicVenueTraitMapping = map[string]func(*schema.MusicVenueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicVenue", func(ctx jsonld.Context) (interface{}) {
		return NewMusicVenueFromContext(ctx)
	})

}

func NewMusicVenueFromContext(ctx jsonld.Context) (x schema.MusicVenue) {
	x.Type = "http://schema.org/MusicVenue"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicVenueTrait(ctx, &x.MusicVenueTrait)
	MapCustomToMusicVenueTrait(ctx, &x.MusicVenueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicVenueTrait(ctx jsonld.Context, x *schema.MusicVenueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicVenueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicVenueTrait(ctx jsonld.Context, x *schema.MusicVenueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicVenueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}