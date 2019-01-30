package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMovieClip strings.Replacer
var strconvMovieClip strconv.NumError

var basicMovieClipTraitMapping = map[string]func(*schema.MovieClipTrait, []string){}
var customMovieClipTraitMapping = map[string]func(*schema.MovieClipTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MovieClip", func(ctx jsonld.Context) (interface{}) {
		return NewMovieClipFromContext(ctx)
	})

}

func NewMovieClipFromContext(ctx jsonld.Context) (x schema.MovieClip) {
	x.Type = "http://schema.org/MovieClip"
	MapBasicToClipTrait(ctx, &x.ClipTrait)
	MapCustomToClipTrait(ctx, &x.ClipTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMovieClipTrait(ctx, &x.MovieClipTrait)
	MapCustomToMovieClipTrait(ctx, &x.MovieClipTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMovieClipTrait(ctx jsonld.Context, x *schema.MovieClipTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMovieClipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMovieClipTrait(ctx jsonld.Context, x *schema.MovieClipTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMovieClipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}