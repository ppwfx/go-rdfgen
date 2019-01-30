package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicComposition strings.Replacer
var strconvMusicComposition strconv.NumError

var basicMusicCompositionTraitMapping = map[string]func(*schema.MusicCompositionTrait, []string){}
var customMusicCompositionTraitMapping = map[string]func(*schema.MusicCompositionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicComposition", func(ctx jsonld.Context) (interface{}) {
		return NewMusicCompositionFromContext(ctx)
	})





	basicMusicCompositionTraitMapping["http://schema.org/iswcCode"] = func(x *schema.MusicCompositionTrait, s []string) {
		x.IswcCode = s[0]
	}





	basicMusicCompositionTraitMapping["http://schema.org/musicCompositionForm"] = func(x *schema.MusicCompositionTrait, s []string) {
		x.MusicCompositionForm = s[0]
	}


	basicMusicCompositionTraitMapping["http://schema.org/musicalKey"] = func(x *schema.MusicCompositionTrait, s []string) {
		x.MusicalKey = s[0]
	}



	customMusicCompositionTraitMapping["http://schema.org/composer"] = func(x *schema.MusicCompositionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Composer, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Composer = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Composer = s
		}
	}

	customMusicCompositionTraitMapping["http://schema.org/firstPerformance"] = func(x *schema.MusicCompositionTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.FirstPerformance = &y

		return
	}

	customMusicCompositionTraitMapping["http://schema.org/includedComposition"] = func(x *schema.MusicCompositionTrait, ctx jsonld.Context, s string){
		var y schema.MusicComposition
		if strings.HasPrefix(s, "_:") {
			y = NewMusicCompositionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicComposition()
			y.Id = s
		}

		x.IncludedComposition = &y

		return
	}

	customMusicCompositionTraitMapping["http://schema.org/lyricist"] = func(x *schema.MusicCompositionTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Lyricist = &y

		return
	}

	customMusicCompositionTraitMapping["http://schema.org/lyrics"] = func(x *schema.MusicCompositionTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.Lyrics = &y

		return
	}

	customMusicCompositionTraitMapping["http://schema.org/musicArrangement"] = func(x *schema.MusicCompositionTrait, ctx jsonld.Context, s string){
		var y schema.MusicComposition
		if strings.HasPrefix(s, "_:") {
			y = NewMusicCompositionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicComposition()
			y.Id = s
		}

		x.MusicArrangement = &y

		return
	}

	customMusicCompositionTraitMapping["http://schema.org/recordedAs"] = func(x *schema.MusicCompositionTrait, ctx jsonld.Context, s string){
		var y schema.MusicRecording
		if strings.HasPrefix(s, "_:") {
			y = NewMusicRecordingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicRecording()
			y.Id = s
		}

		x.RecordedAs = &y

		return
	}
}

func NewMusicCompositionFromContext(ctx jsonld.Context) (x schema.MusicComposition) {
	x.Type = "http://schema.org/MusicComposition"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicCompositionTrait(ctx, &x.MusicCompositionTrait)
	MapCustomToMusicCompositionTrait(ctx, &x.MusicCompositionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicCompositionTrait(ctx jsonld.Context, x *schema.MusicCompositionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicCompositionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicCompositionTrait(ctx jsonld.Context, x *schema.MusicCompositionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicCompositionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}