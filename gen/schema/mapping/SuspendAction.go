package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSuspendAction strings.Replacer
var strconvSuspendAction strconv.NumError

var basicSuspendActionTraitMapping = map[string]func(*schema.SuspendActionTrait, []string){}
var customSuspendActionTraitMapping = map[string]func(*schema.SuspendActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SuspendAction", func(ctx jsonld.Context) (interface{}) {
		return NewSuspendActionFromContext(ctx)
	})

}

func NewSuspendActionFromContext(ctx jsonld.Context) (x schema.SuspendAction) {
	x.Type = "http://schema.org/SuspendAction"
	MapBasicToControlActionTrait(ctx, &x.ControlActionTrait)
	MapCustomToControlActionTrait(ctx, &x.ControlActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSuspendActionTrait(ctx, &x.SuspendActionTrait)
	MapCustomToSuspendActionTrait(ctx, &x.SuspendActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSuspendActionTrait(ctx jsonld.Context, x *schema.SuspendActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSuspendActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSuspendActionTrait(ctx jsonld.Context, x *schema.SuspendActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSuspendActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}