package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSendAction strings.Replacer
var strconvSendAction strconv.NumError

var basicSendActionTraitMapping = map[string]func(*schema.SendActionTrait, []string){}
var customSendActionTraitMapping = map[string]func(*schema.SendActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SendAction", func(ctx jsonld.Context) (interface{}) {
		return NewSendActionFromContext(ctx)
	})




	customSendActionTraitMapping["http://schema.org/recipient"] = func(x *schema.SendActionTrait, ctx jsonld.Context, s string){
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

	customSendActionTraitMapping["http://schema.org/deliveryMethod"] = func(x *schema.SendActionTrait, ctx jsonld.Context, s string){
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

func NewSendActionFromContext(ctx jsonld.Context) (x schema.SendAction) {
	x.Type = "http://schema.org/SendAction"
	MapBasicToTransferActionTrait(ctx, &x.TransferActionTrait)
	MapCustomToTransferActionTrait(ctx, &x.TransferActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSendActionTrait(ctx, &x.SendActionTrait)
	MapCustomToSendActionTrait(ctx, &x.SendActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSendActionTrait(ctx jsonld.Context, x *schema.SendActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSendActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSendActionTrait(ctx jsonld.Context, x *schema.SendActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSendActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}