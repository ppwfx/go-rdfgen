package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVideoObject strings.Replacer
var strconvVideoObject strconv.NumError

var basicVideoObjectTraitMapping = map[string]func(*schema.VideoObjectTrait, []string){}
var customVideoObjectTraitMapping = map[string]func(*schema.VideoObjectTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/VideoObject", func(ctx jsonld.Context) (interface{}) {
		return NewVideoObjectFromContext(ctx)
	})


	basicVideoObjectTraitMapping["http://schema.org/caption"] = func(x *schema.VideoObjectTrait, s []string) {
		x.Caption = s[0]
	}








	basicVideoObjectTraitMapping["http://schema.org/transcript"] = func(x *schema.VideoObjectTrait, s []string) {
		x.Transcript = s[0]
	}


	basicVideoObjectTraitMapping["http://schema.org/videoFrameSize"] = func(x *schema.VideoObjectTrait, s []string) {
		x.VideoFrameSize = s[0]
	}


	basicVideoObjectTraitMapping["http://schema.org/videoQuality"] = func(x *schema.VideoObjectTrait, s []string) {
		x.VideoQuality = s[0]
	}


	customVideoObjectTraitMapping["http://schema.org/thumbnail"] = func(x *schema.VideoObjectTrait, ctx jsonld.Context, s string){
		var y schema.ImageObject
		if strings.HasPrefix(s, "_:") {
			y = NewImageObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewImageObject()
			y.Id = s
		}

		x.Thumbnail = &y

		return
	}

	customVideoObjectTraitMapping["http://schema.org/actor"] = func(x *schema.VideoObjectTrait, ctx jsonld.Context, s string){
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

	customVideoObjectTraitMapping["http://schema.org/director"] = func(x *schema.VideoObjectTrait, ctx jsonld.Context, s string){
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

	customVideoObjectTraitMapping["http://schema.org/actors"] = func(x *schema.VideoObjectTrait, ctx jsonld.Context, s string){
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

	customVideoObjectTraitMapping["http://schema.org/directors"] = func(x *schema.VideoObjectTrait, ctx jsonld.Context, s string){
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

	customVideoObjectTraitMapping["http://schema.org/musicBy"] = func(x *schema.VideoObjectTrait, ctx jsonld.Context, s string){
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
}

func NewVideoObjectFromContext(ctx jsonld.Context) (x schema.VideoObject) {
	x.Type = "http://schema.org/VideoObject"
	MapBasicToMediaObjectTrait(ctx, &x.MediaObjectTrait)
	MapCustomToMediaObjectTrait(ctx, &x.MediaObjectTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVideoObjectTrait(ctx, &x.VideoObjectTrait)
	MapCustomToVideoObjectTrait(ctx, &x.VideoObjectTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVideoObjectTrait(ctx jsonld.Context, x *schema.VideoObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVideoObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVideoObjectTrait(ctx jsonld.Context, x *schema.VideoObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVideoObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}