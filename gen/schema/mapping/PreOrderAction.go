package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPreOrderAction strings.Replacer
var strconvPreOrderAction strconv.NumError

var basicPreOrderActionTraitMapping = map[string]func(*schema.PreOrderActionTrait, []string){}
var customPreOrderActionTraitMapping = map[string]func(*schema.PreOrderActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PreOrderAction", func(ctx jsonld.Context) (interface{}) {
		return NewPreOrderActionFromContext(ctx)
	})

}

func NewPreOrderActionFromContext(ctx jsonld.Context) (x schema.PreOrderAction) {
	x.Type = "http://schema.org/PreOrderAction"
	MapBasicToTradeActionTrait(ctx, &x.TradeActionTrait)
	MapCustomToTradeActionTrait(ctx, &x.TradeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPreOrderActionTrait(ctx, &x.PreOrderActionTrait)
	MapCustomToPreOrderActionTrait(ctx, &x.PreOrderActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPreOrderActionTrait(ctx jsonld.Context, x *schema.PreOrderActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPreOrderActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPreOrderActionTrait(ctx jsonld.Context, x *schema.PreOrderActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPreOrderActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}