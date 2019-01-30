package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsQuoteAction strings.Replacer
var strconvQuoteAction strconv.NumError

var basicQuoteActionTraitMapping = map[string]func(*schema.QuoteActionTrait, []string){}
var customQuoteActionTraitMapping = map[string]func(*schema.QuoteActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/QuoteAction", func(ctx jsonld.Context) (interface{}) {
		return NewQuoteActionFromContext(ctx)
	})

}

func NewQuoteActionFromContext(ctx jsonld.Context) (x schema.QuoteAction) {
	x.Type = "http://schema.org/QuoteAction"
	MapBasicToTradeActionTrait(ctx, &x.TradeActionTrait)
	MapCustomToTradeActionTrait(ctx, &x.TradeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToQuoteActionTrait(ctx, &x.QuoteActionTrait)
	MapCustomToQuoteActionTrait(ctx, &x.QuoteActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToQuoteActionTrait(ctx jsonld.Context, x *schema.QuoteActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicQuoteActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToQuoteActionTrait(ctx jsonld.Context, x *schema.QuoteActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customQuoteActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}