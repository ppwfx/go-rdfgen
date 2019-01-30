package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRegisterAction strings.Replacer
var strconvRegisterAction strconv.NumError

var basicRegisterActionTraitMapping = map[string]func(*schema.RegisterActionTrait, []string){}
var customRegisterActionTraitMapping = map[string]func(*schema.RegisterActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RegisterAction", func(ctx jsonld.Context) (interface{}) {
		return NewRegisterActionFromContext(ctx)
	})

}

func NewRegisterActionFromContext(ctx jsonld.Context) (x schema.RegisterAction) {
	x.Type = "http://schema.org/RegisterAction"
	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRegisterActionTrait(ctx, &x.RegisterActionTrait)
	MapCustomToRegisterActionTrait(ctx, &x.RegisterActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRegisterActionTrait(ctx jsonld.Context, x *schema.RegisterActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRegisterActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRegisterActionTrait(ctx jsonld.Context, x *schema.RegisterActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRegisterActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}