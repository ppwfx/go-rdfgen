package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPerformingGroup strings.Replacer
var strconvPerformingGroup strconv.NumError

var basicPerformingGroupTraitMapping = map[string]func(*schema.PerformingGroupTrait, []string){}
var customPerformingGroupTraitMapping = map[string]func(*schema.PerformingGroupTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PerformingGroup", func(ctx jsonld.Context) (interface{}) {
		return NewPerformingGroupFromContext(ctx)
	})

}

func NewPerformingGroupFromContext(ctx jsonld.Context) (x schema.PerformingGroup) {
	x.Type = "http://schema.org/PerformingGroup"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPerformingGroupTrait(ctx, &x.PerformingGroupTrait)
	MapCustomToPerformingGroupTrait(ctx, &x.PerformingGroupTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPerformingGroupTrait(ctx jsonld.Context, x *schema.PerformingGroupTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPerformingGroupTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPerformingGroupTrait(ctx jsonld.Context, x *schema.PerformingGroupTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPerformingGroupTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}