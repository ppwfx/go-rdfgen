package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDiscussionForumPosting strings.Replacer
var strconvDiscussionForumPosting strconv.NumError

var basicDiscussionForumPostingTraitMapping = map[string]func(*schema.DiscussionForumPostingTrait, []string){}
var customDiscussionForumPostingTraitMapping = map[string]func(*schema.DiscussionForumPostingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DiscussionForumPosting", func(ctx jsonld.Context) (interface{}) {
		return NewDiscussionForumPostingFromContext(ctx)
	})

}

func NewDiscussionForumPostingFromContext(ctx jsonld.Context) (x schema.DiscussionForumPosting) {
	x.Type = "http://schema.org/DiscussionForumPosting"
	MapBasicToSocialMediaPostingTrait(ctx, &x.SocialMediaPostingTrait)
	MapCustomToSocialMediaPostingTrait(ctx, &x.SocialMediaPostingTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDiscussionForumPostingTrait(ctx, &x.DiscussionForumPostingTrait)
	MapCustomToDiscussionForumPostingTrait(ctx, &x.DiscussionForumPostingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDiscussionForumPostingTrait(ctx jsonld.Context, x *schema.DiscussionForumPostingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDiscussionForumPostingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDiscussionForumPostingTrait(ctx jsonld.Context, x *schema.DiscussionForumPostingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDiscussionForumPostingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}