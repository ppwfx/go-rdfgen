package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLiveBlogPosting strings.Replacer
var strconvLiveBlogPosting strconv.NumError

var basicLiveBlogPostingTraitMapping = map[string]func(*schema.LiveBlogPostingTrait, []string){}
var customLiveBlogPostingTraitMapping = map[string]func(*schema.LiveBlogPostingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LiveBlogPosting", func(ctx jsonld.Context) (interface{}) {
		return NewLiveBlogPostingFromContext(ctx)
	})





	customLiveBlogPostingTraitMapping["http://schema.org/coverageStartTime"] = func(x *schema.LiveBlogPostingTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.CoverageStartTime = &y

		return
	}

	customLiveBlogPostingTraitMapping["http://schema.org/coverageEndTime"] = func(x *schema.LiveBlogPostingTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.CoverageEndTime = &y

		return
	}

	customLiveBlogPostingTraitMapping["http://schema.org/liveBlogUpdate"] = func(x *schema.LiveBlogPostingTrait, ctx jsonld.Context, s string){
		var y schema.BlogPosting
		if strings.HasPrefix(s, "_:") {
			y = NewBlogPostingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBlogPosting()
			y.Id = s
		}

		x.LiveBlogUpdate = &y

		return
	}
}

func NewLiveBlogPostingFromContext(ctx jsonld.Context) (x schema.LiveBlogPosting) {
	x.Type = "http://schema.org/LiveBlogPosting"
	MapBasicToBlogPostingTrait(ctx, &x.BlogPostingTrait)
	MapCustomToBlogPostingTrait(ctx, &x.BlogPostingTrait)

	MapBasicToSocialMediaPostingTrait(ctx, &x.SocialMediaPostingTrait)
	MapCustomToSocialMediaPostingTrait(ctx, &x.SocialMediaPostingTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLiveBlogPostingTrait(ctx, &x.LiveBlogPostingTrait)
	MapCustomToLiveBlogPostingTrait(ctx, &x.LiveBlogPostingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLiveBlogPostingTrait(ctx jsonld.Context, x *schema.LiveBlogPostingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLiveBlogPostingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLiveBlogPostingTrait(ctx jsonld.Context, x *schema.LiveBlogPostingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLiveBlogPostingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}