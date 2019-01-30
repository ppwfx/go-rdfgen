package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsExchangeRateSpecification strings.Replacer
var strconvExchangeRateSpecification strconv.NumError

var basicExchangeRateSpecificationTraitMapping = map[string]func(*schema.ExchangeRateSpecificationTrait, []string){}
var customExchangeRateSpecificationTraitMapping = map[string]func(*schema.ExchangeRateSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ExchangeRateSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewExchangeRateSpecificationFromContext(ctx)
	})


	basicExchangeRateSpecificationTraitMapping["http://schema.org/currency"] = func(x *schema.ExchangeRateSpecificationTrait, s []string) {
		x.Currency = s[0]
	}




	customExchangeRateSpecificationTraitMapping["http://schema.org/currentExchangeRate"] = func(x *schema.ExchangeRateSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.UnitPriceSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewUnitPriceSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewUnitPriceSpecification()
			y.Id = s
		}

		x.CurrentExchangeRate = &y

		return
	}

	customExchangeRateSpecificationTraitMapping["http://schema.org/exchangeRateSpread"] = func(x *schema.ExchangeRateSpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ExchangeRateSpread, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ExchangeRateSpread = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ExchangeRateSpread = s
		}
	}
}

func NewExchangeRateSpecificationFromContext(ctx jsonld.Context) (x schema.ExchangeRateSpecification) {
	x.Type = "http://schema.org/ExchangeRateSpecification"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToExchangeRateSpecificationTrait(ctx, &x.ExchangeRateSpecificationTrait)
	MapCustomToExchangeRateSpecificationTrait(ctx, &x.ExchangeRateSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToExchangeRateSpecificationTrait(ctx jsonld.Context, x *schema.ExchangeRateSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicExchangeRateSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToExchangeRateSpecificationTrait(ctx jsonld.Context, x *schema.ExchangeRateSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customExchangeRateSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}