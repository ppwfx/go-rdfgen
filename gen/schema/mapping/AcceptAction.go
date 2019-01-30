package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAcceptAction strings.Replacer
var strconvAcceptAction strconv.NumError

var basicAcceptActionTraitMapping = map[string]func(*schema.AcceptActionTrait, []string){}
var customAcceptActionTraitMapping = map[string]func(*schema.AcceptActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AcceptAction", func(ctx jsonld.Context) (interface{}) {
		return NewAcceptActionFromContext(ctx)
	})

}

func NewAcceptActionFromContext(ctx jsonld.Context) (x schema.AcceptAction) {
	x.Type = "http://schema.org/AcceptAction"
	MapBasicToAllocateActionTrait(ctx, &x.AllocateActionTrait)
	MapCustomToAllocateActionTrait(ctx, &x.AllocateActionTrait)

	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAcceptActionTrait(ctx, &x.AcceptActionTrait)
	MapCustomToAcceptActionTrait(ctx, &x.AcceptActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAcceptActionTrait(ctx jsonld.Context, x *schema.AcceptActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAcceptActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAcceptActionTrait(ctx jsonld.Context, x *schema.AcceptActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAcceptActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}