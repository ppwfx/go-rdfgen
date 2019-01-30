package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSellAction strings.Replacer
var strconvSellAction strconv.NumError

var basicSellActionTraitMapping = map[string]func(*schema.SellActionTrait, []string){}
var customSellActionTraitMapping = map[string]func(*schema.SellActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SellAction", func(ctx jsonld.Context) (interface{}) {
		return NewSellActionFromContext(ctx)
	})




	customSellActionTraitMapping["http://schema.org/warrantyPromise"] = func(x *schema.SellActionTrait, ctx jsonld.Context, s string){
		var y schema.WarrantyPromise
		if strings.HasPrefix(s, "_:") {
			y = NewWarrantyPromiseFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewWarrantyPromise()
			y.Id = s
		}

		x.WarrantyPromise = &y

		return
	}

	customSellActionTraitMapping["http://schema.org/buyer"] = func(x *schema.SellActionTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Buyer = &y

		return
	}
}

func NewSellActionFromContext(ctx jsonld.Context) (x schema.SellAction) {
	x.Type = "http://schema.org/SellAction"
	MapBasicToTradeActionTrait(ctx, &x.TradeActionTrait)
	MapCustomToTradeActionTrait(ctx, &x.TradeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSellActionTrait(ctx, &x.SellActionTrait)
	MapCustomToSellActionTrait(ctx, &x.SellActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSellActionTrait(ctx jsonld.Context, x *schema.SellActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSellActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSellActionTrait(ctx jsonld.Context, x *schema.SellActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSellActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}