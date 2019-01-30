package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReview strings.Replacer
var strconvReview strconv.NumError

var basicReviewTraitMapping = map[string]func(*schema.ReviewTrait, []string){}
var customReviewTraitMapping = map[string]func(*schema.ReviewTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Review", func(ctx jsonld.Context) (interface{}) {
		return NewReviewFromContext(ctx)
	})


	basicReviewTraitMapping["http://schema.org/reviewAspect"] = func(x *schema.ReviewTrait, s []string) {
		x.ReviewAspect = s[0]
	}



	basicReviewTraitMapping["http://schema.org/reviewBody"] = func(x *schema.ReviewTrait, s []string) {
		x.ReviewBody = s[0]
	}



	customReviewTraitMapping["http://schema.org/itemReviewed"] = func(x *schema.ReviewTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.ItemReviewed = &y

		return
	}

	customReviewTraitMapping["http://schema.org/reviewRating"] = func(x *schema.ReviewTrait, ctx jsonld.Context, s string){
		var y schema.Rating
		if strings.HasPrefix(s, "_:") {
			y = NewRatingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewRating()
			y.Id = s
		}

		x.ReviewRating = &y

		return
	}
}

func NewReviewFromContext(ctx jsonld.Context) (x schema.Review) {
	x.Type = "http://schema.org/Review"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReviewTrait(ctx, &x.ReviewTrait)
	MapCustomToReviewTrait(ctx, &x.ReviewTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReviewTrait(ctx jsonld.Context, x *schema.ReviewTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReviewTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReviewTrait(ctx jsonld.Context, x *schema.ReviewTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReviewTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}