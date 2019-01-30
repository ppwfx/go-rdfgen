package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAudiobook strings.Replacer
var strconvAudiobook strconv.NumError

var basicAudiobookTraitMapping = map[string]func(*schema.AudiobookTrait, []string){}
var customAudiobookTraitMapping = map[string]func(*schema.AudiobookTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Audiobook", func(ctx jsonld.Context) (interface{}) {
		return NewAudiobookFromContext(ctx)
	})




	customAudiobookTraitMapping["http://schema.org/duration"] = func(x *schema.AudiobookTrait, ctx jsonld.Context, s string){
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

	customAudiobookTraitMapping["http://schema.org/readBy"] = func(x *schema.AudiobookTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.ReadBy = &y

		return
	}
}

func NewAudiobookFromContext(ctx jsonld.Context) (x schema.Audiobook) {
	x.Type = "http://schema.org/Audiobook"
	MapBasicToAudioObjectTrait(ctx, &x.AudioObjectTrait)
	MapCustomToAudioObjectTrait(ctx, &x.AudioObjectTrait)

	MapBasicToMediaObjectTrait(ctx, &x.MediaObjectTrait)
	MapCustomToMediaObjectTrait(ctx, &x.MediaObjectTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToBookTrait(ctx, &x.BookTrait)
	MapCustomToBookTrait(ctx, &x.BookTrait)


	MapBasicToAudiobookTrait(ctx, &x.AudiobookTrait)
	MapCustomToAudiobookTrait(ctx, &x.AudiobookTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAudiobookTrait(ctx jsonld.Context, x *schema.AudiobookTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAudiobookTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAudiobookTrait(ctx jsonld.Context, x *schema.AudiobookTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAudiobookTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}