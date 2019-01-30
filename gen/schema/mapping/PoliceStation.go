package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPoliceStation strings.Replacer
var strconvPoliceStation strconv.NumError

var basicPoliceStationTraitMapping = map[string]func(*schema.PoliceStationTrait, []string){}
var customPoliceStationTraitMapping = map[string]func(*schema.PoliceStationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PoliceStation", func(ctx jsonld.Context) (interface{}) {
		return NewPoliceStationFromContext(ctx)
	})

}

func NewPoliceStationFromContext(ctx jsonld.Context) (x schema.PoliceStation) {
	x.Type = "http://schema.org/PoliceStation"
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


	MapBasicToPoliceStationTrait(ctx, &x.PoliceStationTrait)
	MapCustomToPoliceStationTrait(ctx, &x.PoliceStationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPoliceStationTrait(ctx jsonld.Context, x *schema.PoliceStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPoliceStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPoliceStationTrait(ctx jsonld.Context, x *schema.PoliceStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPoliceStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}