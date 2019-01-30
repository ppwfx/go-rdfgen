package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSocialMediaPosting strings.Replacer
var strconvSocialMediaPosting strconv.NumError

var basicSocialMediaPostingTraitMapping = map[string]func(*schema.SocialMediaPostingTrait, []string){}
var customSocialMediaPostingTraitMapping = map[string]func(*schema.SocialMediaPostingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SocialMediaPosting", func(ctx jsonld.Context) (interface{}) {
		return NewSocialMediaPostingFromContext(ctx)
	})



	customSocialMediaPostingTraitMapping["http://schema.org/sharedContent"] = func(x *schema.SocialMediaPostingTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.SharedContent = &y

		return
	}
}

func NewSocialMediaPostingFromContext(ctx jsonld.Context) (x schema.SocialMediaPosting) {
	x.Type = "http://schema.org/SocialMediaPosting"
	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSocialMediaPostingTrait(ctx, &x.SocialMediaPostingTrait)
	MapCustomToSocialMediaPostingTrait(ctx, &x.SocialMediaPostingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSocialMediaPostingTrait(ctx jsonld.Context, x *schema.SocialMediaPostingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSocialMediaPostingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSocialMediaPostingTrait(ctx jsonld.Context, x *schema.SocialMediaPostingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSocialMediaPostingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}