package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBankAccount strings.Replacer
var strconvBankAccount strconv.NumError

var basicBankAccountTraitMapping = map[string]func(*schema.BankAccountTrait, []string){}
var customBankAccountTraitMapping = map[string]func(*schema.BankAccountTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BankAccount", func(ctx jsonld.Context) (interface{}) {
		return NewBankAccountFromContext(ctx)
	})





	customBankAccountTraitMapping["http://schema.org/bankAccountType"] = func(x *schema.BankAccountTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BankAccountType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BankAccountType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BankAccountType = s
		}
	}

	customBankAccountTraitMapping["http://schema.org/accountOverdraftLimit"] = func(x *schema.BankAccountTrait, ctx jsonld.Context, s string){
		var y schema.MonetaryAmount
		if strings.HasPrefix(s, "_:") {
			y = NewMonetaryAmountFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMonetaryAmount()
			y.Id = s
		}

		x.AccountOverdraftLimit = &y

		return
	}

	customBankAccountTraitMapping["http://schema.org/accountMinimumInflow"] = func(x *schema.BankAccountTrait, ctx jsonld.Context, s string){
		var y schema.MonetaryAmount
		if strings.HasPrefix(s, "_:") {
			y = NewMonetaryAmountFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMonetaryAmount()
			y.Id = s
		}

		x.AccountMinimumInflow = &y

		return
	}
}

func NewBankAccountFromContext(ctx jsonld.Context) (x schema.BankAccount) {
	x.Type = "http://schema.org/BankAccount"
	MapBasicToFinancialProductTrait(ctx, &x.FinancialProductTrait)
	MapCustomToFinancialProductTrait(ctx, &x.FinancialProductTrait)

	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBankAccountTrait(ctx, &x.BankAccountTrait)
	MapCustomToBankAccountTrait(ctx, &x.BankAccountTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBankAccountTrait(ctx jsonld.Context, x *schema.BankAccountTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBankAccountTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBankAccountTrait(ctx jsonld.Context, x *schema.BankAccountTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBankAccountTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}