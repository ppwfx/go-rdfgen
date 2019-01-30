package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDrugCostCategory strings.Replacer
var strconvDrugCostCategory strconv.NumError

var basicDrugCostCategoryTraitMapping = map[string]func(*schema.DrugCostCategoryTrait, []string){}
var customDrugCostCategoryTraitMapping = map[string]func(*schema.DrugCostCategoryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DrugCostCategory", func(ctx jsonld.Context) (interface{}) {
		return NewDrugCostCategoryFromContext(ctx)
	})

}

func NewDrugCostCategoryFromContext(ctx jsonld.Context) (x schema.DrugCostCategory) {
	x.Type = "http://schema.org/DrugCostCategory"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDrugCostCategoryTrait(ctx, &x.DrugCostCategoryTrait)
	MapCustomToDrugCostCategoryTrait(ctx, &x.DrugCostCategoryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDrugCostCategoryTrait(ctx jsonld.Context, x *schema.DrugCostCategoryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDrugCostCategoryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDrugCostCategoryTrait(ctx jsonld.Context, x *schema.DrugCostCategoryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDrugCostCategoryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}