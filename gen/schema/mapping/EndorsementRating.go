package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEndorsementRating strings.Replacer
var strconvEndorsementRating strconv.NumError

var basicEndorsementRatingTraitMapping = map[string]func(*schema.EndorsementRatingTrait, []string){}
var customEndorsementRatingTraitMapping = map[string]func(*schema.EndorsementRatingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EndorsementRating", func(ctx jsonld.Context) (interface{}) {
		return NewEndorsementRatingFromContext(ctx)
	})

}

func NewEndorsementRatingFromContext(ctx jsonld.Context) (x schema.EndorsementRating) {
	x.Type = "http://schema.org/EndorsementRating"
	MapBasicToRatingTrait(ctx, &x.RatingTrait)
	MapCustomToRatingTrait(ctx, &x.RatingTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEndorsementRatingTrait(ctx, &x.EndorsementRatingTrait)
	MapCustomToEndorsementRatingTrait(ctx, &x.EndorsementRatingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEndorsementRatingTrait(ctx jsonld.Context, x *schema.EndorsementRatingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEndorsementRatingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEndorsementRatingTrait(ctx jsonld.Context, x *schema.EndorsementRatingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEndorsementRatingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}