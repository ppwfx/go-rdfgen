package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGovernmentOffice strings.Replacer
var strconvGovernmentOffice strconv.NumError

var basicGovernmentOfficeTraitMapping = map[string]func(*schema.GovernmentOfficeTrait, []string){}
var customGovernmentOfficeTraitMapping = map[string]func(*schema.GovernmentOfficeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GovernmentOffice", func(ctx jsonld.Context) (interface{}) {
		return NewGovernmentOfficeFromContext(ctx)
	})

}

func NewGovernmentOfficeFromContext(ctx jsonld.Context) (x schema.GovernmentOffice) {
	x.Type = "http://schema.org/GovernmentOffice"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToGovernmentOfficeTrait(ctx, &x.GovernmentOfficeTrait)
	MapCustomToGovernmentOfficeTrait(ctx, &x.GovernmentOfficeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGovernmentOfficeTrait(ctx jsonld.Context, x *schema.GovernmentOfficeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGovernmentOfficeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGovernmentOfficeTrait(ctx jsonld.Context, x *schema.GovernmentOfficeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGovernmentOfficeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}