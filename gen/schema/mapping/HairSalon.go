package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHairSalon strings.Replacer
var strconvHairSalon strconv.NumError

var basicHairSalonTraitMapping = map[string]func(*schema.HairSalonTrait, []string){}
var customHairSalonTraitMapping = map[string]func(*schema.HairSalonTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HairSalon", func(ctx jsonld.Context) (interface{}) {
		return NewHairSalonFromContext(ctx)
	})

}

func NewHairSalonFromContext(ctx jsonld.Context) (x schema.HairSalon) {
	x.Type = "http://schema.org/HairSalon"
	MapBasicToHealthAndBeautyBusinessTrait(ctx, &x.HealthAndBeautyBusinessTrait)
	MapCustomToHealthAndBeautyBusinessTrait(ctx, &x.HealthAndBeautyBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToHairSalonTrait(ctx, &x.HairSalonTrait)
	MapCustomToHairSalonTrait(ctx, &x.HairSalonTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHairSalonTrait(ctx jsonld.Context, x *schema.HairSalonTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHairSalonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHairSalonTrait(ctx jsonld.Context, x *schema.HairSalonTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHairSalonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}