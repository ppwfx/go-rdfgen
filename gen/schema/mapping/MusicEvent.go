package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicEvent strings.Replacer
var strconvMusicEvent strconv.NumError

var basicMusicEventTraitMapping = map[string]func(*schema.MusicEventTrait, []string){}
var customMusicEventTraitMapping = map[string]func(*schema.MusicEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicEvent", func(ctx jsonld.Context) (interface{}) {
		return NewMusicEventFromContext(ctx)
	})

}

func NewMusicEventFromContext(ctx jsonld.Context) (x schema.MusicEvent) {
	x.Type = "http://schema.org/MusicEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicEventTrait(ctx, &x.MusicEventTrait)
	MapCustomToMusicEventTrait(ctx, &x.MusicEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicEventTrait(ctx jsonld.Context, x *schema.MusicEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicEventTrait(ctx jsonld.Context, x *schema.MusicEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}