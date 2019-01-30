package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFinancialProduct strings.Replacer
var strconvFinancialProduct strconv.NumError

var basicFinancialProductTraitMapping = map[string]func(*schema.FinancialProductTrait, []string){}
var customFinancialProductTraitMapping = map[string]func(*schema.FinancialProductTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FinancialProduct", func(ctx jsonld.Context) (interface{}) {
		return NewFinancialProductFromContext(ctx)
	})





	customFinancialProductTraitMapping["http://schema.org/feesAndCommissionsSpecification"] = func(x *schema.FinancialProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.FeesAndCommissionsSpecification, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.FeesAndCommissionsSpecification = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.FeesAndCommissionsSpecification = s
		}
	}

	customFinancialProductTraitMapping["http://schema.org/annualPercentageRate"] = func(x *schema.FinancialProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AnnualPercentageRate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AnnualPercentageRate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AnnualPercentageRate = s
		}
	}

	customFinancialProductTraitMapping["http://schema.org/interestRate"] = func(x *schema.FinancialProductTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.InterestRate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.InterestRate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.InterestRate = s
		}
	}
}

func NewFinancialProductFromContext(ctx jsonld.Context) (x schema.FinancialProduct) {
	x.Type = "http://schema.org/FinancialProduct"
	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFinancialProductTrait(ctx, &x.FinancialProductTrait)
	MapCustomToFinancialProductTrait(ctx, &x.FinancialProductTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFinancialProductTrait(ctx jsonld.Context, x *schema.FinancialProductTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFinancialProductTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFinancialProductTrait(ctx jsonld.Context, x *schema.FinancialProductTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFinancialProductTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}