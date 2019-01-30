package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEmploymentAgency strings.Replacer
var strconvEmploymentAgency strconv.NumError

var basicEmploymentAgencyTraitMapping = map[string]func(*schema.EmploymentAgencyTrait, []string){}
var customEmploymentAgencyTraitMapping = map[string]func(*schema.EmploymentAgencyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EmploymentAgency", func(ctx jsonld.Context) (interface{}) {
		return NewEmploymentAgencyFromContext(ctx)
	})

}

func NewEmploymentAgencyFromContext(ctx jsonld.Context) (x schema.EmploymentAgency) {
	x.Type = "http://schema.org/EmploymentAgency"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToEmploymentAgencyTrait(ctx, &x.EmploymentAgencyTrait)
	MapCustomToEmploymentAgencyTrait(ctx, &x.EmploymentAgencyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEmploymentAgencyTrait(ctx jsonld.Context, x *schema.EmploymentAgencyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEmploymentAgencyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEmploymentAgencyTrait(ctx jsonld.Context, x *schema.EmploymentAgencyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEmploymentAgencyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}