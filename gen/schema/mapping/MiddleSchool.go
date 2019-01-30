package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMiddleSchool strings.Replacer
var strconvMiddleSchool strconv.NumError

var basicMiddleSchoolTraitMapping = map[string]func(*schema.MiddleSchoolTrait, []string){}
var customMiddleSchoolTraitMapping = map[string]func(*schema.MiddleSchoolTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MiddleSchool", func(ctx jsonld.Context) (interface{}) {
		return NewMiddleSchoolFromContext(ctx)
	})

}

func NewMiddleSchoolFromContext(ctx jsonld.Context) (x schema.MiddleSchool) {
	x.Type = "http://schema.org/MiddleSchool"
	MapBasicToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)
	MapCustomToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMiddleSchoolTrait(ctx, &x.MiddleSchoolTrait)
	MapCustomToMiddleSchoolTrait(ctx, &x.MiddleSchoolTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMiddleSchoolTrait(ctx jsonld.Context, x *schema.MiddleSchoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMiddleSchoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMiddleSchoolTrait(ctx jsonld.Context, x *schema.MiddleSchoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMiddleSchoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}