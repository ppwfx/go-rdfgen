package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAutoRental strings.Replacer
var strconvAutoRental strconv.NumError

var basicAutoRentalTraitMapping = map[string]func(*schema.AutoRentalTrait, []string){}
var customAutoRentalTraitMapping = map[string]func(*schema.AutoRentalTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AutoRental", func(ctx jsonld.Context) (interface{}) {
		return NewAutoRentalFromContext(ctx)
	})

}

func NewAutoRentalFromContext(ctx jsonld.Context) (x schema.AutoRental) {
	x.Type = "http://schema.org/AutoRental"
	MapBasicToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)
	MapCustomToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToAutoRentalTrait(ctx, &x.AutoRentalTrait)
	MapCustomToAutoRentalTrait(ctx, &x.AutoRentalTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAutoRentalTrait(ctx jsonld.Context, x *schema.AutoRentalTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAutoRentalTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAutoRentalTrait(ctx jsonld.Context, x *schema.AutoRentalTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAutoRentalTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}