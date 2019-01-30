package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAggregateOffer strings.Replacer
var strconvAggregateOffer strconv.NumError

var basicAggregateOfferTraitMapping = map[string]func(*schema.AggregateOfferTrait, []string){}
var customAggregateOfferTraitMapping = map[string]func(*schema.AggregateOfferTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AggregateOffer", func(ctx jsonld.Context) (interface{}) {
		return NewAggregateOfferFromContext(ctx)
	})






	customAggregateOfferTraitMapping["http://schema.org/offers"] = func(x *schema.AggregateOfferTrait, ctx jsonld.Context, s string){
		var y schema.Offer
		if strings.HasPrefix(s, "_:") {
			y = NewOfferFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOffer()
			y.Id = s
		}

		x.Offers = &y

		return
	}

	customAggregateOfferTraitMapping["http://schema.org/highPrice"] = func(x *schema.AggregateOfferTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.HighPrice, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.HighPrice = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.HighPrice = s
		}
	}

	customAggregateOfferTraitMapping["http://schema.org/lowPrice"] = func(x *schema.AggregateOfferTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.LowPrice, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.LowPrice = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.LowPrice = s
		}
	}

	customAggregateOfferTraitMapping["http://schema.org/offerCount"] = func(x *schema.AggregateOfferTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.OfferCount = &y

		return
	}
}

func NewAggregateOfferFromContext(ctx jsonld.Context) (x schema.AggregateOffer) {
	x.Type = "http://schema.org/AggregateOffer"
	MapBasicToOfferTrait(ctx, &x.OfferTrait)
	MapCustomToOfferTrait(ctx, &x.OfferTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAggregateOfferTrait(ctx, &x.AggregateOfferTrait)
	MapCustomToAggregateOfferTrait(ctx, &x.AggregateOfferTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAggregateOfferTrait(ctx jsonld.Context, x *schema.AggregateOfferTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAggregateOfferTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAggregateOfferTrait(ctx jsonld.Context, x *schema.AggregateOfferTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAggregateOfferTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}