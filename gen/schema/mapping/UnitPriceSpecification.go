package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUnitPriceSpecification strings.Replacer
var strconvUnitPriceSpecification strconv.NumError

var basicUnitPriceSpecificationTraitMapping = map[string]func(*schema.UnitPriceSpecificationTrait, []string){}
var customUnitPriceSpecificationTraitMapping = map[string]func(*schema.UnitPriceSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UnitPriceSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewUnitPriceSpecificationFromContext(ctx)
	})



	basicUnitPriceSpecificationTraitMapping["http://schema.org/unitText"] = func(x *schema.UnitPriceSpecificationTrait, s []string) {
		x.UnitText = s[0]
	}


	basicUnitPriceSpecificationTraitMapping["http://schema.org/priceType"] = func(x *schema.UnitPriceSpecificationTrait, s []string) {
		x.PriceType = s[0]
	}


	basicUnitPriceSpecificationTraitMapping["http://schema.org/billingIncrement"] = func(x *schema.UnitPriceSpecificationTrait, s []string) {
		var err error
		x.BillingIncrement, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}



	customUnitPriceSpecificationTraitMapping["http://schema.org/unitCode"] = func(x *schema.UnitPriceSpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.UnitCode, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.UnitCode = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.UnitCode = s
		}
	}

	customUnitPriceSpecificationTraitMapping["http://schema.org/referenceQuantity"] = func(x *schema.UnitPriceSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.ReferenceQuantity = &y

		return
	}
}

func NewUnitPriceSpecificationFromContext(ctx jsonld.Context) (x schema.UnitPriceSpecification) {
	x.Type = "http://schema.org/UnitPriceSpecification"
	MapBasicToPriceSpecificationTrait(ctx, &x.PriceSpecificationTrait)
	MapCustomToPriceSpecificationTrait(ctx, &x.PriceSpecificationTrait)

	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUnitPriceSpecificationTrait(ctx, &x.UnitPriceSpecificationTrait)
	MapCustomToUnitPriceSpecificationTrait(ctx, &x.UnitPriceSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUnitPriceSpecificationTrait(ctx jsonld.Context, x *schema.UnitPriceSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUnitPriceSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUnitPriceSpecificationTrait(ctx jsonld.Context, x *schema.UnitPriceSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUnitPriceSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}