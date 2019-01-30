package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAquarium strings.Replacer
var strconvAquarium strconv.NumError

var basicAquariumTraitMapping = map[string]func(*schema.AquariumTrait, []string){}
var customAquariumTraitMapping = map[string]func(*schema.AquariumTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Aquarium", func(ctx jsonld.Context) (interface{}) {
		return NewAquariumFromContext(ctx)
	})

}

func NewAquariumFromContext(ctx jsonld.Context) (x schema.Aquarium) {
	x.Type = "http://schema.org/Aquarium"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAquariumTrait(ctx, &x.AquariumTrait)
	MapCustomToAquariumTrait(ctx, &x.AquariumTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAquariumTrait(ctx jsonld.Context, x *schema.AquariumTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAquariumTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAquariumTrait(ctx jsonld.Context, x *schema.AquariumTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAquariumTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}