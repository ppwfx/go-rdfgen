package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicReleaseFormatType strings.Replacer
var strconvMusicReleaseFormatType strconv.NumError

var basicMusicReleaseFormatTypeTraitMapping = map[string]func(*schema.MusicReleaseFormatTypeTrait, []string){}
var customMusicReleaseFormatTypeTraitMapping = map[string]func(*schema.MusicReleaseFormatTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicReleaseFormatType", func(ctx jsonld.Context) (interface{}) {
		return NewMusicReleaseFormatTypeFromContext(ctx)
	})

}

func NewMusicReleaseFormatTypeFromContext(ctx jsonld.Context) (x schema.MusicReleaseFormatType) {
	x.Type = "http://schema.org/MusicReleaseFormatType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicReleaseFormatTypeTrait(ctx, &x.MusicReleaseFormatTypeTrait)
	MapCustomToMusicReleaseFormatTypeTrait(ctx, &x.MusicReleaseFormatTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicReleaseFormatTypeTrait(ctx jsonld.Context, x *schema.MusicReleaseFormatTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicReleaseFormatTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicReleaseFormatTypeTrait(ctx jsonld.Context, x *schema.MusicReleaseFormatTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicReleaseFormatTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}