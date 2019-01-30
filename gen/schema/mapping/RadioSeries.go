package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRadioSeries strings.Replacer
var strconvRadioSeries strconv.NumError

var basicRadioSeriesTraitMapping = map[string]func(*schema.RadioSeriesTrait, []string){}
var customRadioSeriesTraitMapping = map[string]func(*schema.RadioSeriesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RadioSeries", func(ctx jsonld.Context) (interface{}) {
		return NewRadioSeriesFromContext(ctx)
	})
















	customRadioSeriesTraitMapping["http://schema.org/productionCompany"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/actor"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/director"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/actors"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/directors"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/musicBy"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/trailer"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/containsSeason"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/episode"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/episodes"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/numberOfEpisodes"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/numberOfSeasons"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/season"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

	customRadioSeriesTraitMapping["http://schema.org/seasons"] = func(x *schema.RadioSeriesTrait, ctx jsonld.Context, s string){
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

func NewRadioSeriesFromContext(ctx jsonld.Context) (x schema.RadioSeries) {
	x.Type = "http://schema.org/RadioSeries"
	MapBasicToCreativeWorkSeriesTrait(ctx, &x.CreativeWorkSeriesTrait)
	MapCustomToCreativeWorkSeriesTrait(ctx, &x.CreativeWorkSeriesTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToSeriesTrait(ctx, &x.SeriesTrait)
	MapCustomToSeriesTrait(ctx, &x.SeriesTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)


	MapBasicToRadioSeriesTrait(ctx, &x.RadioSeriesTrait)
	MapCustomToRadioSeriesTrait(ctx, &x.RadioSeriesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRadioSeriesTrait(ctx jsonld.Context, x *schema.RadioSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRadioSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRadioSeriesTrait(ctx jsonld.Context, x *schema.RadioSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRadioSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}