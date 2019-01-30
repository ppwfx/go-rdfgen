package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOrderAction strings.Replacer
var strconvOrderAction strconv.NumError

var basicOrderActionTraitMapping = map[string]func(*schema.OrderActionTrait, []string){}
var customOrderActionTraitMapping = map[string]func(*schema.OrderActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OrderAction", func(ctx jsonld.Context) (interface{}) {
		return NewOrderActionFromContext(ctx)
	})



	customOrderActionTraitMapping["http://schema.org/deliveryMethod"] = func(x *schema.OrderActionTrait, ctx jsonld.Context, s string){
		var y schema.DeliveryMethod
		if strings.HasPrefix(s, "_:") {
			y = NewDeliveryMethodFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDeliveryMethod()
			y.Id = s
		}

		x.DeliveryMethod = &y

		return
	}
}

func NewOrderActionFromContext(ctx jsonld.Context) (x schema.OrderAction) {
	x.Type = "http://schema.org/OrderAction"
	MapBasicToTradeActionTrait(ctx, &x.TradeActionTrait)
	MapCustomToTradeActionTrait(ctx, &x.TradeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOrderActionTrait(ctx, &x.OrderActionTrait)
	MapCustomToOrderActionTrait(ctx, &x.OrderActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOrderActionTrait(ctx jsonld.Context, x *schema.OrderActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOrderActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOrderActionTrait(ctx jsonld.Context, x *schema.OrderActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOrderActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}