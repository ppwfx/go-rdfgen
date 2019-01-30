package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDrugPregnancyCategory strings.Replacer
var strconvDrugPregnancyCategory strconv.NumError

var basicDrugPregnancyCategoryTraitMapping = map[string]func(*schema.DrugPregnancyCategoryTrait, []string){}
var customDrugPregnancyCategoryTraitMapping = map[string]func(*schema.DrugPregnancyCategoryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DrugPregnancyCategory", func(ctx jsonld.Context) (interface{}) {
		return NewDrugPregnancyCategoryFromContext(ctx)
	})

}

func NewDrugPregnancyCategoryFromContext(ctx jsonld.Context) (x schema.DrugPregnancyCategory) {
	x.Type = "http://schema.org/DrugPregnancyCategory"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDrugPregnancyCategoryTrait(ctx, &x.DrugPregnancyCategoryTrait)
	MapCustomToDrugPregnancyCategoryTrait(ctx, &x.DrugPregnancyCategoryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDrugPregnancyCategoryTrait(ctx jsonld.Context, x *schema.DrugPregnancyCategoryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDrugPregnancyCategoryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDrugPregnancyCategoryTrait(ctx jsonld.Context, x *schema.DrugPregnancyCategoryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDrugPregnancyCategoryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}