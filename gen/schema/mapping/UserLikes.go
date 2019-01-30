package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserLikes strings.Replacer
var strconvUserLikes strconv.NumError

var basicUserLikesTraitMapping = map[string]func(*schema.UserLikesTrait, []string){}
var customUserLikesTraitMapping = map[string]func(*schema.UserLikesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserLikes", func(ctx jsonld.Context) (interface{}) {
		return NewUserLikesFromContext(ctx)
	})

}

func NewUserLikesFromContext(ctx jsonld.Context) (x schema.UserLikes) {
	x.Type = "http://schema.org/UserLikes"
	MapBasicToUserInteractionTrait(ctx, &x.UserInteractionTrait)
	MapCustomToUserInteractionTrait(ctx, &x.UserInteractionTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserLikesTrait(ctx, &x.UserLikesTrait)
	MapCustomToUserLikesTrait(ctx, &x.UserLikesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserLikesTrait(ctx jsonld.Context, x *schema.UserLikesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserLikesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserLikesTrait(ctx jsonld.Context, x *schema.UserLikesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserLikesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}