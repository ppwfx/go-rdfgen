package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDanceGroup strings.Replacer
var strconvDanceGroup strconv.NumError

var basicDanceGroupTraitMapping = map[string]func(*schema.DanceGroupTrait, []string){}
var customDanceGroupTraitMapping = map[string]func(*schema.DanceGroupTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DanceGroup", func(ctx jsonld.Context) (interface{}) {
		return NewDanceGroupFromContext(ctx)
	})

}

func NewDanceGroupFromContext(ctx jsonld.Context) (x schema.DanceGroup) {
	x.Type = "http://schema.org/DanceGroup"
	MapBasicToPerformingGroupTrait(ctx, &x.PerformingGroupTrait)
	MapCustomToPerformingGroupTrait(ctx, &x.PerformingGroupTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDanceGroupTrait(ctx, &x.DanceGroupTrait)
	MapCustomToDanceGroupTrait(ctx, &x.DanceGroupTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDanceGroupTrait(ctx jsonld.Context, x *schema.DanceGroupTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDanceGroupTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDanceGroupTrait(ctx jsonld.Context, x *schema.DanceGroupTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDanceGroupTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}