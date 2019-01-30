package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHinduTemple strings.Replacer
var strconvHinduTemple strconv.NumError

var basicHinduTempleTraitMapping = map[string]func(*schema.HinduTempleTrait, []string){}
var customHinduTempleTraitMapping = map[string]func(*schema.HinduTempleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HinduTemple", func(ctx jsonld.Context) (interface{}) {
		return NewHinduTempleFromContext(ctx)
	})

}

func NewHinduTempleFromContext(ctx jsonld.Context) (x schema.HinduTemple) {
	x.Type = "http://schema.org/HinduTemple"
	MapBasicToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)
	MapCustomToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHinduTempleTrait(ctx, &x.HinduTempleTrait)
	MapCustomToHinduTempleTrait(ctx, &x.HinduTempleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHinduTempleTrait(ctx jsonld.Context, x *schema.HinduTempleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHinduTempleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHinduTempleTrait(ctx jsonld.Context, x *schema.HinduTempleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHinduTempleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}