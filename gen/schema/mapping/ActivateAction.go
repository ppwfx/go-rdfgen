package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsActivateAction strings.Replacer
var strconvActivateAction strconv.NumError

var basicActivateActionTraitMapping = map[string]func(*schema.ActivateActionTrait, []string){}
var customActivateActionTraitMapping = map[string]func(*schema.ActivateActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ActivateAction", func(ctx jsonld.Context) (interface{}) {
		return NewActivateActionFromContext(ctx)
	})

}

func NewActivateActionFromContext(ctx jsonld.Context) (x schema.ActivateAction) {
	x.Type = "http://schema.org/ActivateAction"
	MapBasicToControlActionTrait(ctx, &x.ControlActionTrait)
	MapCustomToControlActionTrait(ctx, &x.ControlActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToActivateActionTrait(ctx, &x.ActivateActionTrait)
	MapCustomToActivateActionTrait(ctx, &x.ActivateActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToActivateActionTrait(ctx jsonld.Context, x *schema.ActivateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicActivateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToActivateActionTrait(ctx jsonld.Context, x *schema.ActivateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customActivateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}