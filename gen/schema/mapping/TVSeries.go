package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTVSeries strings.Replacer
var strconvTVSeries strconv.NumError

var basicTVSeriesTraitMapping = map[string]func(*schema.TVSeriesTrait, []string){}
var customTVSeriesTraitMapping = map[string]func(*schema.TVSeriesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TVSeries", func(ctx jsonld.Context) (interface{}) {
		return NewTVSeriesFromContext(ctx)
	})

















	customTVSeriesTraitMapping["http://schema.org/productionCompany"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.ProductionCompany = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/actor"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Actor = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/director"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Director = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/actors"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Actors = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/directors"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Directors = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/musicBy"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MusicBy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MusicBy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MusicBy = s
		}
	}

	customTVSeriesTraitMapping["http://schema.org/trailer"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.VideoObject
		if strings.HasPrefix(s, "_:") {
			y = NewVideoObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewVideoObject()
			y.Id = s
		}

		x.Trailer = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/containsSeason"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWorkSeason
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkSeasonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWorkSeason()
			y.Id = s
		}

		x.ContainsSeason = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/countryOfOrigin"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
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

	customTVSeriesTraitMapping["http://schema.org/episode"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Episode
		if strings.HasPrefix(s, "_:") {
			y = NewEpisodeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEpisode()
			y.Id = s
		}

		x.Episode = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/episodes"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Episode
		if strings.HasPrefix(s, "_:") {
			y = NewEpisodeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEpisode()
			y.Id = s
		}

		x.Episodes = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/numberOfEpisodes"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.NumberOfEpisodes = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/numberOfSeasons"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.NumberOfSeasons = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/season"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWorkSeason
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkSeasonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWorkSeason()
			y.Id = s
		}

		x.Season = &y

		return
	}

	customTVSeriesTraitMapping["http://schema.org/seasons"] = func(x *schema.TVSeriesTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWorkSeason
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkSeasonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWorkSeason()
			y.Id = s
		}

		x.Seasons = &y

		return
	}
}

func NewTVSeriesFromContext(ctx jsonld.Context) (x schema.TVSeries) {
	x.Type = "http://schema.org/TVSeries"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToCreativeWorkSeriesTrait(ctx, &x.CreativeWorkSeriesTrait)
	MapCustomToCreativeWorkSeriesTrait(ctx, &x.CreativeWorkSeriesTrait)

	MapBasicToSeriesTrait(ctx, &x.SeriesTrait)
	MapCustomToSeriesTrait(ctx, &x.SeriesTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)


	MapBasicToTVSeriesTrait(ctx, &x.TVSeriesTrait)
	MapCustomToTVSeriesTrait(ctx, &x.TVSeriesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTVSeriesTrait(ctx jsonld.Context, x *schema.TVSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTVSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTVSeriesTrait(ctx jsonld.Context, x *schema.TVSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTVSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}