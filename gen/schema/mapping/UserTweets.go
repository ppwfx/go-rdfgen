package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserTweets strings.Replacer
var strconvUserTweets strconv.NumError

var basicUserTweetsTraitMapping = map[string]func(*schema.UserTweetsTrait, []string){}
var customUserTweetsTraitMapping = map[string]func(*schema.UserTweetsTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserTweets", func(ctx jsonld.Context) (interface{}) {
		return NewUserTweetsFromContext(ctx)
	})

}

func NewUserTweetsFromContext(ctx jsonld.Context) (x schema.UserTweets) {
	x.Type = "http://schema.org/UserTweets"
	MapBasicToUserInteractionTrait(ctx, &x.UserInteractionTrait)
	MapCustomToUserInteractionTrait(ctx, &x.UserInteractionTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserTweetsTrait(ctx, &x.UserTweetsTrait)
	MapCustomToUserTweetsTrait(ctx, &x.UserTweetsTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserTweetsTrait(ctx jsonld.Context, x *schema.UserTweetsTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserTweetsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserTweetsTrait(ctx jsonld.Context, x *schema.UserTweetsTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserTweetsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}