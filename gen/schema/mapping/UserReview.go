package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserReview strings.Replacer
var strconvUserReview strconv.NumError

var basicUserReviewTraitMapping = map[string]func(*schema.UserReviewTrait, []string){}
var customUserReviewTraitMapping = map[string]func(*schema.UserReviewTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserReview", func(ctx jsonld.Context) (interface{}) {
		return NewUserReviewFromContext(ctx)
	})

}

func NewUserReviewFromContext(ctx jsonld.Context) (x schema.UserReview) {
	x.Type = "http://schema.org/UserReview"
	MapBasicToReviewTrait(ctx, &x.ReviewTrait)
	MapCustomToReviewTrait(ctx, &x.ReviewTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserReviewTrait(ctx, &x.UserReviewTrait)
	MapCustomToUserReviewTrait(ctx, &x.UserReviewTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserReviewTrait(ctx jsonld.Context, x *schema.UserReviewTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserReviewTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserReviewTrait(ctx jsonld.Context, x *schema.UserReviewTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserReviewTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}