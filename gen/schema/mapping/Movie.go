package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMovie strings.Replacer
var strconvMovie strconv.NumError

var basicMovieTraitMapping = map[string]func(*schema.MovieTrait, []string){}
var customMovieTraitMapping = map[string]func(*schema.MovieTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Movie", func(ctx jsonld.Context) (interface{}) {
		return NewMovieFromContext(ctx)
	})












	customMovieTraitMapping["http://schema.org/duration"] = func(x *schema.MovieTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.Duration = &y

		return
	}

	customMovieTraitMapping["http://schema.org/productionCompany"] = func(x *schema.MovieTrait, ctx jsonld.Context, s string){
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

	customMovieTraitMapping["http://schema.org/actor"] = func(x *schema.MovieTrait, ctx jsonld.Context, s string){
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

	customMovieTraitMapping["http://schema.org/director"] = func(x *schema.MovieTrait, ctx jsonld.Context, s string){
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

	customMovieTraitMapping["http://schema.org/actors"] = func(x *schema.MovieTrait, ctx jsonld.Context, s string){
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

	customMovieTraitMapping["http://schema.org/directors"] = func(x *schema.MovieTrait, ctx jsonld.Context, s string){
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

	customMovieTraitMapping["http://schema.org/musicBy"] = func(x *schema.MovieTrait, ctx jsonld.Context, s string){
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

	customMovieTraitMapping["http://schema.org/trailer"] = func(x *schema.MovieTrait, ctx jsonld.Context, s string){
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

	customMovieTraitMapping["http://schema.org/countryOfOrigin"] = func(x *schema.MovieTrait, ctx jsonld.Context, s string){
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

	customMovieTraitMapping["http://schema.org/subtitleLanguage"] = func(x *schema.MovieTrait, ctx jsonld.Context, s string){
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
}

func NewMovieFromContext(ctx jsonld.Context) (x schema.Movie) {
	x.Type = "http://schema.org/Movie"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMovieTrait(ctx, &x.MovieTrait)
	MapCustomToMovieTrait(ctx, &x.MovieTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMovieTrait(ctx jsonld.Context, x *schema.MovieTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMovieTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMovieTrait(ctx jsonld.Context, x *schema.MovieTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMovieTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}