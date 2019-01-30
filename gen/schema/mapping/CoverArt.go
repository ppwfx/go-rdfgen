package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCoverArt strings.Replacer
var strconvCoverArt strconv.NumError

var basicCoverArtTraitMapping = map[string]func(*schema.CoverArtTrait, []string){}
var customCoverArtTraitMapping = map[string]func(*schema.CoverArtTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CoverArt", func(ctx jsonld.Context) (interface{}) {
		return NewCoverArtFromContext(ctx)
	})

}

func NewCoverArtFromContext(ctx jsonld.Context) (x schema.CoverArt) {
	x.Type = "http://schema.org/CoverArt"
	MapBasicToVisualArtworkTrait(ctx, &x.VisualArtworkTrait)
	MapCustomToVisualArtworkTrait(ctx, &x.VisualArtworkTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCoverArtTrait(ctx, &x.CoverArtTrait)
	MapCustomToCoverArtTrait(ctx, &x.CoverArtTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCoverArtTrait(ctx jsonld.Context, x *schema.CoverArtTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCoverArtTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCoverArtTrait(ctx jsonld.Context, x *schema.CoverArtTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCoverArtTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}