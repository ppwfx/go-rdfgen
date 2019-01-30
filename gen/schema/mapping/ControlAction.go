package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsControlAction strings.Replacer
var strconvControlAction strconv.NumError

var basicControlActionTraitMapping = map[string]func(*schema.ControlActionTrait, []string){}
var customControlActionTraitMapping = map[string]func(*schema.ControlActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ControlAction", func(ctx jsonld.Context) (interface{}) {
		return NewControlActionFromContext(ctx)
	})

}

func NewControlActionFromContext(ctx jsonld.Context) (x schema.ControlAction) {
	x.Type = "http://schema.org/ControlAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToControlActionTrait(ctx, &x.ControlActionTrait)
	MapCustomToControlActionTrait(ctx, &x.ControlActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToControlActionTrait(ctx jsonld.Context, x *schema.ControlActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicControlActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToControlActionTrait(ctx jsonld.Context, x *schema.ControlActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customControlActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}