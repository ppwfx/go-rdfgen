package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDrugPrescriptionStatus strings.Replacer
var strconvDrugPrescriptionStatus strconv.NumError

var basicDrugPrescriptionStatusTraitMapping = map[string]func(*schema.DrugPrescriptionStatusTrait, []string){}
var customDrugPrescriptionStatusTraitMapping = map[string]func(*schema.DrugPrescriptionStatusTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DrugPrescriptionStatus", func(ctx jsonld.Context) (interface{}) {
		return NewDrugPrescriptionStatusFromContext(ctx)
	})

}

func NewDrugPrescriptionStatusFromContext(ctx jsonld.Context) (x schema.DrugPrescriptionStatus) {
	x.Type = "http://schema.org/DrugPrescriptionStatus"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDrugPrescriptionStatusTrait(ctx, &x.DrugPrescriptionStatusTrait)
	MapCustomToDrugPrescriptionStatusTrait(ctx, &x.DrugPrescriptionStatusTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDrugPrescriptionStatusTrait(ctx jsonld.Context, x *schema.DrugPrescriptionStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDrugPrescriptionStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDrugPrescriptionStatusTrait(ctx jsonld.Context, x *schema.DrugPrescriptionStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDrugPrescriptionStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}