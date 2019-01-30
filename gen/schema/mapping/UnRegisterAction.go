package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUnRegisterAction strings.Replacer
var strconvUnRegisterAction strconv.NumError

var basicUnRegisterActionTraitMapping = map[string]func(*schema.UnRegisterActionTrait, []string){}
var customUnRegisterActionTraitMapping = map[string]func(*schema.UnRegisterActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UnRegisterAction", func(ctx jsonld.Context) (interface{}) {
		return NewUnRegisterActionFromContext(ctx)
	})

}

func NewUnRegisterActionFromContext(ctx jsonld.Context) (x schema.UnRegisterAction) {
	x.Type = "http://schema.org/UnRegisterAction"
	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUnRegisterActionTrait(ctx, &x.UnRegisterActionTrait)
	MapCustomToUnRegisterActionTrait(ctx, &x.UnRegisterActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUnRegisterActionTrait(ctx jsonld.Context, x *schema.UnRegisterActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUnRegisterActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUnRegisterActionTrait(ctx jsonld.Context, x *schema.UnRegisterActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUnRegisterActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}