package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicAlbumProductionType strings.Replacer
var strconvMusicAlbumProductionType strconv.NumError

var basicMusicAlbumProductionTypeTraitMapping = map[string]func(*schema.MusicAlbumProductionTypeTrait, []string){}
var customMusicAlbumProductionTypeTraitMapping = map[string]func(*schema.MusicAlbumProductionTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicAlbumProductionType", func(ctx jsonld.Context) (interface{}) {
		return NewMusicAlbumProductionTypeFromContext(ctx)
	})

}

func NewMusicAlbumProductionTypeFromContext(ctx jsonld.Context) (x schema.MusicAlbumProductionType) {
	x.Type = "http://schema.org/MusicAlbumProductionType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicAlbumProductionTypeTrait(ctx, &x.MusicAlbumProductionTypeTrait)
	MapCustomToMusicAlbumProductionTypeTrait(ctx, &x.MusicAlbumProductionTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicAlbumProductionTypeTrait(ctx jsonld.Context, x *schema.MusicAlbumProductionTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicAlbumProductionTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicAlbumProductionTypeTrait(ctx jsonld.Context, x *schema.MusicAlbumProductionTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicAlbumProductionTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}