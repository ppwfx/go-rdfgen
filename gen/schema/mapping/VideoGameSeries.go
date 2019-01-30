package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVideoGameSeries strings.Replacer
var strconvVideoGameSeries strconv.NumError

var basicVideoGameSeriesTraitMapping = map[string]func(*schema.VideoGameSeriesTrait, []string){}
var customVideoGameSeriesTraitMapping = map[string]func(*schema.VideoGameSeriesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/VideoGameSeries", func(ctx jsonld.Context) (interface{}) {
		return NewVideoGameSeriesFromContext(ctx)
	})
























	customVideoGameSeriesTraitMapping["http://schema.org/productionCompany"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/actor"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/director"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/actors"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/directors"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/musicBy"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/trailer"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/containsSeason"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/episode"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/episodes"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/numberOfEpisodes"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/numberOfSeasons"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/season"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/seasons"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
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

	customVideoGameSeriesTraitMapping["http://schema.org/characterAttribute"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.CharacterAttribute = &y

		return
	}

	customVideoGameSeriesTraitMapping["http://schema.org/cheatCode"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.CheatCode = &y

		return
	}

	customVideoGameSeriesTraitMapping["http://schema.org/gameItem"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.GameItem = &y

		return
	}

	customVideoGameSeriesTraitMapping["http://schema.org/gameLocation"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GameLocation, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GameLocation = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GameLocation = s
		}
	}

	customVideoGameSeriesTraitMapping["http://schema.org/gamePlatform"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GamePlatform, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GamePlatform = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GamePlatform = s
		}
	}

	customVideoGameSeriesTraitMapping["http://schema.org/numberOfPlayers"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.NumberOfPlayers = &y

		return
	}

	customVideoGameSeriesTraitMapping["http://schema.org/playMode"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
		var y schema.GamePlayMode
		if strings.HasPrefix(s, "_:") {
			y = NewGamePlayModeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewGamePlayMode()
			y.Id = s
		}

		x.PlayMode = &y

		return
	}

	customVideoGameSeriesTraitMapping["http://schema.org/quest"] = func(x *schema.VideoGameSeriesTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Quest = &y

		return
	}
}

func NewVideoGameSeriesFromContext(ctx jsonld.Context) (x schema.VideoGameSeries) {
	x.Type = "http://schema.org/VideoGameSeries"
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


	MapBasicToVideoGameSeriesTrait(ctx, &x.VideoGameSeriesTrait)
	MapCustomToVideoGameSeriesTrait(ctx, &x.VideoGameSeriesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVideoGameSeriesTrait(ctx jsonld.Context, x *schema.VideoGameSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVideoGameSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVideoGameSeriesTrait(ctx jsonld.Context, x *schema.VideoGameSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVideoGameSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}