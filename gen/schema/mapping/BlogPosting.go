package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBlogPosting strings.Replacer
var strconvBlogPosting strconv.NumError

var basicBlogPostingTraitMapping = map[string]func(*schema.BlogPostingTrait, []string){}
var customBlogPostingTraitMapping = map[string]func(*schema.BlogPostingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BlogPosting", func(ctx jsonld.Context) (interface{}) {
		return NewBlogPostingFromContext(ctx)
	})

}

func NewBlogPostingFromContext(ctx jsonld.Context) (x schema.BlogPosting) {
	x.Type = "http://schema.org/BlogPosting"
	MapBasicToSocialMediaPostingTrait(ctx, &x.SocialMediaPostingTrait)
	MapCustomToSocialMediaPostingTrait(ctx, &x.SocialMediaPostingTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBlogPostingTrait(ctx, &x.BlogPostingTrait)
	MapCustomToBlogPostingTrait(ctx, &x.BlogPostingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBlogPostingTrait(ctx jsonld.Context, x *schema.BlogPostingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBlogPostingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBlogPostingTrait(ctx jsonld.Context, x *schema.BlogPostingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBlogPostingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}