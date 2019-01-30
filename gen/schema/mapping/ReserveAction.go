package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReserveAction strings.Replacer
var strconvReserveAction strconv.NumError

var basicReserveActionTraitMapping = map[string]func(*schema.ReserveActionTrait, []string){}
var customReserveActionTraitMapping = map[string]func(*schema.ReserveActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReserveAction", func(ctx jsonld.Context) (interface{}) {
		return NewReserveActionFromContext(ctx)
	})

}

func NewReserveActionFromContext(ctx jsonld.Context) (x schema.ReserveAction) {
	x.Type = "http://schema.org/ReserveAction"
	MapBasicToPlanActionTrait(ctx, &x.PlanActionTrait)
	MapCustomToPlanActionTrait(ctx, &x.PlanActionTrait)

	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReserveActionTrait(ctx, &x.ReserveActionTrait)
	MapCustomToReserveActionTrait(ctx, &x.ReserveActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReserveActionTrait(ctx jsonld.Context, x *schema.ReserveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReserveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReserveActionTrait(ctx jsonld.Context, x *schema.ReserveActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReserveActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}