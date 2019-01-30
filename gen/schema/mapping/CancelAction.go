package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCancelAction strings.Replacer
var strconvCancelAction strconv.NumError

var basicCancelActionTraitMapping = map[string]func(*schema.CancelActionTrait, []string){}
var customCancelActionTraitMapping = map[string]func(*schema.CancelActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CancelAction", func(ctx jsonld.Context) (interface{}) {
		return NewCancelActionFromContext(ctx)
	})

}

func NewCancelActionFromContext(ctx jsonld.Context) (x schema.CancelAction) {
	x.Type = "http://schema.org/CancelAction"
	MapBasicToPlanActionTrait(ctx, &x.PlanActionTrait)
	MapCustomToPlanActionTrait(ctx, &x.PlanActionTrait)

	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCancelActionTrait(ctx, &x.CancelActionTrait)
	MapCustomToCancelActionTrait(ctx, &x.CancelActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCancelActionTrait(ctx jsonld.Context, x *schema.CancelActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCancelActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCancelActionTrait(ctx jsonld.Context, x *schema.CancelActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCancelActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}