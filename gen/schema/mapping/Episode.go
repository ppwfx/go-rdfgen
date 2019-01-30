package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEpisode strings.Replacer
var strconvEpisode strconv.NumError

var basicEpisodeTraitMapping = map[string]func(*schema.EpisodeTrait, []string){}
var customEpisodeTraitMapping = map[string]func(*schema.EpisodeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Episode", func(ctx jsonld.Context) (interface{}) {
		return NewEpisodeFromContext(ctx)
	})












	customEpisodeTraitMapping["http://schema.org/productionCompany"] = func(x *schema.EpisodeTrait, ctx jsonld.Context, s string){
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

	customEpisodeTraitMapping["http://schema.org/actor"] = func(x *schema.EpisodeTrait, ctx jsonld.Context, s string){
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

	customEpisodeTraitMapping["http://schema.org/director"] = func(x *schema.EpisodeTrait, ctx jsonld.Context, s string){
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

	customEpisodeTraitMapping["http://schema.org/actors"] = func(x *schema.EpisodeTrait, ctx jsonld.Context, s string){
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

	customEpisodeTraitMapping["http://schema.org/directors"] = func(x *schema.EpisodeTrait, ctx jsonld.Context, s string){
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

	customEpisodeTraitMapping["http://schema.org/musicBy"] = func(x *schema.EpisodeTrait, ctx jsonld.Context, s string){
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

	customEpisodeTraitMapping["http://schema.org/trailer"] = func(x *schema.EpisodeTrait, ctx jsonld.Context, s string){
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

	customEpisodeTraitMapping["http://schema.org/partOfSeries"] = func(x *schema.EpisodeTrait, ctx jsonld.Context, s string){
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

	customEpisodeTraitMapping["http://schema.org/episodeNumber"] = func(x *schema.EpisodeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EpisodeNumber, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EpisodeNumber = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EpisodeNumber = s
		}
	}

	customEpisodeTraitMapping["http://schema.org/partOfSeason"] = func(x *schema.EpisodeTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWorkSeason
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkSeasonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWorkSeason()
			y.Id = s
		}

		x.PartOfSeason = &y

		return
	}
}

func NewEpisodeFromContext(ctx jsonld.Context) (x schema.Episode) {
	x.Type = "http://schema.org/Episode"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEpisodeTrait(ctx, &x.EpisodeTrait)
	MapCustomToEpisodeTrait(ctx, &x.EpisodeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEpisodeTrait(ctx jsonld.Context, x *schema.EpisodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEpisodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEpisodeTrait(ctx jsonld.Context, x *schema.EpisodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEpisodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}