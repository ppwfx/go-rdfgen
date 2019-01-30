package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAggregateRating strings.Replacer
var strconvAggregateRating strconv.NumError

var basicAggregateRatingTraitMapping = map[string]func(*schema.AggregateRatingTrait, []string){}
var customAggregateRatingTraitMapping = map[string]func(*schema.AggregateRatingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AggregateRating", func(ctx jsonld.Context) (interface{}) {
		return NewAggregateRatingFromContext(ctx)
	})





	customAggregateRatingTraitMapping["http://schema.org/itemReviewed"] = func(x *schema.AggregateRatingTrait, ctx jsonld.Context, s string){
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

	customAggregateRatingTraitMapping["http://schema.org/ratingCount"] = func(x *schema.AggregateRatingTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.RatingCount = &y

		return
	}

	customAggregateRatingTraitMapping["http://schema.org/reviewCount"] = func(x *schema.AggregateRatingTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.ReviewCount = &y

		return
	}
}

func NewAggregateRatingFromContext(ctx jsonld.Context) (x schema.AggregateRating) {
	x.Type = "http://schema.org/AggregateRating"
	MapBasicToRatingTrait(ctx, &x.RatingTrait)
	MapCustomToRatingTrait(ctx, &x.RatingTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAggregateRatingTrait(ctx, &x.AggregateRatingTrait)
	MapCustomToAggregateRatingTrait(ctx, &x.AggregateRatingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAggregateRatingTrait(ctx jsonld.Context, x *schema.AggregateRatingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAggregateRatingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAggregateRatingTrait(ctx jsonld.Context, x *schema.AggregateRatingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAggregateRatingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}