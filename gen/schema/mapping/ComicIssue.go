package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsComicIssue strings.Replacer
var strconvComicIssue strconv.NumError

var basicComicIssueTraitMapping = map[string]func(*schema.ComicIssueTrait, []string){}
var customComicIssueTraitMapping = map[string]func(*schema.ComicIssueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ComicIssue", func(ctx jsonld.Context) (interface{}) {
		return NewComicIssueFromContext(ctx)
	})







	basicComicIssueTraitMapping["http://schema.org/variantCover"] = func(x *schema.ComicIssueTrait, s []string) {
		x.VariantCover = s[0]
	}


	customComicIssueTraitMapping["http://schema.org/artist"] = func(x *schema.ComicIssueTrait, ctx jsonld.Context, s string){
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

	customComicIssueTraitMapping["http://schema.org/colorist"] = func(x *schema.ComicIssueTrait, ctx jsonld.Context, s string){
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

	customComicIssueTraitMapping["http://schema.org/inker"] = func(x *schema.ComicIssueTrait, ctx jsonld.Context, s string){
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

	customComicIssueTraitMapping["http://schema.org/letterer"] = func(x *schema.ComicIssueTrait, ctx jsonld.Context, s string){
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

	customComicIssueTraitMapping["http://schema.org/penciler"] = func(x *schema.ComicIssueTrait, ctx jsonld.Context, s string){
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

func NewComicIssueFromContext(ctx jsonld.Context) (x schema.ComicIssue) {
	x.Type = "http://schema.org/ComicIssue"
	MapBasicToPublicationIssueTrait(ctx, &x.PublicationIssueTrait)
	MapCustomToPublicationIssueTrait(ctx, &x.PublicationIssueTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToComicIssueTrait(ctx, &x.ComicIssueTrait)
	MapCustomToComicIssueTrait(ctx, &x.ComicIssueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToComicIssueTrait(ctx jsonld.Context, x *schema.ComicIssueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicComicIssueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToComicIssueTrait(ctx jsonld.Context, x *schema.ComicIssueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customComicIssueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}