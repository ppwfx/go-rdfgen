package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTheaterGroup strings.Replacer
var strconvTheaterGroup strconv.NumError

var basicTheaterGroupTraitMapping = map[string]func(*schema.TheaterGroupTrait, []string){}
var customTheaterGroupTraitMapping = map[string]func(*schema.TheaterGroupTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TheaterGroup", func(ctx jsonld.Context) (interface{}) {
		return NewTheaterGroupFromContext(ctx)
	})

}

func NewTheaterGroupFromContext(ctx jsonld.Context) (x schema.TheaterGroup) {
	x.Type = "http://schema.org/TheaterGroup"
	MapBasicToPerformingGroupTrait(ctx, &x.PerformingGroupTrait)
	MapCustomToPerformingGroupTrait(ctx, &x.PerformingGroupTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTheaterGroupTrait(ctx, &x.TheaterGroupTrait)
	MapCustomToTheaterGroupTrait(ctx, &x.TheaterGroupTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTheaterGroupTrait(ctx jsonld.Context, x *schema.TheaterGroupTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTheaterGroupTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTheaterGroupTrait(ctx jsonld.Context, x *schema.TheaterGroupTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTheaterGroupTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}