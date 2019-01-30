package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEmployerAggregateRating strings.Replacer
var strconvEmployerAggregateRating strconv.NumError

var basicEmployerAggregateRatingTraitMapping = map[string]func(*schema.EmployerAggregateRatingTrait, []string){}
var customEmployerAggregateRatingTraitMapping = map[string]func(*schema.EmployerAggregateRatingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EmployerAggregateRating", func(ctx jsonld.Context) (interface{}) {
		return NewEmployerAggregateRatingFromContext(ctx)
	})

}

func NewEmployerAggregateRatingFromContext(ctx jsonld.Context) (x schema.EmployerAggregateRating) {
	x.Type = "http://schema.org/EmployerAggregateRating"
	MapBasicToAggregateRatingTrait(ctx, &x.AggregateRatingTrait)
	MapCustomToAggregateRatingTrait(ctx, &x.AggregateRatingTrait)

	MapBasicToRatingTrait(ctx, &x.RatingTrait)
	MapCustomToRatingTrait(ctx, &x.RatingTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEmployerAggregateRatingTrait(ctx, &x.EmployerAggregateRatingTrait)
	MapCustomToEmployerAggregateRatingTrait(ctx, &x.EmployerAggregateRatingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEmployerAggregateRatingTrait(ctx jsonld.Context, x *schema.EmployerAggregateRatingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEmployerAggregateRatingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEmployerAggregateRatingTrait(ctx jsonld.Context, x *schema.EmployerAggregateRatingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEmployerAggregateRatingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}