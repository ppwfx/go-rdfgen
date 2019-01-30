package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCriticReview strings.Replacer
var strconvCriticReview strconv.NumError

var basicCriticReviewTraitMapping = map[string]func(*schema.CriticReviewTrait, []string){}
var customCriticReviewTraitMapping = map[string]func(*schema.CriticReviewTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CriticReview", func(ctx jsonld.Context) (interface{}) {
		return NewCriticReviewFromContext(ctx)
	})

}

func NewCriticReviewFromContext(ctx jsonld.Context) (x schema.CriticReview) {
	x.Type = "http://schema.org/CriticReview"
	MapBasicToReviewTrait(ctx, &x.ReviewTrait)
	MapCustomToReviewTrait(ctx, &x.ReviewTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCriticReviewTrait(ctx, &x.CriticReviewTrait)
	MapCustomToCriticReviewTrait(ctx, &x.CriticReviewTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCriticReviewTrait(ctx jsonld.Context, x *schema.CriticReviewTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCriticReviewTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCriticReviewTrait(ctx jsonld.Context, x *schema.CriticReviewTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCriticReviewTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}