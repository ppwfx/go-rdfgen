package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsExerciseAction strings.Replacer
var strconvExerciseAction strconv.NumError

var basicExerciseActionTraitMapping = map[string]func(*schema.ExerciseActionTrait, []string){}
var customExerciseActionTraitMapping = map[string]func(*schema.ExerciseActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ExerciseAction", func(ctx jsonld.Context) (interface{}) {
		return NewExerciseActionFromContext(ctx)
	})


	basicExerciseActionTraitMapping["http://schema.org/exerciseType"] = func(x *schema.ExerciseActionTrait, s []string) {
		x.ExerciseType = s[0]
	}














	customExerciseActionTraitMapping["http://schema.org/fromLocation"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.FromLocation = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/toLocation"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.ToLocation = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/exerciseCourse"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.ExerciseCourse = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/course"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.Course = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/opponent"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Opponent = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/diet"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.Diet
		if strings.HasPrefix(s, "_:") {
			y = NewDietFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDiet()
			y.Id = s
		}

		x.Diet = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/distance"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.Distance
		if strings.HasPrefix(s, "_:") {
			y = NewDistanceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDistance()
			y.Id = s
		}

		x.Distance = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/exercisePlan"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.ExercisePlan
		if strings.HasPrefix(s, "_:") {
			y = NewExercisePlanFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewExercisePlan()
			y.Id = s
		}

		x.ExercisePlan = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/exerciseRelatedDiet"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.Diet
		if strings.HasPrefix(s, "_:") {
			y = NewDietFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDiet()
			y.Id = s
		}

		x.ExerciseRelatedDiet = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/sportsActivityLocation"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.SportsActivityLocation
		if strings.HasPrefix(s, "_:") {
			y = NewSportsActivityLocationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewSportsActivityLocation()
			y.Id = s
		}

		x.SportsActivityLocation = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/sportsEvent"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.SportsEvent
		if strings.HasPrefix(s, "_:") {
			y = NewSportsEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewSportsEvent()
			y.Id = s
		}

		x.SportsEvent = &y

		return
	}

	customExerciseActionTraitMapping["http://schema.org/sportsTeam"] = func(x *schema.ExerciseActionTrait, ctx jsonld.Context, s string){
		var y schema.SportsTeam
		if strings.HasPrefix(s, "_:") {
			y = NewSportsTeamFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewSportsTeam()
			y.Id = s
		}

		x.SportsTeam = &y

		return
	}
}

func NewExerciseActionFromContext(ctx jsonld.Context) (x schema.ExerciseAction) {
	x.Type = "http://schema.org/ExerciseAction"
	MapBasicToPlayActionTrait(ctx, &x.PlayActionTrait)
	MapCustomToPlayActionTrait(ctx, &x.PlayActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToExerciseActionTrait(ctx, &x.ExerciseActionTrait)
	MapCustomToExerciseActionTrait(ctx, &x.ExerciseActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToExerciseActionTrait(ctx jsonld.Context, x *schema.ExerciseActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicExerciseActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToExerciseActionTrait(ctx jsonld.Context, x *schema.ExerciseActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customExerciseActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}