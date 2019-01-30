package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOrganizationRole strings.Replacer
var strconvOrganizationRole strconv.NumError

var basicOrganizationRoleTraitMapping = map[string]func(*schema.OrganizationRoleTrait, []string){}
var customOrganizationRoleTraitMapping = map[string]func(*schema.OrganizationRoleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OrganizationRole", func(ctx jsonld.Context) (interface{}) {
		return NewOrganizationRoleFromContext(ctx)
	})


	basicOrganizationRoleTraitMapping["http://schema.org/numberedPosition"] = func(x *schema.OrganizationRoleTrait, s []string) {
		var err error
		x.NumberedPosition, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}

}

func NewOrganizationRoleFromContext(ctx jsonld.Context) (x schema.OrganizationRole) {
	x.Type = "http://schema.org/OrganizationRole"
	MapBasicToRoleTrait(ctx, &x.RoleTrait)
	MapCustomToRoleTrait(ctx, &x.RoleTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOrganizationRoleTrait(ctx, &x.OrganizationRoleTrait)
	MapCustomToOrganizationRoleTrait(ctx, &x.OrganizationRoleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOrganizationRoleTrait(ctx jsonld.Context, x *schema.OrganizationRoleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOrganizationRoleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOrganizationRoleTrait(ctx jsonld.Context, x *schema.OrganizationRoleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOrganizationRoleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}