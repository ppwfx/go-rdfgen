package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsComicCoverArt strings.Replacer
var strconvComicCoverArt strconv.NumError

var basicComicCoverArtTraitMapping = map[string]func(*schema.ComicCoverArtTrait, []string){}
var customComicCoverArtTraitMapping = map[string]func(*schema.ComicCoverArtTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ComicCoverArt", func(ctx jsonld.Context) (interface{}) {
		return NewComicCoverArtFromContext(ctx)
	})

}

func NewComicCoverArtFromContext(ctx jsonld.Context) (x schema.ComicCoverArt) {
	x.Type = "http://schema.org/ComicCoverArt"
	MapBasicToComicStoryTrait(ctx, &x.ComicStoryTrait)
	MapCustomToComicStoryTrait(ctx, &x.ComicStoryTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToCoverArtTrait(ctx, &x.CoverArtTrait)
	MapCustomToCoverArtTrait(ctx, &x.CoverArtTrait)

	MapBasicToVisualArtworkTrait(ctx, &x.VisualArtworkTrait)
	MapCustomToVisualArtworkTrait(ctx, &x.VisualArtworkTrait)


	MapBasicToComicCoverArtTrait(ctx, &x.ComicCoverArtTrait)
	MapCustomToComicCoverArtTrait(ctx, &x.ComicCoverArtTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToComicCoverArtTrait(ctx jsonld.Context, x *schema.ComicCoverArtTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicComicCoverArtTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToComicCoverArtTrait(ctx jsonld.Context, x *schema.ComicCoverArtTrait) () {
	for k, v := range ctx.Current {
		f, ok := customComicCoverArtTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}