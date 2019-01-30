package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMoneyTransfer strings.Replacer
var strconvMoneyTransfer strconv.NumError

var basicMoneyTransferTraitMapping = map[string]func(*schema.MoneyTransferTrait, []string){}
var customMoneyTransferTraitMapping = map[string]func(*schema.MoneyTransferTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MoneyTransfer", func(ctx jsonld.Context) (interface{}) {
		return NewMoneyTransferFromContext(ctx)
	})




	customMoneyTransferTraitMapping["http://schema.org/amount"] = func(x *schema.MoneyTransferTrait, ctx jsonld.Context, s string){
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

	customMoneyTransferTraitMapping["http://schema.org/beneficiaryBank"] = func(x *schema.MoneyTransferTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BeneficiaryBank, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BeneficiaryBank = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BeneficiaryBank = s
		}
	}
}

func NewMoneyTransferFromContext(ctx jsonld.Context) (x schema.MoneyTransfer) {
	x.Type = "http://schema.org/MoneyTransfer"
	MapBasicToTransferActionTrait(ctx, &x.TransferActionTrait)
	MapCustomToTransferActionTrait(ctx, &x.TransferActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMoneyTransferTrait(ctx, &x.MoneyTransferTrait)
	MapCustomToMoneyTransferTrait(ctx, &x.MoneyTransferTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMoneyTransferTrait(ctx jsonld.Context, x *schema.MoneyTransferTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMoneyTransferTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMoneyTransferTrait(ctx jsonld.Context, x *schema.MoneyTransferTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMoneyTransferTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}