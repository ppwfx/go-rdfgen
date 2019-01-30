package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPerformanceRole strings.Replacer
var strconvPerformanceRole strconv.NumError

var basicPerformanceRoleTraitMapping = map[string]func(*schema.PerformanceRoleTrait, []string){}
var customPerformanceRoleTraitMapping = map[string]func(*schema.PerformanceRoleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PerformanceRole", func(ctx jsonld.Context) (interface{}) {
		return NewPerformanceRoleFromContext(ctx)
	})


	basicPerformanceRoleTraitMapping["http://schema.org/characterName"] = func(x *schema.PerformanceRoleTrait, s []string) {
		x.CharacterName = s[0]
	}

}

func NewPerformanceRoleFromContext(ctx jsonld.Context) (x schema.PerformanceRole) {
	x.Type = "http://schema.org/PerformanceRole"
	MapBasicToRoleTrait(ctx, &x.RoleTrait)
	MapCustomToRoleTrait(ctx, &x.RoleTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPerformanceRoleTrait(ctx, &x.PerformanceRoleTrait)
	MapCustomToPerformanceRoleTrait(ctx, &x.PerformanceRoleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPerformanceRoleTrait(ctx jsonld.Context, x *schema.PerformanceRoleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPerformanceRoleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPerformanceRoleTrait(ctx jsonld.Context, x *schema.PerformanceRoleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPerformanceRoleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}