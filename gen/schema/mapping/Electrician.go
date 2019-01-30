package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsElectrician strings.Replacer
var strconvElectrician strconv.NumError

var basicElectricianTraitMapping = map[string]func(*schema.ElectricianTrait, []string){}
var customElectricianTraitMapping = map[string]func(*schema.ElectricianTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Electrician", func(ctx jsonld.Context) (interface{}) {
		return NewElectricianFromContext(ctx)
	})

}

func NewElectricianFromContext(ctx jsonld.Context) (x schema.Electrician) {
	x.Type = "http://schema.org/Electrician"
	MapBasicToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)
	MapCustomToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToElectricianTrait(ctx, &x.ElectricianTrait)
	MapCustomToElectricianTrait(ctx, &x.ElectricianTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToElectricianTrait(ctx jsonld.Context, x *schema.ElectricianTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicElectricianTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToElectricianTrait(ctx jsonld.Context, x *schema.ElectricianTrait) () {
	for k, v := range ctx.Current {
		f, ok := customElectricianTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}