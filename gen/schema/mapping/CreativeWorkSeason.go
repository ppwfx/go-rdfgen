package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCreativeWorkSeason strings.Replacer
var strconvCreativeWorkSeason strconv.NumError

var basicCreativeWorkSeasonTraitMapping = map[string]func(*schema.CreativeWorkSeasonTrait, []string){}
var customCreativeWorkSeasonTraitMapping = map[string]func(*schema.CreativeWorkSeasonTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CreativeWorkSeason", func(ctx jsonld.Context) (interface{}) {
		return NewCreativeWorkSeasonFromContext(ctx)
	})













	customCreativeWorkSeasonTraitMapping["http://schema.org/productionCompany"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkSeasonTraitMapping["http://schema.org/actor"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkSeasonTraitMapping["http://schema.org/director"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkSeasonTraitMapping["http://schema.org/endDate"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EndDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EndDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EndDate = s
		}
	}

	customCreativeWorkSeasonTraitMapping["http://schema.org/startDate"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.StartDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.StartDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.StartDate = s
		}
	}

	customCreativeWorkSeasonTraitMapping["http://schema.org/trailer"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkSeasonTraitMapping["http://schema.org/episode"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkSeasonTraitMapping["http://schema.org/episodes"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkSeasonTraitMapping["http://schema.org/numberOfEpisodes"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkSeasonTraitMapping["http://schema.org/partOfSeries"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWorkSeries
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkSeriesFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWorkSeries()
			y.Id = s
		}

		x.PartOfSeries = &y

		return
	}

	customCreativeWorkSeasonTraitMapping["http://schema.org/seasonNumber"] = func(x *schema.CreativeWorkSeasonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SeasonNumber, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SeasonNumber = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SeasonNumber = s
		}
	}
}

func NewCreativeWorkSeasonFromContext(ctx jsonld.Context) (x schema.CreativeWorkSeason) {
	x.Type = "http://schema.org/CreativeWorkSeason"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCreativeWorkSeasonTrait(ctx, &x.CreativeWorkSeasonTrait)
	MapCustomToCreativeWorkSeasonTrait(ctx, &x.CreativeWorkSeasonTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCreativeWorkSeasonTrait(ctx jsonld.Context, x *schema.CreativeWorkSeasonTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCreativeWorkSeasonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCreativeWorkSeasonTrait(ctx jsonld.Context, x *schema.CreativeWorkSeasonTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCreativeWorkSeasonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}