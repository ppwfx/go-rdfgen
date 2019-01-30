package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTipAction strings.Replacer
var strconvTipAction strconv.NumError

var basicTipActionTraitMapping = map[string]func(*schema.TipActionTrait, []string){}
var customTipActionTraitMapping = map[string]func(*schema.TipActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TipAction", func(ctx jsonld.Context) (interface{}) {
		return NewTipActionFromContext(ctx)
	})



	customTipActionTraitMapping["http://schema.org/recipient"] = func(x *schema.TipActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Recipient, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Recipient = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Recipient = s
		}
	}
}

func NewTipActionFromContext(ctx jsonld.Context) (x schema.TipAction) {
	x.Type = "http://schema.org/TipAction"
	MapBasicToTradeActionTrait(ctx, &x.TradeActionTrait)
	MapCustomToTradeActionTrait(ctx, &x.TradeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTipActionTrait(ctx, &x.TipActionTrait)
	MapCustomToTipActionTrait(ctx, &x.TipActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTipActionTrait(ctx jsonld.Context, x *schema.TipActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTipActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTipActionTrait(ctx jsonld.Context, x *schema.TipActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTipActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}