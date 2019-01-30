package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsClip strings.Replacer
var strconvClip strconv.NumError

var basicClipTraitMapping = map[string]func(*schema.ClipTrait, []string){}
var customClipTraitMapping = map[string]func(*schema.ClipTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Clip", func(ctx jsonld.Context) (interface{}) {
		return NewClipFromContext(ctx)
	})











	customClipTraitMapping["http://schema.org/actor"] = func(x *schema.ClipTrait, ctx jsonld.Context, s string){
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

	customClipTraitMapping["http://schema.org/director"] = func(x *schema.ClipTrait, ctx jsonld.Context, s string){
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

	customClipTraitMapping["http://schema.org/actors"] = func(x *schema.ClipTrait, ctx jsonld.Context, s string){
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

	customClipTraitMapping["http://schema.org/directors"] = func(x *schema.ClipTrait, ctx jsonld.Context, s string){
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

	customClipTraitMapping["http://schema.org/musicBy"] = func(x *schema.ClipTrait, ctx jsonld.Context, s string){
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

	customClipTraitMapping["http://schema.org/partOfSeries"] = func(x *schema.ClipTrait, ctx jsonld.Context, s string){
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

	customClipTraitMapping["http://schema.org/partOfSeason"] = func(x *schema.ClipTrait, ctx jsonld.Context, s string){
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

	customClipTraitMapping["http://schema.org/partOfEpisode"] = func(x *schema.ClipTrait, ctx jsonld.Context, s string){
		var y schema.Episode
		if strings.HasPrefix(s, "_:") {
			y = NewEpisodeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEpisode()
			y.Id = s
		}

		x.PartOfEpisode = &y

		return
	}

	customClipTraitMapping["http://schema.org/clipNumber"] = func(x *schema.ClipTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ClipNumber, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ClipNumber = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ClipNumber = s
		}
	}
}

func NewClipFromContext(ctx jsonld.Context) (x schema.Clip) {
	x.Type = "http://schema.org/Clip"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToClipTrait(ctx, &x.ClipTrait)
	MapCustomToClipTrait(ctx, &x.ClipTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToClipTrait(ctx jsonld.Context, x *schema.ClipTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicClipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToClipTrait(ctx jsonld.Context, x *schema.ClipTrait) () {
	for k, v := range ctx.Current {
		f, ok := customClipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}