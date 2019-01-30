package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRoofingContractor strings.Replacer
var strconvRoofingContractor strconv.NumError

var basicRoofingContractorTraitMapping = map[string]func(*schema.RoofingContractorTrait, []string){}
var customRoofingContractorTraitMapping = map[string]func(*schema.RoofingContractorTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RoofingContractor", func(ctx jsonld.Context) (interface{}) {
		return NewRoofingContractorFromContext(ctx)
	})

}

func NewRoofingContractorFromContext(ctx jsonld.Context) (x schema.RoofingContractor) {
	x.Type = "http://schema.org/RoofingContractor"
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


	MapBasicToRoofingContractorTrait(ctx, &x.RoofingContractorTrait)
	MapCustomToRoofingContractorTrait(ctx, &x.RoofingContractorTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRoofingContractorTrait(ctx jsonld.Context, x *schema.RoofingContractorTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRoofingContractorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRoofingContractorTrait(ctx jsonld.Context, x *schema.RoofingContractorTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRoofingContractorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}