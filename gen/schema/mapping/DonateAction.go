package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDonateAction strings.Replacer
var strconvDonateAction strconv.NumError

var basicDonateActionTraitMapping = map[string]func(*schema.DonateActionTrait, []string){}
var customDonateActionTraitMapping = map[string]func(*schema.DonateActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DonateAction", func(ctx jsonld.Context) (interface{}) {
		return NewDonateActionFromContext(ctx)
	})



	customDonateActionTraitMapping["http://schema.org/recipient"] = func(x *schema.DonateActionTrait, ctx jsonld.Context, s string){
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

func NewDonateActionFromContext(ctx jsonld.Context) (x schema.DonateAction) {
	x.Type = "http://schema.org/DonateAction"
	MapBasicToTradeActionTrait(ctx, &x.TradeActionTrait)
	MapCustomToTradeActionTrait(ctx, &x.TradeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDonateActionTrait(ctx, &x.DonateActionTrait)
	MapCustomToDonateActionTrait(ctx, &x.DonateActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDonateActionTrait(ctx jsonld.Context, x *schema.DonateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDonateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDonateActionTrait(ctx jsonld.Context, x *schema.DonateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDonateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}