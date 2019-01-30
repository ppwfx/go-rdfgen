package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReceiveAction strings.Replacer
var strconvReceiveAction strconv.NumError

var basicReceiveActionTraitMapping = map[string]func(*schema.ReceiveActionTrait, []string){}
var customReceiveActionTraitMapping = map[string]func(*schema.ReceiveActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReceiveAction", func(ctx jsonld.Context) (interface{}) {
		return NewReceiveActionFromContext(ctx)
	})




	customReceiveActionTraitMapping["http://schema.org/sender"] = func(x *schema.ReceiveActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Sender, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Sender = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Sender = s
		}
	}

	customReceiveActionTraitMapping["http://schema.org/deliveryMethod"] = func(x *schema.ReceiveActionTrait, ctx jsonld.Context, s string){
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

func NewReceiveActionFromContext(ctx jsonld.Context) (x schema.ReceiveAction) {
	x.Type = "http://schema.org/ReceiveAction"
	MapBasicToTransferActionTrait(ctx, &x.TransferActionTrait)
	MapCustomToTransferActionTrait(ctx, &x.TransferActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReceiveActionTrait(ctx, &x.ReceiveActionTrait)
	MapCustomToReceiveActionTrait(ctx, &x.ReceiveActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReceiveActionTrait(ctx jsonld.Context, x *schema.ReceiveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReceiveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReceiveActionTrait(ctx jsonld.Context, x *schema.ReceiveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReceiveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}