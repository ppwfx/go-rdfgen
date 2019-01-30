package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRadioEpisode strings.Replacer
var strconvRadioEpisode strconv.NumError

var basicRadioEpisodeTraitMapping = map[string]func(*schema.RadioEpisodeTrait, []string){}
var customRadioEpisodeTraitMapping = map[string]func(*schema.RadioEpisodeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RadioEpisode", func(ctx jsonld.Context) (interface{}) {
		return NewRadioEpisodeFromContext(ctx)
	})

}

func NewRadioEpisodeFromContext(ctx jsonld.Context) (x schema.RadioEpisode) {
	x.Type = "http://schema.org/RadioEpisode"
	MapBasicToEpisodeTrait(ctx, &x.EpisodeTrait)
	MapCustomToEpisodeTrait(ctx, &x.EpisodeTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRadioEpisodeTrait(ctx, &x.RadioEpisodeTrait)
	MapCustomToRadioEpisodeTrait(ctx, &x.RadioEpisodeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRadioEpisodeTrait(ctx jsonld.Context, x *schema.RadioEpisodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRadioEpisodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRadioEpisodeTrait(ctx jsonld.Context, x *schema.RadioEpisodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRadioEpisodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}