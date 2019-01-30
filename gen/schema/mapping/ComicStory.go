package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsComicStory strings.Replacer
var strconvComicStory strconv.NumError

var basicComicStoryTraitMapping = map[string]func(*schema.ComicStoryTrait, []string){}
var customComicStoryTraitMapping = map[string]func(*schema.ComicStoryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ComicStory", func(ctx jsonld.Context) (interface{}) {
		return NewComicStoryFromContext(ctx)
	})







	customComicStoryTraitMapping["http://schema.org/artist"] = func(x *schema.ComicStoryTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Artist = &y

		return
	}

	customComicStoryTraitMapping["http://schema.org/colorist"] = func(x *schema.ComicStoryTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Colorist = &y

		return
	}

	customComicStoryTraitMapping["http://schema.org/inker"] = func(x *schema.ComicStoryTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Inker = &y

		return
	}

	customComicStoryTraitMapping["http://schema.org/letterer"] = func(x *schema.ComicStoryTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Letterer = &y

		return
	}

	customComicStoryTraitMapping["http://schema.org/penciler"] = func(x *schema.ComicStoryTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Penciler = &y

		return
	}
}

func NewComicStoryFromContext(ctx jsonld.Context) (x schema.ComicStory) {
	x.Type = "http://schema.org/ComicStory"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToComicStoryTrait(ctx, &x.ComicStoryTrait)
	MapCustomToComicStoryTrait(ctx, &x.ComicStoryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToComicStoryTrait(ctx jsonld.Context, x *schema.ComicStoryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicComicStoryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToComicStoryTrait(ctx jsonld.Context, x *schema.ComicStoryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customComicStoryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}