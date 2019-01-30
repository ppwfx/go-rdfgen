package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDeactivateAction strings.Replacer
var strconvDeactivateAction strconv.NumError

var basicDeactivateActionTraitMapping = map[string]func(*schema.DeactivateActionTrait, []string){}
var customDeactivateActionTraitMapping = map[string]func(*schema.DeactivateActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DeactivateAction", func(ctx jsonld.Context) (interface{}) {
		return NewDeactivateActionFromContext(ctx)
	})

}

func NewDeactivateActionFromContext(ctx jsonld.Context) (x schema.DeactivateAction) {
	x.Type = "http://schema.org/DeactivateAction"
	MapBasicToControlActionTrait(ctx, &x.ControlActionTrait)
	MapCustomToControlActionTrait(ctx, &x.ControlActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDeactivateActionTrait(ctx, &x.DeactivateActionTrait)
	MapCustomToDeactivateActionTrait(ctx, &x.DeactivateActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDeactivateActionTrait(ctx jsonld.Context, x *schema.DeactivateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDeactivateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDeactivateActionTrait(ctx jsonld.Context, x *schema.DeactivateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDeactivateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}