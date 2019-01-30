package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTVSeason strings.Replacer
var strconvTVSeason strconv.NumError

var basicTVSeasonTraitMapping = map[string]func(*schema.TVSeasonTrait, []string){}
var customTVSeasonTraitMapping = map[string]func(*schema.TVSeasonTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TVSeason", func(ctx jsonld.Context) (interface{}) {
		return NewTVSeasonFromContext(ctx)
	})




	customTVSeasonTraitMapping["http://schema.org/countryOfOrigin"] = func(x *schema.TVSeasonTrait, ctx jsonld.Context, s string){
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

	customTVSeasonTraitMapping["http://schema.org/partOfTVSeries"] = func(x *schema.TVSeasonTrait, ctx jsonld.Context, s string){
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

func NewTVSeasonFromContext(ctx jsonld.Context) (x schema.TVSeason) {
	x.Type = "http://schema.org/TVSeason"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToCreativeWorkSeasonTrait(ctx, &x.CreativeWorkSeasonTrait)
	MapCustomToCreativeWorkSeasonTrait(ctx, &x.CreativeWorkSeasonTrait)


	MapBasicToTVSeasonTrait(ctx, &x.TVSeasonTrait)
	MapCustomToTVSeasonTrait(ctx, &x.TVSeasonTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTVSeasonTrait(ctx jsonld.Context, x *schema.TVSeasonTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTVSeasonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTVSeasonTrait(ctx jsonld.Context, x *schema.TVSeasonTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTVSeasonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}