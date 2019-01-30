package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHighSchool strings.Replacer
var strconvHighSchool strconv.NumError

var basicHighSchoolTraitMapping = map[string]func(*schema.HighSchoolTrait, []string){}
var customHighSchoolTraitMapping = map[string]func(*schema.HighSchoolTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HighSchool", func(ctx jsonld.Context) (interface{}) {
		return NewHighSchoolFromContext(ctx)
	})

}

func NewHighSchoolFromContext(ctx jsonld.Context) (x schema.HighSchool) {
	x.Type = "http://schema.org/HighSchool"
	MapBasicToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)
	MapCustomToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHighSchoolTrait(ctx, &x.HighSchoolTrait)
	MapCustomToHighSchoolTrait(ctx, &x.HighSchoolTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHighSchoolTrait(ctx jsonld.Context, x *schema.HighSchoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHighSchoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHighSchoolTrait(ctx jsonld.Context, x *schema.HighSchoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHighSchoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}