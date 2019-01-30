package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTradeAction strings.Replacer
var strconvTradeAction strconv.NumError

var basicTradeActionTraitMapping = map[string]func(*schema.TradeActionTrait, []string){}
var customTradeActionTraitMapping = map[string]func(*schema.TradeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TradeAction", func(ctx jsonld.Context) (interface{}) {
		return NewTradeActionFromContext(ctx)
	})




	customTradeActionTraitMapping["http://schema.org/priceSpecification"] = func(x *schema.TradeActionTrait, ctx jsonld.Context, s string){
		var y schema.PriceSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewPriceSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPriceSpecification()
			y.Id = s
		}

		x.PriceSpecification = &y

		return
	}

	customTradeActionTraitMapping["http://schema.org/price"] = func(x *schema.TradeActionTrait, ctx jsonld.Context, s string){
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

func NewTradeActionFromContext(ctx jsonld.Context) (x schema.TradeAction) {
	x.Type = "http://schema.org/TradeAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTradeActionTrait(ctx, &x.TradeActionTrait)
	MapCustomToTradeActionTrait(ctx, &x.TradeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTradeActionTrait(ctx jsonld.Context, x *schema.TradeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTradeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTradeActionTrait(ctx jsonld.Context, x *schema.TradeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTradeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}