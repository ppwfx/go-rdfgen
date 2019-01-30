package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReturnAction strings.Replacer
var strconvReturnAction strconv.NumError

var basicReturnActionTraitMapping = map[string]func(*schema.ReturnActionTrait, []string){}
var customReturnActionTraitMapping = map[string]func(*schema.ReturnActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReturnAction", func(ctx jsonld.Context) (interface{}) {
		return NewReturnActionFromContext(ctx)
	})



	customReturnActionTraitMapping["http://schema.org/recipient"] = func(x *schema.ReturnActionTrait, ctx jsonld.Context, s string){
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

func NewReturnActionFromContext(ctx jsonld.Context) (x schema.ReturnAction) {
	x.Type = "http://schema.org/ReturnAction"
	MapBasicToTransferActionTrait(ctx, &x.TransferActionTrait)
	MapCustomToTransferActionTrait(ctx, &x.TransferActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReturnActionTrait(ctx, &x.ReturnActionTrait)
	MapCustomToReturnActionTrait(ctx, &x.ReturnActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReturnActionTrait(ctx jsonld.Context, x *schema.ReturnActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReturnActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReturnActionTrait(ctx jsonld.Context, x *schema.ReturnActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReturnActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}