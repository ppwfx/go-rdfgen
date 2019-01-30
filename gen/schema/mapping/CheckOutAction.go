package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCheckOutAction strings.Replacer
var strconvCheckOutAction strconv.NumError

var basicCheckOutActionTraitMapping = map[string]func(*schema.CheckOutActionTrait, []string){}
var customCheckOutActionTraitMapping = map[string]func(*schema.CheckOutActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CheckOutAction", func(ctx jsonld.Context) (interface{}) {
		return NewCheckOutActionFromContext(ctx)
	})

}

func NewCheckOutActionFromContext(ctx jsonld.Context) (x schema.CheckOutAction) {
	x.Type = "http://schema.org/CheckOutAction"
	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCheckOutActionTrait(ctx, &x.CheckOutActionTrait)
	MapCustomToCheckOutActionTrait(ctx, &x.CheckOutActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCheckOutActionTrait(ctx jsonld.Context, x *schema.CheckOutActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCheckOutActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCheckOutActionTrait(ctx jsonld.Context, x *schema.CheckOutActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCheckOutActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}