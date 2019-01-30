package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEmergencyService strings.Replacer
var strconvEmergencyService strconv.NumError

var basicEmergencyServiceTraitMapping = map[string]func(*schema.EmergencyServiceTrait, []string){}
var customEmergencyServiceTraitMapping = map[string]func(*schema.EmergencyServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EmergencyService", func(ctx jsonld.Context) (interface{}) {
		return NewEmergencyServiceFromContext(ctx)
	})

}

func NewEmergencyServiceFromContext(ctx jsonld.Context) (x schema.EmergencyService) {
	x.Type = "http://schema.org/EmergencyService"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToEmergencyServiceTrait(ctx, &x.EmergencyServiceTrait)
	MapCustomToEmergencyServiceTrait(ctx, &x.EmergencyServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEmergencyServiceTrait(ctx jsonld.Context, x *schema.EmergencyServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEmergencyServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEmergencyServiceTrait(ctx jsonld.Context, x *schema.EmergencyServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEmergencyServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}