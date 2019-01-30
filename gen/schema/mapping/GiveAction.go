package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGiveAction strings.Replacer
var strconvGiveAction strconv.NumError

var basicGiveActionTraitMapping = map[string]func(*schema.GiveActionTrait, []string){}
var customGiveActionTraitMapping = map[string]func(*schema.GiveActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GiveAction", func(ctx jsonld.Context) (interface{}) {
		return NewGiveActionFromContext(ctx)
	})



	customGiveActionTraitMapping["http://schema.org/recipient"] = func(x *schema.GiveActionTrait, ctx jsonld.Context, s string){
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

func NewGiveActionFromContext(ctx jsonld.Context) (x schema.GiveAction) {
	x.Type = "http://schema.org/GiveAction"
	MapBasicToTransferActionTrait(ctx, &x.TransferActionTrait)
	MapCustomToTransferActionTrait(ctx, &x.TransferActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGiveActionTrait(ctx, &x.GiveActionTrait)
	MapCustomToGiveActionTrait(ctx, &x.GiveActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGiveActionTrait(ctx jsonld.Context, x *schema.GiveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGiveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGiveActionTrait(ctx jsonld.Context, x *schema.GiveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGiveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}