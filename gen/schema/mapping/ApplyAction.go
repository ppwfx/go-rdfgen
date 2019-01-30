package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsApplyAction strings.Replacer
var strconvApplyAction strconv.NumError

var basicApplyActionTraitMapping = map[string]func(*schema.ApplyActionTrait, []string){}
var customApplyActionTraitMapping = map[string]func(*schema.ApplyActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ApplyAction", func(ctx jsonld.Context) (interface{}) {
		return NewApplyActionFromContext(ctx)
	})

}

func NewApplyActionFromContext(ctx jsonld.Context) (x schema.ApplyAction) {
	x.Type = "http://schema.org/ApplyAction"
	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToApplyActionTrait(ctx, &x.ApplyActionTrait)
	MapCustomToApplyActionTrait(ctx, &x.ApplyActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToApplyActionTrait(ctx jsonld.Context, x *schema.ApplyActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicApplyActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToApplyActionTrait(ctx jsonld.Context, x *schema.ApplyActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customApplyActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}