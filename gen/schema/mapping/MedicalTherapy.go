package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalTherapy strings.Replacer
var strconvMedicalTherapy strconv.NumError

var basicMedicalTherapyTraitMapping = map[string]func(*schema.MedicalTherapyTrait, []string){}
var customMedicalTherapyTraitMapping = map[string]func(*schema.MedicalTherapyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalTherapy", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalTherapyFromContext(ctx)
	})





	customMedicalTherapyTraitMapping["http://schema.org/contraindication"] = func(x *schema.MedicalTherapyTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Contraindication, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Contraindication = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Contraindication = s
		}
	}

	customMedicalTherapyTraitMapping["http://schema.org/seriousAdverseOutcome"] = func(x *schema.MedicalTherapyTrait, ctx jsonld.Context, s string){
		var y schema.MedicalEntity
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalEntityFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalEntity()
			y.Id = s
		}

		x.SeriousAdverseOutcome = &y

		return
	}

	customMedicalTherapyTraitMapping["http://schema.org/duplicateTherapy"] = func(x *schema.MedicalTherapyTrait, ctx jsonld.Context, s string){
		var y schema.MedicalTherapy
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalTherapyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalTherapy()
			y.Id = s
		}

		x.DuplicateTherapy = &y

		return
	}
}

func NewMedicalTherapyFromContext(ctx jsonld.Context) (x schema.MedicalTherapy) {
	x.Type = "http://schema.org/MedicalTherapy"
	MapBasicToTherapeuticProcedureTrait(ctx, &x.TherapeuticProcedureTrait)
	MapCustomToTherapeuticProcedureTrait(ctx, &x.TherapeuticProcedureTrait)

	MapBasicToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)
	MapCustomToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalTherapyTrait(ctx, &x.MedicalTherapyTrait)
	MapCustomToMedicalTherapyTrait(ctx, &x.MedicalTherapyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalTherapyTrait(ctx jsonld.Context, x *schema.MedicalTherapyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalTherapyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalTherapyTrait(ctx jsonld.Context, x *schema.MedicalTherapyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalTherapyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}