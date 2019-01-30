package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFireStation strings.Replacer
var strconvFireStation strconv.NumError

var basicFireStationTraitMapping = map[string]func(*schema.FireStationTrait, []string){}
var customFireStationTraitMapping = map[string]func(*schema.FireStationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FireStation", func(ctx jsonld.Context) (interface{}) {
		return NewFireStationFromContext(ctx)
	})

}

func NewFireStationFromContext(ctx jsonld.Context) (x schema.FireStation) {
	x.Type = "http://schema.org/FireStation"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToEmergencyServiceTrait(ctx, &x.EmergencyServiceTrait)
	MapCustomToEmergencyServiceTrait(ctx, &x.EmergencyServiceTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToFireStationTrait(ctx, &x.FireStationTrait)
	MapCustomToFireStationTrait(ctx, &x.FireStationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFireStationTrait(ctx jsonld.Context, x *schema.FireStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFireStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFireStationTrait(ctx jsonld.Context, x *schema.FireStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFireStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}