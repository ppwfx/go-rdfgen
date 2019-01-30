package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVideoGallery strings.Replacer
var strconvVideoGallery strconv.NumError

var basicVideoGalleryTraitMapping = map[string]func(*schema.VideoGalleryTrait, []string){}
var customVideoGalleryTraitMapping = map[string]func(*schema.VideoGalleryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/VideoGallery", func(ctx jsonld.Context) (interface{}) {
		return NewVideoGalleryFromContext(ctx)
	})

}

func NewVideoGalleryFromContext(ctx jsonld.Context) (x schema.VideoGallery) {
	x.Type = "http://schema.org/VideoGallery"
	MapBasicToCollectionPageTrait(ctx, &x.CollectionPageTrait)
	MapCustomToCollectionPageTrait(ctx, &x.CollectionPageTrait)

	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVideoGalleryTrait(ctx, &x.VideoGalleryTrait)
	MapCustomToVideoGalleryTrait(ctx, &x.VideoGalleryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVideoGalleryTrait(ctx jsonld.Context, x *schema.VideoGalleryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVideoGalleryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVideoGalleryTrait(ctx jsonld.Context, x *schema.VideoGalleryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVideoGalleryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}