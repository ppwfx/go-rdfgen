package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEmployerReview strings.Replacer
var strconvEmployerReview strconv.NumError

var basicEmployerReviewTraitMapping = map[string]func(*schema.EmployerReviewTrait, []string){}
var customEmployerReviewTraitMapping = map[string]func(*schema.EmployerReviewTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EmployerReview", func(ctx jsonld.Context) (interface{}) {
		return NewEmployerReviewFromContext(ctx)
	})

}

func NewEmployerReviewFromContext(ctx jsonld.Context) (x schema.EmployerReview) {
	x.Type = "http://schema.org/EmployerReview"
	MapBasicToReviewTrait(ctx, &x.ReviewTrait)
	MapCustomToReviewTrait(ctx, &x.ReviewTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEmployerReviewTrait(ctx, &x.EmployerReviewTrait)
	MapCustomToEmployerReviewTrait(ctx, &x.EmployerReviewTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEmployerReviewTrait(ctx jsonld.Context, x *schema.EmployerReviewTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEmployerReviewTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEmployerReviewTrait(ctx jsonld.Context, x *schema.EmployerReviewTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEmployerReviewTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}