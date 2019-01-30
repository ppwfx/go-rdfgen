package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTakeAction strings.Replacer
var strconvTakeAction strconv.NumError

var basicTakeActionTraitMapping = map[string]func(*schema.TakeActionTrait, []string){}
var customTakeActionTraitMapping = map[string]func(*schema.TakeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TakeAction", func(ctx jsonld.Context) (interface{}) {
		return NewTakeActionFromContext(ctx)
	})

}

func NewTakeActionFromContext(ctx jsonld.Context) (x schema.TakeAction) {
	x.Type = "http://schema.org/TakeAction"
	MapBasicToTransferActionTrait(ctx, &x.TransferActionTrait)
	MapCustomToTransferActionTrait(ctx, &x.TransferActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTakeActionTrait(ctx, &x.TakeActionTrait)
	MapCustomToTakeActionTrait(ctx, &x.TakeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTakeActionTrait(ctx jsonld.Context, x *schema.TakeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTakeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTakeActionTrait(ctx jsonld.Context, x *schema.TakeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTakeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}