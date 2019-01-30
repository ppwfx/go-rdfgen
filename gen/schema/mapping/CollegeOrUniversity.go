package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCollegeOrUniversity strings.Replacer
var strconvCollegeOrUniversity strconv.NumError

var basicCollegeOrUniversityTraitMapping = map[string]func(*schema.CollegeOrUniversityTrait, []string){}
var customCollegeOrUniversityTraitMapping = map[string]func(*schema.CollegeOrUniversityTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CollegeOrUniversity", func(ctx jsonld.Context) (interface{}) {
		return NewCollegeOrUniversityFromContext(ctx)
	})

}

func NewCollegeOrUniversityFromContext(ctx jsonld.Context) (x schema.CollegeOrUniversity) {
	x.Type = "http://schema.org/CollegeOrUniversity"
	MapBasicToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)
	MapCustomToEducationalOrganizationTrait(ctx, &x.EducationalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCollegeOrUniversityTrait(ctx, &x.CollegeOrUniversityTrait)
	MapCustomToCollegeOrUniversityTrait(ctx, &x.CollegeOrUniversityTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCollegeOrUniversityTrait(ctx jsonld.Context, x *schema.CollegeOrUniversityTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCollegeOrUniversityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCollegeOrUniversityTrait(ctx jsonld.Context, x *schema.CollegeOrUniversityTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCollegeOrUniversityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}