package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCourthouse strings.Replacer
var strconvCourthouse strconv.NumError

var basicCourthouseTraitMapping = map[string]func(*schema.CourthouseTrait, []string){}
var customCourthouseTraitMapping = map[string]func(*schema.CourthouseTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Courthouse", func(ctx jsonld.Context) (interface{}) {
		return NewCourthouseFromContext(ctx)
	})

}

func NewCourthouseFromContext(ctx jsonld.Context) (x schema.Courthouse) {
	x.Type = "http://schema.org/Courthouse"
	MapBasicToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)
	MapCustomToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCourthouseTrait(ctx, &x.CourthouseTrait)
	MapCustomToCourthouseTrait(ctx, &x.CourthouseTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCourthouseTrait(ctx jsonld.Context, x *schema.CourthouseTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCourthouseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCourthouseTrait(ctx jsonld.Context, x *schema.CourthouseTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCourthouseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}