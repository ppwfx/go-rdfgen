package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsArriveAction strings.Replacer
var strconvArriveAction strconv.NumError

var basicArriveActionTraitMapping = map[string]func(*schema.ArriveActionTrait, []string){}
var customArriveActionTraitMapping = map[string]func(*schema.ArriveActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ArriveAction", func(ctx jsonld.Context) (interface{}) {
		return NewArriveActionFromContext(ctx)
	})

}

func NewArriveActionFromContext(ctx jsonld.Context) (x schema.ArriveAction) {
	x.Type = "http://schema.org/ArriveAction"
	MapBasicToMoveActionTrait(ctx, &x.MoveActionTrait)
	MapCustomToMoveActionTrait(ctx, &x.MoveActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToArriveActionTrait(ctx, &x.ArriveActionTrait)
	MapCustomToArriveActionTrait(ctx, &x.ArriveActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToArriveActionTrait(ctx jsonld.Context, x *schema.ArriveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicArriveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToArriveActionTrait(ctx jsonld.Context, x *schema.ArriveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customArriveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}