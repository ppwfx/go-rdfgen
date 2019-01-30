package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVideoGameClip strings.Replacer
var strconvVideoGameClip strconv.NumError

var basicVideoGameClipTraitMapping = map[string]func(*schema.VideoGameClipTrait, []string){}
var customVideoGameClipTraitMapping = map[string]func(*schema.VideoGameClipTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/VideoGameClip", func(ctx jsonld.Context) (interface{}) {
		return NewVideoGameClipFromContext(ctx)
	})

}

func NewVideoGameClipFromContext(ctx jsonld.Context) (x schema.VideoGameClip) {
	x.Type = "http://schema.org/VideoGameClip"
	MapBasicToClipTrait(ctx, &x.ClipTrait)
	MapCustomToClipTrait(ctx, &x.ClipTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVideoGameClipTrait(ctx, &x.VideoGameClipTrait)
	MapCustomToVideoGameClipTrait(ctx, &x.VideoGameClipTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVideoGameClipTrait(ctx jsonld.Context, x *schema.VideoGameClipTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVideoGameClipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVideoGameClipTrait(ctx jsonld.Context, x *schema.VideoGameClipTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVideoGameClipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}