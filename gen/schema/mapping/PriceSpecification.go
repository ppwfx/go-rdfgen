package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPriceSpecification strings.Replacer
var strconvPriceSpecification strconv.NumError

var basicPriceSpecificationTraitMapping = map[string]func(*schema.PriceSpecificationTrait, []string){}
var customPriceSpecificationTraitMapping = map[string]func(*schema.PriceSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PriceSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewPriceSpecificationFromContext(ctx)
	})






	basicPriceSpecificationTraitMapping["http://schema.org/maxPrice"] = func(x *schema.PriceSpecificationTrait, s []string) {
		var err error
		x.MaxPrice, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicPriceSpecificationTraitMapping["http://schema.org/minPrice"] = func(x *schema.PriceSpecificationTrait, s []string) {
		var err error
		x.MinPrice, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}



	basicPriceSpecificationTraitMapping["http://schema.org/priceCurrency"] = func(x *schema.PriceSpecificationTrait, s []string) {
		x.PriceCurrency = s[0]
	}


	basicPriceSpecificationTraitMapping["http://schema.org/valueAddedTaxIncluded"] = func(x *schema.PriceSpecificationTrait, s []string) {
		var err error
		x.ValueAddedTaxIncluded, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	customPriceSpecificationTraitMapping["http://schema.org/validFrom"] = func(x *schema.PriceSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ValidFrom = &y

		return
	}

	customPriceSpecificationTraitMapping["http://schema.org/validThrough"] = func(x *schema.PriceSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ValidThrough = &y

		return
	}

	customPriceSpecificationTraitMapping["http://schema.org/eligibleQuantity"] = func(x *schema.PriceSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.EligibleQuantity = &y

		return
	}

	customPriceSpecificationTraitMapping["http://schema.org/eligibleTransactionVolume"] = func(x *schema.PriceSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.PriceSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewPriceSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPriceSpecification()
			y.Id = s
		}

		x.EligibleTransactionVolume = &y

		return
	}

	customPriceSpecificationTraitMapping["http://schema.org/price"] = func(x *schema.PriceSpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Price, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Price = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Price = s
		}
	}
}

func NewPriceSpecificationFromContext(ctx jsonld.Context) (x schema.PriceSpecification) {
	x.Type = "http://schema.org/PriceSpecification"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPriceSpecificationTrait(ctx, &x.PriceSpecificationTrait)
	MapCustomToPriceSpecificationTrait(ctx, &x.PriceSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPriceSpecificationTrait(ctx jsonld.Context, x *schema.PriceSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPriceSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPriceSpecificationTrait(ctx jsonld.Context, x *schema.PriceSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPriceSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}