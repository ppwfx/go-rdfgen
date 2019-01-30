package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMotorcycleRepair strings.Replacer
var strconvMotorcycleRepair strconv.NumError

var basicMotorcycleRepairTraitMapping = map[string]func(*schema.MotorcycleRepairTrait, []string){}
var customMotorcycleRepairTraitMapping = map[string]func(*schema.MotorcycleRepairTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MotorcycleRepair", func(ctx jsonld.Context) (interface{}) {
		return NewMotorcycleRepairFromContext(ctx)
	})

}

func NewMotorcycleRepairFromContext(ctx jsonld.Context) (x schema.MotorcycleRepair) {
	x.Type = "http://schema.org/MotorcycleRepair"
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


	MapBasicToMotorcycleRepairTrait(ctx, &x.MotorcycleRepairTrait)
	MapCustomToMotorcycleRepairTrait(ctx, &x.MotorcycleRepairTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMotorcycleRepairTrait(ctx jsonld.Context, x *schema.MotorcycleRepairTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMotorcycleRepairTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMotorcycleRepairTrait(ctx jsonld.Context, x *schema.MotorcycleRepairTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMotorcycleRepairTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}