package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsExercisePlan strings.Replacer
var strconvExercisePlan strconv.NumError

var basicExercisePlanTraitMapping = map[string]func(*schema.ExercisePlanTrait, []string){}
var customExercisePlanTraitMapping = map[string]func(*schema.ExercisePlanTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ExercisePlan", func(ctx jsonld.Context) (interface{}) {
		return NewExercisePlanFromContext(ctx)
	})




	basicExercisePlanTraitMapping["http://schema.org/exerciseType"] = func(x *schema.ExercisePlanTrait, s []string) {
		x.ExerciseType = s[0]
	}



	basicExercisePlanTraitMapping["http://schema.org/additionalVariable"] = func(x *schema.ExercisePlanTrait, s []string) {
		x.AdditionalVariable = s[0]
	}





	customExercisePlanTraitMapping["http://schema.org/restPeriods"] = func(x *schema.ExercisePlanTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RestPeriods, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RestPeriods = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RestPeriods = s
		}
	}

	customExercisePlanTraitMapping["http://schema.org/activityFrequency"] = func(x *schema.ExercisePlanTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ActivityFrequency, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ActivityFrequency = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ActivityFrequency = s
		}
	}

	customExercisePlanTraitMapping["http://schema.org/intensity"] = func(x *schema.ExercisePlanTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Intensity, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Intensity = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Intensity = s
		}
	}

	customExercisePlanTraitMapping["http://schema.org/activityDuration"] = func(x *schema.ExercisePlanTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ActivityDuration, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ActivityDuration = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ActivityDuration = s
		}
	}

	customExercisePlanTraitMapping["http://schema.org/repetitions"] = func(x *schema.ExercisePlanTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Repetitions, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Repetitions = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Repetitions = s
		}
	}

	customExercisePlanTraitMapping["http://schema.org/workload"] = func(x *schema.ExercisePlanTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Workload, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Workload = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Workload = s
		}
	}
}

func NewExercisePlanFromContext(ctx jsonld.Context) (x schema.ExercisePlan) {
	x.Type = "http://schema.org/ExercisePlan"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToPhysicalActivityTrait(ctx, &x.PhysicalActivityTrait)
	MapCustomToPhysicalActivityTrait(ctx, &x.PhysicalActivityTrait)

	MapBasicToLifestyleModificationTrait(ctx, &x.LifestyleModificationTrait)
	MapCustomToLifestyleModificationTrait(ctx, &x.LifestyleModificationTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)


	MapBasicToExercisePlanTrait(ctx, &x.ExercisePlanTrait)
	MapCustomToExercisePlanTrait(ctx, &x.ExercisePlanTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToExercisePlanTrait(ctx jsonld.Context, x *schema.ExercisePlanTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicExercisePlanTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToExercisePlanTrait(ctx jsonld.Context, x *schema.ExercisePlanTrait) () {
	for k, v := range ctx.Current {
		f, ok := customExercisePlanTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}