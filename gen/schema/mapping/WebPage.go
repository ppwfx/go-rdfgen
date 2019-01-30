package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWebPage strings.Replacer
var strconvWebPage strconv.NumError

var basicWebPageTraitMapping = map[string]func(*schema.WebPageTrait, []string){}
var customWebPageTraitMapping = map[string]func(*schema.WebPageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WebPage", func(ctx jsonld.Context) (interface{}) {
		return NewWebPageFromContext(ctx)
	})







	basicWebPageTraitMapping["http://schema.org/relatedLink"] = func(x *schema.WebPageTrait, s []string) {
		x.RelatedLink = s[0]
	}



	basicWebPageTraitMapping["http://schema.org/significantLink"] = func(x *schema.WebPageTrait, s []string) {
		x.SignificantLink = s[0]
	}


	basicWebPageTraitMapping["http://schema.org/significantLinks"] = func(x *schema.WebPageTrait, s []string) {
		x.SignificantLinks = s[0]
	}



	customWebPageTraitMapping["http://schema.org/speakable"] = func(x *schema.WebPageTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Speakable, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Speakable = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Speakable = s
		}
	}

	customWebPageTraitMapping["http://schema.org/primaryImageOfPage"] = func(x *schema.WebPageTrait, ctx jsonld.Context, s string){
		var y schema.ImageObject
		if strings.HasPrefix(s, "_:") {
			y = NewImageObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewImageObject()
			y.Id = s
		}

		x.PrimaryImageOfPage = &y

		return
	}

	customWebPageTraitMapping["http://schema.org/breadcrumb"] = func(x *schema.WebPageTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Breadcrumb, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Breadcrumb = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Breadcrumb = s
		}
	}

	customWebPageTraitMapping["http://schema.org/lastReviewed"] = func(x *schema.WebPageTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.LastReviewed = &y

		return
	}

	customWebPageTraitMapping["http://schema.org/mainContentOfPage"] = func(x *schema.WebPageTrait, ctx jsonld.Context, s string){
		var y schema.WebPageElement
		if strings.HasPrefix(s, "_:") {
			y = NewWebPageElementFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewWebPageElement()
			y.Id = s
		}

		x.MainContentOfPage = &y

		return
	}

	customWebPageTraitMapping["http://schema.org/reviewedBy"] = func(x *schema.WebPageTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ReviewedBy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ReviewedBy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ReviewedBy = s
		}
	}

	customWebPageTraitMapping["http://schema.org/specialty"] = func(x *schema.WebPageTrait, ctx jsonld.Context, s string){
		var y schema.Specialty
		if strings.HasPrefix(s, "_:") {
			y = NewSpecialtyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewSpecialty()
			y.Id = s
		}

		x.Specialty = &y

		return
	}
}

func NewWebPageFromContext(ctx jsonld.Context) (x schema.WebPage) {
	x.Type = "http://schema.org/WebPage"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWebPageTrait(ctx jsonld.Context, x *schema.WebPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWebPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWebPageTrait(ctx jsonld.Context, x *schema.WebPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWebPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}