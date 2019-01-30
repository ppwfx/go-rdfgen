package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBedAndBreakfast strings.Replacer
var strconvBedAndBreakfast strconv.NumError

var basicBedAndBreakfastTraitMapping = map[string]func(*schema.BedAndBreakfastTrait, []string){}
var customBedAndBreakfastTraitMapping = map[string]func(*schema.BedAndBreakfastTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BedAndBreakfast", func(ctx jsonld.Context) (interface{}) {
		return NewBedAndBreakfastFromContext(ctx)
	})

}

func NewBedAndBreakfastFromContext(ctx jsonld.Context) (x schema.BedAndBreakfast) {
	x.Type = "http://schema.org/BedAndBreakfast"
	MapBasicToLodgingBusinessTrait(ctx, &x.LodgingBusinessTrait)
	MapCustomToLodgingBusinessTrait(ctx, &x.LodgingBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToBedAndBreakfastTrait(ctx, &x.BedAndBreakfastTrait)
	MapCustomToBedAndBreakfastTrait(ctx, &x.BedAndBreakfastTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBedAndBreakfastTrait(ctx jsonld.Context, x *schema.BedAndBreakfastTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBedAndBreakfastTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBedAndBreakfastTrait(ctx jsonld.Context, x *schema.BedAndBreakfastTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBedAndBreakfastTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}