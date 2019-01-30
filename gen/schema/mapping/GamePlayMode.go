package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGamePlayMode strings.Replacer
var strconvGamePlayMode strconv.NumError

var basicGamePlayModeTraitMapping = map[string]func(*schema.GamePlayModeTrait, []string){}
var customGamePlayModeTraitMapping = map[string]func(*schema.GamePlayModeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GamePlayMode", func(ctx jsonld.Context) (interface{}) {
		return NewGamePlayModeFromContext(ctx)
	})

}

func NewGamePlayModeFromContext(ctx jsonld.Context) (x schema.GamePlayMode) {
	x.Type = "http://schema.org/GamePlayMode"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGamePlayModeTrait(ctx, &x.GamePlayModeTrait)
	MapCustomToGamePlayModeTrait(ctx, &x.GamePlayModeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGamePlayModeTrait(ctx jsonld.Context, x *schema.GamePlayModeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGamePlayModeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGamePlayModeTrait(ctx jsonld.Context, x *schema.GamePlayModeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGamePlayModeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}