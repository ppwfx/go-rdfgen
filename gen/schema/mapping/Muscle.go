package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMuscle strings.Replacer
var strconvMuscle strconv.NumError

var basicMuscleTraitMapping = map[string]func(*schema.MuscleTrait, []string){}
var customMuscleTraitMapping = map[string]func(*schema.MuscleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Muscle", func(ctx jsonld.Context) (interface{}) {
		return NewMuscleFromContext(ctx)
	})


	basicMuscleTraitMapping["http://schema.org/action"] = func(x *schema.MuscleTrait, s []string) {
		x.Action = s[0]
	}


	basicMuscleTraitMapping["http://schema.org/muscleAction"] = func(x *schema.MuscleTrait, s []string) {
		x.MuscleAction = s[0]
	}







	customMuscleTraitMapping["http://schema.org/antagonist"] = func(x *schema.MuscleTrait, ctx jsonld.Context, s string){
		var y schema.Muscle
		if strings.HasPrefix(s, "_:") {
			y = NewMuscleFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMuscle()
			y.Id = s
		}

		x.Antagonist = &y

		return
	}

	customMuscleTraitMapping["http://schema.org/bloodSupply"] = func(x *schema.MuscleTrait, ctx jsonld.Context, s string){
		var y schema.Vessel
		if strings.HasPrefix(s, "_:") {
			y = NewVesselFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewVessel()
			y.Id = s
		}

		x.BloodSupply = &y

		return
	}

	customMuscleTraitMapping["http://schema.org/insertion"] = func(x *schema.MuscleTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalStructure
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalStructure()
			y.Id = s
		}

		x.Insertion = &y

		return
	}

	customMuscleTraitMapping["http://schema.org/nerve"] = func(x *schema.MuscleTrait, ctx jsonld.Context, s string){
		var y schema.Nerve
		if strings.HasPrefix(s, "_:") {
			y = NewNerveFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewNerve()
			y.Id = s
		}

		x.Nerve = &y

		return
	}

	customMuscleTraitMapping["http://schema.org/origin"] = func(x *schema.MuscleTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalStructure
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalStructure()
			y.Id = s
		}

		x.Origin = &y

		return
	}
}

func NewMuscleFromContext(ctx jsonld.Context) (x schema.Muscle) {
	x.Type = "http://schema.org/Muscle"
	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMuscleTrait(ctx, &x.MuscleTrait)
	MapCustomToMuscleTrait(ctx, &x.MuscleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMuscleTrait(ctx jsonld.Context, x *schema.MuscleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMuscleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMuscleTrait(ctx jsonld.Context, x *schema.MuscleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMuscleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}