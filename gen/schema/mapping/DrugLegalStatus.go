package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDrugLegalStatus strings.Replacer
var strconvDrugLegalStatus strconv.NumError

var basicDrugLegalStatusTraitMapping = map[string]func(*schema.DrugLegalStatusTrait, []string){}
var customDrugLegalStatusTraitMapping = map[string]func(*schema.DrugLegalStatusTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DrugLegalStatus", func(ctx jsonld.Context) (interface{}) {
		return NewDrugLegalStatusFromContext(ctx)
	})



	customDrugLegalStatusTraitMapping["http://schema.org/applicableLocation"] = func(x *schema.DrugLegalStatusTrait, ctx jsonld.Context, s string){
		var y schema.AdministrativeArea
		if strings.HasPrefix(s, "_:") {
			y = NewAdministrativeAreaFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAdministrativeArea()
			y.Id = s
		}

		x.ApplicableLocation = &y

		return
	}
}

func NewDrugLegalStatusFromContext(ctx jsonld.Context) (x schema.DrugLegalStatus) {
	x.Type = "http://schema.org/DrugLegalStatus"
	MapBasicToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)
	MapCustomToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDrugLegalStatusTrait(ctx, &x.DrugLegalStatusTrait)
	MapCustomToDrugLegalStatusTrait(ctx, &x.DrugLegalStatusTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDrugLegalStatusTrait(ctx jsonld.Context, x *schema.DrugLegalStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDrugLegalStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDrugLegalStatusTrait(ctx jsonld.Context, x *schema.DrugLegalStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDrugLegalStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}