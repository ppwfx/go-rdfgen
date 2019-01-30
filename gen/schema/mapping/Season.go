package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSeason strings.Replacer
var strconvSeason strconv.NumError

var basicSeasonTraitMapping = map[string]func(*schema.SeasonTrait, []string){}
var customSeasonTraitMapping = map[string]func(*schema.SeasonTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Season", func(ctx jsonld.Context) (interface{}) {
		return NewSeasonFromContext(ctx)
	})

}

func NewSeasonFromContext(ctx jsonld.Context) (x schema.Season) {
	x.Type = "http://schema.org/Season"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSeasonTrait(ctx, &x.SeasonTrait)
	MapCustomToSeasonTrait(ctx, &x.SeasonTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSeasonTrait(ctx jsonld.Context, x *schema.SeasonTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSeasonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSeasonTrait(ctx jsonld.Context, x *schema.SeasonTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSeasonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}