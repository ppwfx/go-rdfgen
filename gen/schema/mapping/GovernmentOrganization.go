package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGovernmentOrganization strings.Replacer
var strconvGovernmentOrganization strconv.NumError

var basicGovernmentOrganizationTraitMapping = map[string]func(*schema.GovernmentOrganizationTrait, []string){}
var customGovernmentOrganizationTraitMapping = map[string]func(*schema.GovernmentOrganizationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GovernmentOrganization", func(ctx jsonld.Context) (interface{}) {
		return NewGovernmentOrganizationFromContext(ctx)
	})

}

func NewGovernmentOrganizationFromContext(ctx jsonld.Context) (x schema.GovernmentOrganization) {
	x.Type = "http://schema.org/GovernmentOrganization"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGovernmentOrganizationTrait(ctx, &x.GovernmentOrganizationTrait)
	MapCustomToGovernmentOrganizationTrait(ctx, &x.GovernmentOrganizationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGovernmentOrganizationTrait(ctx jsonld.Context, x *schema.GovernmentOrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGovernmentOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGovernmentOrganizationTrait(ctx jsonld.Context, x *schema.GovernmentOrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGovernmentOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}