package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTravelAgency strings.Replacer
var strconvTravelAgency strconv.NumError

var basicTravelAgencyTraitMapping = map[string]func(*schema.TravelAgencyTrait, []string){}
var customTravelAgencyTraitMapping = map[string]func(*schema.TravelAgencyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TravelAgency", func(ctx jsonld.Context) (interface{}) {
		return NewTravelAgencyFromContext(ctx)
	})

}

func NewTravelAgencyFromContext(ctx jsonld.Context) (x schema.TravelAgency) {
	x.Type = "http://schema.org/TravelAgency"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToTravelAgencyTrait(ctx, &x.TravelAgencyTrait)
	MapCustomToTravelAgencyTrait(ctx, &x.TravelAgencyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTravelAgencyTrait(ctx jsonld.Context, x *schema.TravelAgencyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTravelAgencyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTravelAgencyTrait(ctx jsonld.Context, x *schema.TravelAgencyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTravelAgencyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}