package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVideoGame strings.Replacer
var strconvVideoGame strconv.NumError

var basicVideoGameTraitMapping = map[string]func(*schema.VideoGameTrait, []string){}
var customVideoGameTraitMapping = map[string]func(*schema.VideoGameTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/VideoGame", func(ctx jsonld.Context) (interface{}) {
		return NewVideoGameFromContext(ctx)
	})













	customVideoGameTraitMapping["http://schema.org/actor"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
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

	customVideoGameTraitMapping["http://schema.org/director"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
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

	customVideoGameTraitMapping["http://schema.org/actors"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
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

	customVideoGameTraitMapping["http://schema.org/directors"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
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

	customVideoGameTraitMapping["http://schema.org/musicBy"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
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

	customVideoGameTraitMapping["http://schema.org/trailer"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
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

	customVideoGameTraitMapping["http://schema.org/cheatCode"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
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

	customVideoGameTraitMapping["http://schema.org/gamePlatform"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
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

	customVideoGameTraitMapping["http://schema.org/playMode"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
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

	customVideoGameTraitMapping["http://schema.org/gameServer"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
		var y schema.GameServer
		if strings.HasPrefix(s, "_:") {
			y = NewGameServerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewGameServer()
			y.Id = s
		}

		x.GameServer = &y

		return
	}

	customVideoGameTraitMapping["http://schema.org/gameTip"] = func(x *schema.VideoGameTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.GameTip = &y

		return
	}
}

func NewVideoGameFromContext(ctx jsonld.Context) (x schema.VideoGame) {
	x.Type = "http://schema.org/VideoGame"
	MapBasicToGameTrait(ctx, &x.GameTrait)
	MapCustomToGameTrait(ctx, &x.GameTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToSoftwareApplicationTrait(ctx, &x.SoftwareApplicationTrait)
	MapCustomToSoftwareApplicationTrait(ctx, &x.SoftwareApplicationTrait)


	MapBasicToVideoGameTrait(ctx, &x.VideoGameTrait)
	MapCustomToVideoGameTrait(ctx, &x.VideoGameTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVideoGameTrait(ctx jsonld.Context, x *schema.VideoGameTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVideoGameTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVideoGameTrait(ctx jsonld.Context, x *schema.VideoGameTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVideoGameTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}