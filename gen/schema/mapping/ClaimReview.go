package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsClaimReview strings.Replacer
var strconvClaimReview strconv.NumError

var basicClaimReviewTraitMapping = map[string]func(*schema.ClaimReviewTrait, []string){}
var customClaimReviewTraitMapping = map[string]func(*schema.ClaimReviewTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ClaimReview", func(ctx jsonld.Context) (interface{}) {
		return NewClaimReviewFromContext(ctx)
	})


	basicClaimReviewTraitMapping["http://schema.org/claimReviewed"] = func(x *schema.ClaimReviewTrait, s []string) {
		x.ClaimReviewed = s[0]
	}

}

func NewClaimReviewFromContext(ctx jsonld.Context) (x schema.ClaimReview) {
	x.Type = "http://schema.org/ClaimReview"
	MapBasicToReviewTrait(ctx, &x.ReviewTrait)
	MapCustomToReviewTrait(ctx, &x.ReviewTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToClaimReviewTrait(ctx, &x.ClaimReviewTrait)
	MapCustomToClaimReviewTrait(ctx, &x.ClaimReviewTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToClaimReviewTrait(ctx jsonld.Context, x *schema.ClaimReviewTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicClaimReviewTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToClaimReviewTrait(ctx jsonld.Context, x *schema.ClaimReviewTrait) () {
	for k, v := range ctx.Current {
		f, ok := customClaimReviewTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}