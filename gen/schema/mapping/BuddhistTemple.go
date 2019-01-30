package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBuddhistTemple strings.Replacer
var strconvBuddhistTemple strconv.NumError

var basicBuddhistTempleTraitMapping = map[string]func(*schema.BuddhistTempleTrait, []string){}
var customBuddhistTempleTraitMapping = map[string]func(*schema.BuddhistTempleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BuddhistTemple", func(ctx jsonld.Context) (interface{}) {
		return NewBuddhistTempleFromContext(ctx)
	})

}

func NewBuddhistTempleFromContext(ctx jsonld.Context) (x schema.BuddhistTemple) {
	x.Type = "http://schema.org/BuddhistTemple"
	MapBasicToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)
	MapCustomToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBuddhistTempleTrait(ctx, &x.BuddhistTempleTrait)
	MapCustomToBuddhistTempleTrait(ctx, &x.BuddhistTempleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBuddhistTempleTrait(ctx jsonld.Context, x *schema.BuddhistTempleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBuddhistTempleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBuddhistTempleTrait(ctx jsonld.Context, x *schema.BuddhistTempleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBuddhistTempleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}