package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsArtGallery strings.Replacer
var strconvArtGallery strconv.NumError

var basicArtGalleryTraitMapping = map[string]func(*schema.ArtGalleryTrait, []string){}
var customArtGalleryTraitMapping = map[string]func(*schema.ArtGalleryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ArtGallery", func(ctx jsonld.Context) (interface{}) {
		return NewArtGalleryFromContext(ctx)
	})

}

func NewArtGalleryFromContext(ctx jsonld.Context) (x schema.ArtGallery) {
	x.Type = "http://schema.org/ArtGallery"
	MapBasicToEntertainmentBusinessTrait(ctx, &x.EntertainmentBusinessTrait)
	MapCustomToEntertainmentBusinessTrait(ctx, &x.EntertainmentBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToArtGalleryTrait(ctx, &x.ArtGalleryTrait)
	MapCustomToArtGalleryTrait(ctx, &x.ArtGalleryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToArtGalleryTrait(ctx jsonld.Context, x *schema.ArtGalleryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicArtGalleryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToArtGalleryTrait(ctx jsonld.Context, x *schema.ArtGalleryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customArtGalleryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}