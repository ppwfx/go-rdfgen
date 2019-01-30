package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsConfirmAction strings.Replacer
var strconvConfirmAction strconv.NumError

var basicConfirmActionTraitMapping = map[string]func(*schema.ConfirmActionTrait, []string){}
var customConfirmActionTraitMapping = map[string]func(*schema.ConfirmActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ConfirmAction", func(ctx jsonld.Context) (interface{}) {
		return NewConfirmActionFromContext(ctx)
	})

}

func NewConfirmActionFromContext(ctx jsonld.Context) (x schema.ConfirmAction) {
	x.Type = "http://schema.org/ConfirmAction"
	MapBasicToInformActionTrait(ctx, &x.InformActionTrait)
	MapCustomToInformActionTrait(ctx, &x.InformActionTrait)

	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToConfirmActionTrait(ctx, &x.ConfirmActionTrait)
	MapCustomToConfirmActionTrait(ctx, &x.ConfirmActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToConfirmActionTrait(ctx jsonld.Context, x *schema.ConfirmActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicConfirmActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToConfirmActionTrait(ctx jsonld.Context, x *schema.ConfirmActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customConfirmActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}