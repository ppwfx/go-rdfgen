package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCampingPitch strings.Replacer
var strconvCampingPitch strconv.NumError

var basicCampingPitchTraitMapping = map[string]func(*schema.CampingPitchTrait, []string){}
var customCampingPitchTraitMapping = map[string]func(*schema.CampingPitchTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CampingPitch", func(ctx jsonld.Context) (interface{}) {
		return NewCampingPitchFromContext(ctx)
	})

}

func NewCampingPitchFromContext(ctx jsonld.Context) (x schema.CampingPitch) {
	x.Type = "http://schema.org/CampingPitch"
	MapBasicToAccommodationTrait(ctx, &x.AccommodationTrait)
	MapCustomToAccommodationTrait(ctx, &x.AccommodationTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCampingPitchTrait(ctx, &x.CampingPitchTrait)
	MapCustomToCampingPitchTrait(ctx, &x.CampingPitchTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCampingPitchTrait(ctx jsonld.Context, x *schema.CampingPitchTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCampingPitchTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCampingPitchTrait(ctx jsonld.Context, x *schema.CampingPitchTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCampingPitchTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}