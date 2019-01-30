package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDatedMoneySpecification strings.Replacer
var strconvDatedMoneySpecification strconv.NumError

var basicDatedMoneySpecificationTraitMapping = map[string]func(*schema.DatedMoneySpecificationTrait, []string){}
var customDatedMoneySpecificationTraitMapping = map[string]func(*schema.DatedMoneySpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DatedMoneySpecification", func(ctx jsonld.Context) (interface{}) {
		return NewDatedMoneySpecificationFromContext(ctx)
	})




	basicDatedMoneySpecificationTraitMapping["http://schema.org/currency"] = func(x *schema.DatedMoneySpecificationTrait, s []string) {
		x.Currency = s[0]
	}



	customDatedMoneySpecificationTraitMapping["http://schema.org/endDate"] = func(x *schema.DatedMoneySpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EndDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EndDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EndDate = s
		}
	}

	customDatedMoneySpecificationTraitMapping["http://schema.org/startDate"] = func(x *schema.DatedMoneySpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.StartDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.StartDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.StartDate = s
		}
	}

	customDatedMoneySpecificationTraitMapping["http://schema.org/amount"] = func(x *schema.DatedMoneySpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Amount, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Amount = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Amount = s
		}
	}
}

func NewDatedMoneySpecificationFromContext(ctx jsonld.Context) (x schema.DatedMoneySpecification) {
	x.Type = "http://schema.org/DatedMoneySpecification"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDatedMoneySpecificationTrait(ctx, &x.DatedMoneySpecificationTrait)
	MapCustomToDatedMoneySpecificationTrait(ctx, &x.DatedMoneySpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDatedMoneySpecificationTrait(ctx jsonld.Context, x *schema.DatedMoneySpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDatedMoneySpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDatedMoneySpecificationTrait(ctx jsonld.Context, x *schema.DatedMoneySpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDatedMoneySpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}