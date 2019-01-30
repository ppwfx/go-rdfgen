package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHotel strings.Replacer
var strconvHotel strconv.NumError

var basicHotelTraitMapping = map[string]func(*schema.HotelTrait, []string){}
var customHotelTraitMapping = map[string]func(*schema.HotelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Hotel", func(ctx jsonld.Context) (interface{}) {
		return NewHotelFromContext(ctx)
	})

}

func NewHotelFromContext(ctx jsonld.Context) (x schema.Hotel) {
	x.Type = "http://schema.org/Hotel"
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


	MapBasicToHotelTrait(ctx, &x.HotelTrait)
	MapCustomToHotelTrait(ctx, &x.HotelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHotelTrait(ctx jsonld.Context, x *schema.HotelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHotelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHotelTrait(ctx jsonld.Context, x *schema.HotelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHotelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}