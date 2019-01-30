package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicVideoObject strings.Replacer
var strconvMusicVideoObject strconv.NumError

var basicMusicVideoObjectTraitMapping = map[string]func(*schema.MusicVideoObjectTrait, []string){}
var customMusicVideoObjectTraitMapping = map[string]func(*schema.MusicVideoObjectTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicVideoObject", func(ctx jsonld.Context) (interface{}) {
		return NewMusicVideoObjectFromContext(ctx)
	})

}

func NewMusicVideoObjectFromContext(ctx jsonld.Context) (x schema.MusicVideoObject) {
	x.Type = "http://schema.org/MusicVideoObject"
	MapBasicToMediaObjectTrait(ctx, &x.MediaObjectTrait)
	MapCustomToMediaObjectTrait(ctx, &x.MediaObjectTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicVideoObjectTrait(ctx, &x.MusicVideoObjectTrait)
	MapCustomToMusicVideoObjectTrait(ctx, &x.MusicVideoObjectTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicVideoObjectTrait(ctx jsonld.Context, x *schema.MusicVideoObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicVideoObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicVideoObjectTrait(ctx jsonld.Context, x *schema.MusicVideoObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicVideoObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}