package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGeneralContractor strings.Replacer
var strconvGeneralContractor strconv.NumError

var basicGeneralContractorTraitMapping = map[string]func(*schema.GeneralContractorTrait, []string){}
var customGeneralContractorTraitMapping = map[string]func(*schema.GeneralContractorTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GeneralContractor", func(ctx jsonld.Context) (interface{}) {
		return NewGeneralContractorFromContext(ctx)
	})

}

func NewGeneralContractorFromContext(ctx jsonld.Context) (x schema.GeneralContractor) {
	x.Type = "http://schema.org/GeneralContractor"
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


	MapBasicToGeneralContractorTrait(ctx, &x.GeneralContractorTrait)
	MapCustomToGeneralContractorTrait(ctx, &x.GeneralContractorTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGeneralContractorTrait(ctx jsonld.Context, x *schema.GeneralContractorTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGeneralContractorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGeneralContractorTrait(ctx jsonld.Context, x *schema.GeneralContractorTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGeneralContractorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}