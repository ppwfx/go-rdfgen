package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFinancialService strings.Replacer
var strconvFinancialService strconv.NumError

var basicFinancialServiceTraitMapping = map[string]func(*schema.FinancialServiceTrait, []string){}
var customFinancialServiceTraitMapping = map[string]func(*schema.FinancialServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FinancialService", func(ctx jsonld.Context) (interface{}) {
		return NewFinancialServiceFromContext(ctx)
	})



	customFinancialServiceTraitMapping["http://schema.org/feesAndCommissionsSpecification"] = func(x *schema.FinancialServiceTrait, ctx jsonld.Context, s string){
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
}

func NewFinancialServiceFromContext(ctx jsonld.Context) (x schema.FinancialService) {
	x.Type = "http://schema.org/FinancialService"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToFinancialServiceTrait(ctx, &x.FinancialServiceTrait)
	MapCustomToFinancialServiceTrait(ctx, &x.FinancialServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFinancialServiceTrait(ctx jsonld.Context, x *schema.FinancialServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFinancialServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFinancialServiceTrait(ctx jsonld.Context, x *schema.FinancialServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFinancialServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}