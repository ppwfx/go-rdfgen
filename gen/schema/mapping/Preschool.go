package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPreschool strings.Replacer
var strconvPreschool strconv.NumError

var basicPreschoolTraitMapping = map[string]func(*schema.PreschoolTrait, []string){}
var customPreschoolTraitMapping = map[string]func(*schema.PreschoolTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Preschool", func(ctx jsonld.Context) (interface{}) {
		return NewPreschoolFromContext(ctx)
	})

}

func NewPreschoolFromContext(ctx jsonld.Context) (x schema.Preschool) {
	x.Type = "http://schema.org/Preschool"
	MapBasicToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)
	MapCustomToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPreschoolTrait(ctx, &x.PreschoolTrait)
	MapCustomToPreschoolTrait(ctx, &x.PreschoolTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPreschoolTrait(ctx jsonld.Context, x *schema.PreschoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPreschoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPreschoolTrait(ctx jsonld.Context, x *schema.PreschoolTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPreschoolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}