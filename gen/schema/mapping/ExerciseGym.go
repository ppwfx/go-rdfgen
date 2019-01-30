package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsExerciseGym strings.Replacer
var strconvExerciseGym strconv.NumError

var basicExerciseGymTraitMapping = map[string]func(*schema.ExerciseGymTrait, []string){}
var customExerciseGymTraitMapping = map[string]func(*schema.ExerciseGymTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ExerciseGym", func(ctx jsonld.Context) (interface{}) {
		return NewExerciseGymFromContext(ctx)
	})

}

func NewExerciseGymFromContext(ctx jsonld.Context) (x schema.ExerciseGym) {
	x.Type = "http://schema.org/ExerciseGym"
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


	MapBasicToExerciseGymTrait(ctx, &x.ExerciseGymTrait)
	MapCustomToExerciseGymTrait(ctx, &x.ExerciseGymTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToExerciseGymTrait(ctx jsonld.Context, x *schema.ExerciseGymTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicExerciseGymTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToExerciseGymTrait(ctx jsonld.Context, x *schema.ExerciseGymTrait) () {
	for k, v := range ctx.Current {
		f, ok := customExerciseGymTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}