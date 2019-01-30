package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTVEpisode strings.Replacer
var strconvTVEpisode strconv.NumError

var basicTVEpisodeTraitMapping = map[string]func(*schema.TVEpisodeTrait, []string){}
var customTVEpisodeTraitMapping = map[string]func(*schema.TVEpisodeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TVEpisode", func(ctx jsonld.Context) (interface{}) {
		return NewTVEpisodeFromContext(ctx)
	})





	customTVEpisodeTraitMapping["http://schema.org/countryOfOrigin"] = func(x *schema.TVEpisodeTrait, ctx jsonld.Context, s string){
		var y schema.Country
		if strings.HasPrefix(s, "_:") {
			y = NewCountryFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCountry()
			y.Id = s
		}

		x.CountryOfOrigin = &y

		return
	}

	customTVEpisodeTraitMapping["http://schema.org/subtitleLanguage"] = func(x *schema.TVEpisodeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SubtitleLanguage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SubtitleLanguage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SubtitleLanguage = s
		}
	}

	customTVEpisodeTraitMapping["http://schema.org/partOfTVSeries"] = func(x *schema.TVEpisodeTrait, ctx jsonld.Context, s string){
		var y schema.TVSeries
		if strings.HasPrefix(s, "_:") {
			y = NewTVSeriesFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewTVSeries()
			y.Id = s
		}

		x.PartOfTVSeries = &y

		return
	}
}

func NewTVEpisodeFromContext(ctx jsonld.Context) (x schema.TVEpisode) {
	x.Type = "http://schema.org/TVEpisode"
	MapBasicToEpisodeTrait(ctx, &x.EpisodeTrait)
	MapCustomToEpisodeTrait(ctx, &x.EpisodeTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTVEpisodeTrait(ctx, &x.TVEpisodeTrait)
	MapCustomToTVEpisodeTrait(ctx, &x.TVEpisodeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTVEpisodeTrait(ctx jsonld.Context, x *schema.TVEpisodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTVEpisodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTVEpisodeTrait(ctx jsonld.Context, x *schema.TVEpisodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTVEpisodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}