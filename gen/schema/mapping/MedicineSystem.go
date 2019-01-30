package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicineSystem strings.Replacer
var strconvMedicineSystem strconv.NumError

var basicMedicineSystemTraitMapping = map[string]func(*schema.MedicineSystemTrait, []string){}
var customMedicineSystemTraitMapping = map[string]func(*schema.MedicineSystemTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicineSystem", func(ctx jsonld.Context) (interface{}) {
		return NewMedicineSystemFromContext(ctx)
	})

}

func NewMedicineSystemFromContext(ctx jsonld.Context) (x schema.MedicineSystem) {
	x.Type = "http://schema.org/MedicineSystem"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicineSystemTrait(ctx, &x.MedicineSystemTrait)
	MapCustomToMedicineSystemTrait(ctx, &x.MedicineSystemTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicineSystemTrait(ctx jsonld.Context, x *schema.MedicineSystemTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicineSystemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicineSystemTrait(ctx jsonld.Context, x *schema.MedicineSystemTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicineSystemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}