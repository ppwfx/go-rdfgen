package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOrganizeAction strings.Replacer
var strconvOrganizeAction strconv.NumError

var basicOrganizeActionTraitMapping = map[string]func(*schema.OrganizeActionTrait, []string){}
var customOrganizeActionTraitMapping = map[string]func(*schema.OrganizeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OrganizeAction", func(ctx jsonld.Context) (interface{}) {
		return NewOrganizeActionFromContext(ctx)
	})

}

func NewOrganizeActionFromContext(ctx jsonld.Context) (x schema.OrganizeAction) {
	x.Type = "http://schema.org/OrganizeAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOrganizeActionTrait(ctx jsonld.Context, x *schema.OrganizeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOrganizeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOrganizeActionTrait(ctx jsonld.Context, x *schema.OrganizeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOrganizeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}