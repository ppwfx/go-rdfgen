package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDrug strings.Replacer
var strconvDrug strconv.NumError

var basicDrugTraitMapping = map[string]func(*schema.DrugTrait, []string){}
var customDrugTraitMapping = map[string]func(*schema.DrugTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Drug", func(ctx jsonld.Context) (interface{}) {
		return NewDrugFromContext(ctx)
	})





	basicDrugTraitMapping["http://schema.org/activeIngredient"] = func(x *schema.DrugTrait, s []string) {
		x.ActiveIngredient = s[0]
	}


	basicDrugTraitMapping["http://schema.org/dosageForm"] = func(x *schema.DrugTrait, s []string) {
		x.DosageForm = s[0]
	}


	basicDrugTraitMapping["http://schema.org/clincalPharmacology"] = func(x *schema.DrugTrait, s []string) {
		x.ClincalPharmacology = s[0]
	}


	basicDrugTraitMapping["http://schema.org/clinicalPharmacology"] = func(x *schema.DrugTrait, s []string) {
		x.ClinicalPharmacology = s[0]
	}


	basicDrugTraitMapping["http://schema.org/overdosage"] = func(x *schema.DrugTrait, s []string) {
		x.Overdosage = s[0]
	}


	basicDrugTraitMapping["http://schema.org/pregnancyWarning"] = func(x *schema.DrugTrait, s []string) {
		x.PregnancyWarning = s[0]
	}


	basicDrugTraitMapping["http://schema.org/drugUnit"] = func(x *schema.DrugTrait, s []string) {
		x.DrugUnit = s[0]
	}


	basicDrugTraitMapping["http://schema.org/foodWarning"] = func(x *schema.DrugTrait, s []string) {
		x.FoodWarning = s[0]
	}


	basicDrugTraitMapping["http://schema.org/nonProprietaryName"] = func(x *schema.DrugTrait, s []string) {
		x.NonProprietaryName = s[0]
	}


	basicDrugTraitMapping["http://schema.org/alcoholWarning"] = func(x *schema.DrugTrait, s []string) {
		x.AlcoholWarning = s[0]
	}


	basicDrugTraitMapping["http://schema.org/mechanismOfAction"] = func(x *schema.DrugTrait, s []string) {
		x.MechanismOfAction = s[0]
	}


	basicDrugTraitMapping["http://schema.org/breastfeedingWarning"] = func(x *schema.DrugTrait, s []string) {
		x.BreastfeedingWarning = s[0]
	}


	basicDrugTraitMapping["http://schema.org/administrationRoute"] = func(x *schema.DrugTrait, s []string) {
		x.AdministrationRoute = s[0]
	}


	basicDrugTraitMapping["http://schema.org/rxcui"] = func(x *schema.DrugTrait, s []string) {
		x.Rxcui = s[0]
	}



	basicDrugTraitMapping["http://schema.org/proprietaryName"] = func(x *schema.DrugTrait, s []string) {
		x.ProprietaryName = s[0]
	}


	basicDrugTraitMapping["http://schema.org/prescribingInfo"] = func(x *schema.DrugTrait, s []string) {
		x.PrescribingInfo = s[0]
	}


	basicDrugTraitMapping["http://schema.org/labelDetails"] = func(x *schema.DrugTrait, s []string) {
		x.LabelDetails = s[0]
	}


	basicDrugTraitMapping["http://schema.org/isAvailableGenerically"] = func(x *schema.DrugTrait, s []string) {
		var err error
		x.IsAvailableGenerically, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	basicDrugTraitMapping["http://schema.org/isProprietary"] = func(x *schema.DrugTrait, s []string) {
		var err error
		x.IsProprietary, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}











	customDrugTraitMapping["http://schema.org/manufacturer"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.Manufacturer = &y

		return
	}

	customDrugTraitMapping["http://schema.org/legalStatus"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
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

	customDrugTraitMapping["http://schema.org/prescriptionStatus"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PrescriptionStatus, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PrescriptionStatus = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PrescriptionStatus = s
		}
	}

	customDrugTraitMapping["http://schema.org/warning"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Warning, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Warning = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Warning = s
		}
	}

	customDrugTraitMapping["http://schema.org/maximumIntake"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		var y schema.MaximumDoseSchedule
		if strings.HasPrefix(s, "_:") {
			y = NewMaximumDoseScheduleFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMaximumDoseSchedule()
			y.Id = s
		}

		x.MaximumIntake = &y

		return
	}

	customDrugTraitMapping["http://schema.org/availableStrength"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		var y schema.DrugStrength
		if strings.HasPrefix(s, "_:") {
			y = NewDrugStrengthFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDrugStrength()
			y.Id = s
		}

		x.AvailableStrength = &y

		return
	}

	customDrugTraitMapping["http://schema.org/cost"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		var y schema.DrugCost
		if strings.HasPrefix(s, "_:") {
			y = NewDrugCostFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDrugCost()
			y.Id = s
		}

		x.Cost = &y

		return
	}

	customDrugTraitMapping["http://schema.org/doseSchedule"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		var y schema.DoseSchedule
		if strings.HasPrefix(s, "_:") {
			y = NewDoseScheduleFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDoseSchedule()
			y.Id = s
		}

		x.DoseSchedule = &y

		return
	}

	customDrugTraitMapping["http://schema.org/drugClass"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		var y schema.DrugClass
		if strings.HasPrefix(s, "_:") {
			y = NewDrugClassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDrugClass()
			y.Id = s
		}

		x.DrugClass = &y

		return
	}

	customDrugTraitMapping["http://schema.org/includedInHealthInsurancePlan"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		var y schema.HealthInsurancePlan
		if strings.HasPrefix(s, "_:") {
			y = NewHealthInsurancePlanFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewHealthInsurancePlan()
			y.Id = s
		}

		x.IncludedInHealthInsurancePlan = &y

		return
	}

	customDrugTraitMapping["http://schema.org/interactingDrug"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		var y schema.Drug
		if strings.HasPrefix(s, "_:") {
			y = NewDrugFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDrug()
			y.Id = s
		}

		x.InteractingDrug = &y

		return
	}

	customDrugTraitMapping["http://schema.org/pregnancyCategory"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		var y schema.DrugPregnancyCategory
		if strings.HasPrefix(s, "_:") {
			y = NewDrugPregnancyCategoryFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDrugPregnancyCategory()
			y.Id = s
		}

		x.PregnancyCategory = &y

		return
	}

	customDrugTraitMapping["http://schema.org/relatedDrug"] = func(x *schema.DrugTrait, ctx jsonld.Context, s string){
		var y schema.Drug
		if strings.HasPrefix(s, "_:") {
			y = NewDrugFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDrug()
			y.Id = s
		}

		x.RelatedDrug = &y

		return
	}
}

func NewDrugFromContext(ctx jsonld.Context) (x schema.Drug) {
	x.Type = "http://schema.org/Drug"
	MapBasicToSubstanceTrait(ctx, &x.SubstanceTrait)
	MapCustomToSubstanceTrait(ctx, &x.SubstanceTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDrugTrait(ctx, &x.DrugTrait)
	MapCustomToDrugTrait(ctx, &x.DrugTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDrugTrait(ctx jsonld.Context, x *schema.DrugTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDrugTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDrugTrait(ctx jsonld.Context, x *schema.DrugTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDrugTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}