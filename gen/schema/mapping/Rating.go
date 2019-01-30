package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRating strings.Replacer
var strconvRating strconv.NumError

var basicRatingTraitMapping = map[string]func(*schema.RatingTrait, []string){}
var customRatingTraitMapping = map[string]func(*schema.RatingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Rating", func(ctx jsonld.Context) (interface{}) {
		return NewRatingFromContext(ctx)
	})





	basicRatingTraitMapping["http://schema.org/reviewAspect"] = func(x *schema.RatingTrait, s []string) {
		x.ReviewAspect = s[0]
	}



	customRatingTraitMapping["http://schema.org/author"] = func(x *schema.RatingTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Author, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Author = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Author = s
		}
	}

	customRatingTraitMapping["http://schema.org/bestRating"] = func(x *schema.RatingTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BestRating, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BestRating = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BestRating = s
		}
	}

	customRatingTraitMapping["http://schema.org/ratingValue"] = func(x *schema.RatingTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RatingValue, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RatingValue = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RatingValue = s
		}
	}

	customRatingTraitMapping["http://schema.org/worstRating"] = func(x *schema.RatingTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.WorstRating, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.WorstRating = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.WorstRating = s
		}
	}
}

func NewRatingFromContext(ctx jsonld.Context) (x schema.Rating) {
	x.Type = "http://schema.org/Rating"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRatingTrait(ctx, &x.RatingTrait)
	MapCustomToRatingTrait(ctx, &x.RatingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRatingTrait(ctx jsonld.Context, x *schema.RatingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRatingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRatingTrait(ctx jsonld.Context, x *schema.RatingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRatingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}