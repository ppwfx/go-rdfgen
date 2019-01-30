package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCivicStructure strings.Replacer
var strconvCivicStructure strconv.NumError

var basicCivicStructureTraitMapping = map[string]func(*schema.CivicStructureTrait, []string){}
var customCivicStructureTraitMapping = map[string]func(*schema.CivicStructureTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CivicStructure", func(ctx jsonld.Context) (interface{}) {
		return NewCivicStructureFromContext(ctx)
	})


	basicCivicStructureTraitMapping["http://schema.org/openingHours"] = func(x *schema.CivicStructureTrait, s []string) {
		x.OpeningHours = s[0]
	}

}

func NewCivicStructureFromContext(ctx jsonld.Context) (x schema.CivicStructure) {
	x.Type = "http://schema.org/CivicStructure"
	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCivicStructureTrait(ctx jsonld.Context, x *schema.CivicStructureTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCivicStructureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCivicStructureTrait(ctx jsonld.Context, x *schema.CivicStructureTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCivicStructureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}