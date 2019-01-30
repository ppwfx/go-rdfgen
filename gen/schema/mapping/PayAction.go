package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPayAction strings.Replacer
var strconvPayAction strconv.NumError

var basicPayActionTraitMapping = map[string]func(*schema.PayActionTrait, []string){}
var customPayActionTraitMapping = map[string]func(*schema.PayActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PayAction", func(ctx jsonld.Context) (interface{}) {
		return NewPayActionFromContext(ctx)
	})




	customPayActionTraitMapping["http://schema.org/recipient"] = func(x *schema.PayActionTrait, ctx jsonld.Context, s string){
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

	customPayActionTraitMapping["http://schema.org/purpose"] = func(x *schema.PayActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Purpose, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Purpose = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Purpose = s
		}
	}
}

func NewPayActionFromContext(ctx jsonld.Context) (x schema.PayAction) {
	x.Type = "http://schema.org/PayAction"
	MapBasicToTradeActionTrait(ctx, &x.TradeActionTrait)
	MapCustomToTradeActionTrait(ctx, &x.TradeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPayActionTrait(ctx, &x.PayActionTrait)
	MapCustomToPayActionTrait(ctx, &x.PayActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPayActionTrait(ctx jsonld.Context, x *schema.PayActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPayActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPayActionTrait(ctx jsonld.Context, x *schema.PayActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPayActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}