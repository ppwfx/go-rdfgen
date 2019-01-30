package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGolfCourse strings.Replacer
var strconvGolfCourse strconv.NumError

var basicGolfCourseTraitMapping = map[string]func(*schema.GolfCourseTrait, []string){}
var customGolfCourseTraitMapping = map[string]func(*schema.GolfCourseTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GolfCourse", func(ctx jsonld.Context) (interface{}) {
		return NewGolfCourseFromContext(ctx)
	})

}

func NewGolfCourseFromContext(ctx jsonld.Context) (x schema.GolfCourse) {
	x.Type = "http://schema.org/GolfCourse"
	MapBasicToSportsActivityLocationTrait(ctx, &x.SportsActivityLocationTrait)
	MapCustomToSportsActivityLocationTrait(ctx, &x.SportsActivityLocationTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToGolfCourseTrait(ctx, &x.GolfCourseTrait)
	MapCustomToGolfCourseTrait(ctx, &x.GolfCourseTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGolfCourseTrait(ctx jsonld.Context, x *schema.GolfCourseTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGolfCourseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGolfCourseTrait(ctx jsonld.Context, x *schema.GolfCourseTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGolfCourseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}