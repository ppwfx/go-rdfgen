package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAudioObject strings.Replacer
var strconvAudioObject strconv.NumError

var basicAudioObjectTraitMapping = map[string]func(*schema.AudioObjectTrait, []string){}
var customAudioObjectTraitMapping = map[string]func(*schema.AudioObjectTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AudioObject", func(ctx jsonld.Context) (interface{}) {
		return NewAudioObjectFromContext(ctx)
	})


	basicAudioObjectTraitMapping["http://schema.org/transcript"] = func(x *schema.AudioObjectTrait, s []string) {
		x.Transcript = s[0]
	}

}

func NewAudioObjectFromContext(ctx jsonld.Context) (x schema.AudioObject) {
	x.Type = "http://schema.org/AudioObject"
	MapBasicToMediaObjectTrait(ctx, &x.MediaObjectTrait)
	MapCustomToMediaObjectTrait(ctx, &x.MediaObjectTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAudioObjectTrait(ctx, &x.AudioObjectTrait)
	MapCustomToAudioObjectTrait(ctx, &x.AudioObjectTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAudioObjectTrait(ctx jsonld.Context, x *schema.AudioObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAudioObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAudioObjectTrait(ctx jsonld.Context, x *schema.AudioObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAudioObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}