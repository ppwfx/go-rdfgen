package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicStore strings.Replacer
var strconvMusicStore strconv.NumError

var basicMusicStoreTraitMapping = map[string]func(*schema.MusicStoreTrait, []string){}
var customMusicStoreTraitMapping = map[string]func(*schema.MusicStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicStore", func(ctx jsonld.Context) (interface{}) {
		return NewMusicStoreFromContext(ctx)
	})

}

func NewMusicStoreFromContext(ctx jsonld.Context) (x schema.MusicStore) {
	x.Type = "http://schema.org/MusicStore"
	MapBasicToStoreTrait(ctx, &x.StoreTrait)
	MapCustomToStoreTrait(ctx, &x.StoreTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToMusicStoreTrait(ctx, &x.MusicStoreTrait)
	MapCustomToMusicStoreTrait(ctx, &x.MusicStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicStoreTrait(ctx jsonld.Context, x *schema.MusicStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicStoreTrait(ctx jsonld.Context, x *schema.MusicStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}