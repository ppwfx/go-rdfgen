package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMediaObject strings.Replacer
var strconvMediaObject strconv.NumError

var basicMediaObjectTraitMapping = map[string]func(*schema.MediaObjectTrait, []string){}
var customMediaObjectTraitMapping = map[string]func(*schema.MediaObjectTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MediaObject", func(ctx jsonld.Context) (interface{}) {
		return NewMediaObjectFromContext(ctx)
	})



	basicMediaObjectTraitMapping["http://schema.org/bitrate"] = func(x *schema.MediaObjectTrait, s []string) {
		x.Bitrate = s[0]
	}


	basicMediaObjectTraitMapping["http://schema.org/contentSize"] = func(x *schema.MediaObjectTrait, s []string) {
		x.ContentSize = s[0]
	}


	basicMediaObjectTraitMapping["http://schema.org/contentUrl"] = func(x *schema.MediaObjectTrait, s []string) {
		x.ContentUrl = s[0]
	}



	basicMediaObjectTraitMapping["http://schema.org/embedUrl"] = func(x *schema.MediaObjectTrait, s []string) {
		x.EmbedUrl = s[0]
	}





	basicMediaObjectTraitMapping["http://schema.org/playerType"] = func(x *schema.MediaObjectTrait, s []string) {
		x.PlayerType = s[0]
	}







	customMediaObjectTraitMapping["http://schema.org/associatedArticle"] = func(x *schema.MediaObjectTrait, ctx jsonld.Context, s string){
		var y schema.NewsArticle
		if strings.HasPrefix(s, "_:") {
			y = NewNewsArticleFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewNewsArticle()
			y.Id = s
		}

		x.AssociatedArticle = &y

		return
	}

	customMediaObjectTraitMapping["http://schema.org/duration"] = func(x *schema.MediaObjectTrait, ctx jsonld.Context, s string){
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

	customMediaObjectTraitMapping["http://schema.org/encodesCreativeWork"] = func(x *schema.MediaObjectTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.EncodesCreativeWork = &y

		return
	}

	customMediaObjectTraitMapping["http://schema.org/encodingFormat"] = func(x *schema.MediaObjectTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EncodingFormat, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EncodingFormat = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EncodingFormat = s
		}
	}

	customMediaObjectTraitMapping["http://schema.org/height"] = func(x *schema.MediaObjectTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Height, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Height = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Height = s
		}
	}

	customMediaObjectTraitMapping["http://schema.org/productionCompany"] = func(x *schema.MediaObjectTrait, ctx jsonld.Context, s string){
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

	customMediaObjectTraitMapping["http://schema.org/regionsAllowed"] = func(x *schema.MediaObjectTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.RegionsAllowed = &y

		return
	}

	customMediaObjectTraitMapping["http://schema.org/requiresSubscription"] = func(x *schema.MediaObjectTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RequiresSubscription, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RequiresSubscription = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RequiresSubscription = s
		}
	}

	customMediaObjectTraitMapping["http://schema.org/uploadDate"] = func(x *schema.MediaObjectTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.UploadDate = &y

		return
	}

	customMediaObjectTraitMapping["http://schema.org/width"] = func(x *schema.MediaObjectTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Width, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Width = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Width = s
		}
	}
}

func NewMediaObjectFromContext(ctx jsonld.Context) (x schema.MediaObject) {
	x.Type = "http://schema.org/MediaObject"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMediaObjectTrait(ctx, &x.MediaObjectTrait)
	MapCustomToMediaObjectTrait(ctx, &x.MediaObjectTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMediaObjectTrait(ctx jsonld.Context, x *schema.MediaObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMediaObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMediaObjectTrait(ctx jsonld.Context, x *schema.MediaObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMediaObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}