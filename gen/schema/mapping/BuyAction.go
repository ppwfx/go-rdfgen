package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBuyAction strings.Replacer
var strconvBuyAction strconv.NumError

var basicBuyActionTraitMapping = map[string]func(*schema.BuyActionTrait, []string){}
var customBuyActionTraitMapping = map[string]func(*schema.BuyActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BuyAction", func(ctx jsonld.Context) (interface{}) {
		return NewBuyActionFromContext(ctx)
	})





	customBuyActionTraitMapping["http://schema.org/seller"] = func(x *schema.BuyActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Seller, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Seller = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Seller = s
		}
	}

	customBuyActionTraitMapping["http://schema.org/vendor"] = func(x *schema.BuyActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Vendor, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Vendor = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Vendor = s
		}
	}

	customBuyActionTraitMapping["http://schema.org/warrantyPromise"] = func(x *schema.BuyActionTrait, ctx jsonld.Context, s string){
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
}

func NewBuyActionFromContext(ctx jsonld.Context) (x schema.BuyAction) {
	x.Type = "http://schema.org/BuyAction"
	MapBasicToTradeActionTrait(ctx, &x.TradeActionTrait)
	MapCustomToTradeActionTrait(ctx, &x.TradeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBuyActionTrait(ctx, &x.BuyActionTrait)
	MapCustomToBuyActionTrait(ctx, &x.BuyActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBuyActionTrait(ctx jsonld.Context, x *schema.BuyActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBuyActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBuyActionTrait(ctx jsonld.Context, x *schema.BuyActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBuyActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}