package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGasStation strings.Replacer
var strconvGasStation strconv.NumError

var basicGasStationTraitMapping = map[string]func(*schema.GasStationTrait, []string){}
var customGasStationTraitMapping = map[string]func(*schema.GasStationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GasStation", func(ctx jsonld.Context) (interface{}) {
		return NewGasStationFromContext(ctx)
	})

}

func NewGasStationFromContext(ctx jsonld.Context) (x schema.GasStation) {
	x.Type = "http://schema.org/GasStation"
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


	MapBasicToGasStationTrait(ctx, &x.GasStationTrait)
	MapCustomToGasStationTrait(ctx, &x.GasStationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGasStationTrait(ctx jsonld.Context, x *schema.GasStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGasStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGasStationTrait(ctx jsonld.Context, x *schema.GasStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGasStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}