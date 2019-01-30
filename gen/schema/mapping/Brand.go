package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBrand strings.Replacer
var strconvBrand strconv.NumError

var basicBrandTraitMapping = map[string]func(*schema.BrandTrait, []string){}
var customBrandTraitMapping = map[string]func(*schema.BrandTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Brand", func(ctx jsonld.Context) (interface{}) {
		return NewBrandFromContext(ctx)
	})





	customBrandTraitMapping["http://schema.org/aggregateRating"] = func(x *schema.BrandTrait, ctx jsonld.Context, s string){
		var y schema.AggregateRating
		if strings.HasPrefix(s, "_:") {
			y = NewAggregateRatingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAggregateRating()
			y.Id = s
		}

		x.AggregateRating = &y

		return
	}

	customBrandTraitMapping["http://schema.org/review"] = func(x *schema.BrandTrait, ctx jsonld.Context, s string){
		var y schema.Review
		if strings.HasPrefix(s, "_:") {
			y = NewReviewFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewReview()
			y.Id = s
		}

		x.Review = &y

		return
	}

	customBrandTraitMapping["http://schema.org/logo"] = func(x *schema.BrandTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Logo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Logo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Logo = s
		}
	}
}

func NewBrandFromContext(ctx jsonld.Context) (x schema.Brand) {
	x.Type = "http://schema.org/Brand"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBrandTrait(ctx, &x.BrandTrait)
	MapCustomToBrandTrait(ctx, &x.BrandTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBrandTrait(ctx jsonld.Context, x *schema.BrandTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBrandTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBrandTrait(ctx jsonld.Context, x *schema.BrandTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBrandTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}