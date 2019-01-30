package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNerve strings.Replacer
var strconvNerve strconv.NumError

var basicNerveTraitMapping = map[string]func(*schema.NerveTrait, []string){}
var customNerveTraitMapping = map[string]func(*schema.NerveTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Nerve", func(ctx jsonld.Context) (interface{}) {
		return NewNerveFromContext(ctx)
	})






	customNerveTraitMapping["http://schema.org/branch"] = func(x *schema.NerveTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalStructure
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalStructure()
			y.Id = s
		}

		x.Branch = &y

		return
	}

	customNerveTraitMapping["http://schema.org/nerveMotor"] = func(x *schema.NerveTrait, ctx jsonld.Context, s string){
		var y schema.Muscle
		if strings.HasPrefix(s, "_:") {
			y = NewMuscleFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMuscle()
			y.Id = s
		}

		x.NerveMotor = &y

		return
	}

	customNerveTraitMapping["http://schema.org/sensoryUnit"] = func(x *schema.NerveTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SensoryUnit, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SensoryUnit = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SensoryUnit = s
		}
	}

	customNerveTraitMapping["http://schema.org/sourcedFrom"] = func(x *schema.NerveTrait, ctx jsonld.Context, s string){
		var y schema.BrainStructure
		if strings.HasPrefix(s, "_:") {
			y = NewBrainStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBrainStructure()
			y.Id = s
		}

		x.SourcedFrom = &y

		return
	}
}

func NewNerveFromContext(ctx jsonld.Context) (x schema.Nerve) {
	x.Type = "http://schema.org/Nerve"
	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToNerveTrait(ctx, &x.NerveTrait)
	MapCustomToNerveTrait(ctx, &x.NerveTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNerveTrait(ctx jsonld.Context, x *schema.NerveTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNerveTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNerveTrait(ctx jsonld.Context, x *schema.NerveTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNerveTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}