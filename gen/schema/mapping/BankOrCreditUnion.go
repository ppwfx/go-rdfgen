package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBankOrCreditUnion strings.Replacer
var strconvBankOrCreditUnion strconv.NumError

var basicBankOrCreditUnionTraitMapping = map[string]func(*schema.BankOrCreditUnionTrait, []string){}
var customBankOrCreditUnionTraitMapping = map[string]func(*schema.BankOrCreditUnionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BankOrCreditUnion", func(ctx jsonld.Context) (interface{}) {
		return NewBankOrCreditUnionFromContext(ctx)
	})

}

func NewBankOrCreditUnionFromContext(ctx jsonld.Context) (x schema.BankOrCreditUnion) {
	x.Type = "http://schema.org/BankOrCreditUnion"
	MapBasicToFinancialServiceTrait(ctx, &x.FinancialServiceTrait)
	MapCustomToFinancialServiceTrait(ctx, &x.FinancialServiceTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToBankOrCreditUnionTrait(ctx, &x.BankOrCreditUnionTrait)
	MapCustomToBankOrCreditUnionTrait(ctx, &x.BankOrCreditUnionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBankOrCreditUnionTrait(ctx jsonld.Context, x *schema.BankOrCreditUnionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBankOrCreditUnionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBankOrCreditUnionTrait(ctx jsonld.Context, x *schema.BankOrCreditUnionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBankOrCreditUnionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}