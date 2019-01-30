package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsElementarySchool strings.Replacer
var strconvElementarySchool strconv.NumError

var basicElementarySchoolTraitMapping = map[string]func(*schema.ElementarySchoolTrait, []string){}
var customElementarySchoolTraitMapping = map[string]func(*schema.ElementarySchoolTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ElementarySchool", func(ctx jsonld.Context) (interface{}) {
		return NewElementarySchoolFromContext(ctx)
	})

}

func NewElementarySchoolFromContext(ctx jsonld.Context) (x schema.ElementarySchool) {
	x.Type = "http://schema.org/ElementarySchool"
	MapBasicToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)
	MapCustomToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToElementarySchoolTrait(ctx, &x.ElementarySchoolTrait)
	MapCustomToElementarySchoolTrait(ctx, &x.ElementarySchoolTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToElementarySchoolTrait(ctx jsonld.Context, x *schema.ElementarySchoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicElementarySchoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToElementarySchoolTrait(ctx jsonld.Context, x *schema.ElementarySchoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := customElementarySchoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}