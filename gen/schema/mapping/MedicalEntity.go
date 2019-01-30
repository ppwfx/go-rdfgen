package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalEntity strings.Replacer
var strconvMedicalEntity strconv.NumError

var basicMedicalEntityTraitMapping = map[string]func(*schema.MedicalEntityTrait, []string){}
var customMedicalEntityTraitMapping = map[string]func(*schema.MedicalEntityTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalEntity", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalEntityFromContext(ctx)
	})









	customMedicalEntityTraitMapping["http://schema.org/code"] = func(x *schema.MedicalEntityTrait, ctx jsonld.Context, s string){
		var y schema.MedicalCode
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalCodeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalCode()
			y.Id = s
		}

		x.Code = &y

		return
	}

	customMedicalEntityTraitMapping["http://schema.org/guideline"] = func(x *schema.MedicalEntityTrait, ctx jsonld.Context, s string){
		var y schema.MedicalGuideline
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalGuidelineFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalGuideline()
			y.Id = s
		}

		x.Guideline = &y

		return
	}

	customMedicalEntityTraitMapping["http://schema.org/legalStatus"] = func(x *schema.MedicalEntityTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.LegalStatus, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.LegalStatus = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.LegalStatus = s
		}
	}

	customMedicalEntityTraitMapping["http://schema.org/medicineSystem"] = func(x *schema.MedicalEntityTrait, ctx jsonld.Context, s string){
		var y schema.MedicineSystem
		if strings.HasPrefix(s, "_:") {
			y = NewMedicineSystemFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicineSystem()
			y.Id = s
		}

		x.MedicineSystem = &y

		return
	}

	customMedicalEntityTraitMapping["http://schema.org/recognizingAuthority"] = func(x *schema.MedicalEntityTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.RecognizingAuthority = &y

		return
	}

	customMedicalEntityTraitMapping["http://schema.org/relevantSpecialty"] = func(x *schema.MedicalEntityTrait, ctx jsonld.Context, s string){
		var y schema.MedicalSpecialty
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalSpecialtyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalSpecialty()
			y.Id = s
		}

		x.RelevantSpecialty = &y

		return
	}

	customMedicalEntityTraitMapping["http://schema.org/study"] = func(x *schema.MedicalEntityTrait, ctx jsonld.Context, s string){
		var y schema.MedicalStudy
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalStudyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalStudy()
			y.Id = s
		}

		x.Study = &y

		return
	}
}

func NewMedicalEntityFromContext(ctx jsonld.Context) (x schema.MedicalEntity) {
	x.Type = "http://schema.org/MedicalEntity"
	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalEntityTrait(ctx jsonld.Context, x *schema.MedicalEntityTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalEntityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalEntityTrait(ctx jsonld.Context, x *schema.MedicalEntityTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalEntityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}