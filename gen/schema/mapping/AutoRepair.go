package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAutoRepair strings.Replacer
var strconvAutoRepair strconv.NumError

var basicAutoRepairTraitMapping = map[string]func(*schema.AutoRepairTrait, []string){}
var customAutoRepairTraitMapping = map[string]func(*schema.AutoRepairTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AutoRepair", func(ctx jsonld.Context) (interface{}) {
		return NewAutoRepairFromContext(ctx)
	})

}

func NewAutoRepairFromContext(ctx jsonld.Context) (x schema.AutoRepair) {
	x.Type = "http://schema.org/AutoRepair"
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


	MapBasicToAutoRepairTrait(ctx, &x.AutoRepairTrait)
	MapCustomToAutoRepairTrait(ctx, &x.AutoRepairTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAutoRepairTrait(ctx jsonld.Context, x *schema.AutoRepairTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAutoRepairTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAutoRepairTrait(ctx jsonld.Context, x *schema.AutoRepairTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAutoRepairTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}