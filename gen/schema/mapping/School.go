package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSchool strings.Replacer
var strconvSchool strconv.NumError

var basicSchoolTraitMapping = map[string]func(*schema.SchoolTrait, []string){}
var customSchoolTraitMapping = map[string]func(*schema.SchoolTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/School", func(ctx jsonld.Context) (interface{}) {
		return NewSchoolFromContext(ctx)
	})

}

func NewSchoolFromContext(ctx jsonld.Context) (x schema.School) {
	x.Type = "http://schema.org/School"
	MapBasicToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)
	MapCustomToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSchoolTrait(ctx, &x.SchoolTrait)
	MapCustomToSchoolTrait(ctx, &x.SchoolTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSchoolTrait(ctx jsonld.Context, x *schema.SchoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSchoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSchoolTrait(ctx jsonld.Context, x *schema.SchoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSchoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}