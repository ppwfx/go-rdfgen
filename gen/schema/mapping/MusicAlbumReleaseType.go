package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicAlbumReleaseType strings.Replacer
var strconvMusicAlbumReleaseType strconv.NumError

var basicMusicAlbumReleaseTypeTraitMapping = map[string]func(*schema.MusicAlbumReleaseTypeTrait, []string){}
var customMusicAlbumReleaseTypeTraitMapping = map[string]func(*schema.MusicAlbumReleaseTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicAlbumReleaseType", func(ctx jsonld.Context) (interface{}) {
		return NewMusicAlbumReleaseTypeFromContext(ctx)
	})

}

func NewMusicAlbumReleaseTypeFromContext(ctx jsonld.Context) (x schema.MusicAlbumReleaseType) {
	x.Type = "http://schema.org/MusicAlbumReleaseType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicAlbumReleaseTypeTrait(ctx, &x.MusicAlbumReleaseTypeTrait)
	MapCustomToMusicAlbumReleaseTypeTrait(ctx, &x.MusicAlbumReleaseTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicAlbumReleaseTypeTrait(ctx jsonld.Context, x *schema.MusicAlbumReleaseTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicAlbumReleaseTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicAlbumReleaseTypeTrait(ctx jsonld.Context, x *schema.MusicAlbumReleaseTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicAlbumReleaseTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}