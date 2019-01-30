package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCatholicChurch strings.Replacer
var strconvCatholicChurch strconv.NumError

var basicCatholicChurchTraitMapping = map[string]func(*schema.CatholicChurchTrait, []string){}
var customCatholicChurchTraitMapping = map[string]func(*schema.CatholicChurchTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CatholicChurch", func(ctx jsonld.Context) (interface{}) {
		return NewCatholicChurchFromContext(ctx)
	})

}

func NewCatholicChurchFromContext(ctx jsonld.Context) (x schema.CatholicChurch) {
	x.Type = "http://schema.org/CatholicChurch"
	MapBasicToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)
	MapCustomToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCatholicChurchTrait(ctx, &x.CatholicChurchTrait)
	MapCustomToCatholicChurchTrait(ctx, &x.CatholicChurchTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCatholicChurchTrait(ctx jsonld.Context, x *schema.CatholicChurchTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCatholicChurchTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCatholicChurchTrait(ctx jsonld.Context, x *schema.CatholicChurchTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCatholicChurchTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}