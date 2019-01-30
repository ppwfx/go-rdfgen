package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRadioSeason strings.Replacer
var strconvRadioSeason strconv.NumError

var basicRadioSeasonTraitMapping = map[string]func(*schema.RadioSeasonTrait, []string){}
var customRadioSeasonTraitMapping = map[string]func(*schema.RadioSeasonTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RadioSeason", func(ctx jsonld.Context) (interface{}) {
		return NewRadioSeasonFromContext(ctx)
	})

}

func NewRadioSeasonFromContext(ctx jsonld.Context) (x schema.RadioSeason) {
	x.Type = "http://schema.org/RadioSeason"
	MapBasicToCreativeWorkSeasonTrait(ctx, &x.CreativeWorkSeasonTrait)
	MapCustomToCreativeWorkSeasonTrait(ctx, &x.CreativeWorkSeasonTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRadioSeasonTrait(ctx, &x.RadioSeasonTrait)
	MapCustomToRadioSeasonTrait(ctx, &x.RadioSeasonTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRadioSeasonTrait(ctx jsonld.Context, x *schema.RadioSeasonTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRadioSeasonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRadioSeasonTrait(ctx jsonld.Context, x *schema.RadioSeasonTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRadioSeasonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}