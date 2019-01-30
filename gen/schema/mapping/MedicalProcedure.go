package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalProcedure strings.Replacer
var strconvMedicalProcedure strconv.NumError

var basicMedicalProcedureTraitMapping = map[string]func(*schema.MedicalProcedureTrait, []string){}
var customMedicalProcedureTraitMapping = map[string]func(*schema.MedicalProcedureTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalProcedure", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalProcedureFromContext(ctx)
	})




	basicMedicalProcedureTraitMapping["http://schema.org/followup"] = func(x *schema.MedicalProcedureTrait, s []string) {
		x.Followup = s[0]
	}


	basicMedicalProcedureTraitMapping["http://schema.org/howPerformed"] = func(x *schema.MedicalProcedureTrait, s []string) {
		x.HowPerformed = s[0]
	}



	basicMedicalProcedureTraitMapping["http://schema.org/bodyLocation"] = func(x *schema.MedicalProcedureTrait, s []string) {
		x.BodyLocation = s[0]
	}




	customMedicalProcedureTraitMapping["http://schema.org/outcome"] = func(x *schema.MedicalProcedureTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Outcome, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Outcome = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Outcome = s
		}
	}

	customMedicalProcedureTraitMapping["http://schema.org/status"] = func(x *schema.MedicalProcedureTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Status, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Status = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Status = s
		}
	}

	customMedicalProcedureTraitMapping["http://schema.org/preparation"] = func(x *schema.MedicalProcedureTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Preparation, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Preparation = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Preparation = s
		}
	}

	customMedicalProcedureTraitMapping["http://schema.org/indication"] = func(x *schema.MedicalProcedureTrait, ctx jsonld.Context, s string){
		var y schema.MedicalIndication
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalIndicationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalIndication()
			y.Id = s
		}

		x.Indication = &y

		return
	}

	customMedicalProcedureTraitMapping["http://schema.org/procedureType"] = func(x *schema.MedicalProcedureTrait, ctx jsonld.Context, s string){
		var y schema.MedicalProcedureType
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalProcedureTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalProcedureType()
			y.Id = s
		}

		x.ProcedureType = &y

		return
	}
}

func NewMedicalProcedureFromContext(ctx jsonld.Context) (x schema.MedicalProcedure) {
	x.Type = "http://schema.org/MedicalProcedure"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)
	MapCustomToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalProcedureTrait(ctx jsonld.Context, x *schema.MedicalProcedureTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalProcedureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalProcedureTrait(ctx jsonld.Context, x *schema.MedicalProcedureTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalProcedureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}