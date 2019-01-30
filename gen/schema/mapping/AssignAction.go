package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAssignAction strings.Replacer
var strconvAssignAction strconv.NumError

var basicAssignActionTraitMapping = map[string]func(*schema.AssignActionTrait, []string){}
var customAssignActionTraitMapping = map[string]func(*schema.AssignActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AssignAction", func(ctx jsonld.Context) (interface{}) {
		return NewAssignActionFromContext(ctx)
	})

}

func NewAssignActionFromContext(ctx jsonld.Context) (x schema.AssignAction) {
	x.Type = "http://schema.org/AssignAction"
	MapBasicToAllocateActionTrait(ctx, &x.AllocateActionTrait)
	MapCustomToAllocateActionTrait(ctx, &x.AllocateActionTrait)

	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAssignActionTrait(ctx, &x.AssignActionTrait)
	MapCustomToAssignActionTrait(ctx, &x.AssignActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAssignActionTrait(ctx jsonld.Context, x *schema.AssignActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAssignActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAssignActionTrait(ctx jsonld.Context, x *schema.AssignActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAssignActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}