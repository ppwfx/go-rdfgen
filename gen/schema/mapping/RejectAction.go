package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRejectAction strings.Replacer
var strconvRejectAction strconv.NumError

var basicRejectActionTraitMapping = map[string]func(*schema.RejectActionTrait, []string){}
var customRejectActionTraitMapping = map[string]func(*schema.RejectActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RejectAction", func(ctx jsonld.Context) (interface{}) {
		return NewRejectActionFromContext(ctx)
	})

}

func NewRejectActionFromContext(ctx jsonld.Context) (x schema.RejectAction) {
	x.Type = "http://schema.org/RejectAction"
	MapBasicToAllocateActionTrait(ctx, &x.AllocateActionTrait)
	MapCustomToAllocateActionTrait(ctx, &x.AllocateActionTrait)

	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRejectActionTrait(ctx, &x.RejectActionTrait)
	MapCustomToRejectActionTrait(ctx, &x.RejectActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRejectActionTrait(ctx jsonld.Context, x *schema.RejectActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRejectActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRejectActionTrait(ctx jsonld.Context, x *schema.RejectActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRejectActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}