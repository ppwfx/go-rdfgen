package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSubscribeAction strings.Replacer
var strconvSubscribeAction strconv.NumError

var basicSubscribeActionTraitMapping = map[string]func(*schema.SubscribeActionTrait, []string){}
var customSubscribeActionTraitMapping = map[string]func(*schema.SubscribeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SubscribeAction", func(ctx jsonld.Context) (interface{}) {
		return NewSubscribeActionFromContext(ctx)
	})

}

func NewSubscribeActionFromContext(ctx jsonld.Context) (x schema.SubscribeAction) {
	x.Type = "http://schema.org/SubscribeAction"
	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSubscribeActionTrait(ctx, &x.SubscribeActionTrait)
	MapCustomToSubscribeActionTrait(ctx, &x.SubscribeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSubscribeActionTrait(ctx jsonld.Context, x *schema.SubscribeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSubscribeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSubscribeActionTrait(ctx jsonld.Context, x *schema.SubscribeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSubscribeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}