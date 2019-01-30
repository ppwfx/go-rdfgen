package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsImageGallery strings.Replacer
var strconvImageGallery strconv.NumError

var basicImageGalleryTraitMapping = map[string]func(*schema.ImageGalleryTrait, []string){}
var customImageGalleryTraitMapping = map[string]func(*schema.ImageGalleryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ImageGallery", func(ctx jsonld.Context) (interface{}) {
		return NewImageGalleryFromContext(ctx)
	})

}

func NewImageGalleryFromContext(ctx jsonld.Context) (x schema.ImageGallery) {
	x.Type = "http://schema.org/ImageGallery"
	MapBasicToCollectionPageTrait(ctx, &x.CollectionPageTrait)
	MapCustomToCollectionPageTrait(ctx, &x.CollectionPageTrait)

	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToImageGalleryTrait(ctx, &x.ImageGalleryTrait)
	MapCustomToImageGalleryTrait(ctx, &x.ImageGalleryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToImageGalleryTrait(ctx jsonld.Context, x *schema.ImageGalleryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicImageGalleryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToImageGalleryTrait(ctx jsonld.Context, x *schema.ImageGalleryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customImageGalleryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}