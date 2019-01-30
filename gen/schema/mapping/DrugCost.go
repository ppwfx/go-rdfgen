package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDrugCost strings.Replacer
var strconvDrugCost strconv.NumError

var basicDrugCostTraitMapping = map[string]func(*schema.DrugCostTrait, []string){}
var customDrugCostTraitMapping = map[string]func(*schema.DrugCostTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DrugCost", func(ctx jsonld.Context) (interface{}) {
		return NewDrugCostFromContext(ctx)
	})



	basicDrugCostTraitMapping["http://schema.org/drugUnit"] = func(x *schema.DrugCostTrait, s []string) {
		x.DrugUnit = s[0]
	}


	basicDrugCostTraitMapping["http://schema.org/costCurrency"] = func(x *schema.DrugCostTrait, s []string) {
		x.CostCurrency = s[0]
	}


	basicDrugCostTraitMapping["http://schema.org/costOrigin"] = func(x *schema.DrugCostTrait, s []string) {
		x.CostOrigin = s[0]
	}




	customDrugCostTraitMapping["http://schema.org/costPerUnit"] = func(x *schema.DrugCostTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.CostPerUnit, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.CostPerUnit = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.CostPerUnit = s
		}
	}

	customDrugCostTraitMapping["http://schema.org/applicableLocation"] = func(x *schema.DrugCostTrait, ctx jsonld.Context, s string){
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

	customDrugCostTraitMapping["http://schema.org/costCategory"] = func(x *schema.DrugCostTrait, ctx jsonld.Context, s string){
		var y schema.DrugCostCategory
		if strings.HasPrefix(s, "_:") {
			y = NewDrugCostCategoryFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDrugCostCategory()
			y.Id = s
		}

		x.CostCategory = &y

		return
	}
}

func NewDrugCostFromContext(ctx jsonld.Context) (x schema.DrugCost) {
	x.Type = "http://schema.org/DrugCost"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDrugCostTrait(ctx, &x.DrugCostTrait)
	MapCustomToDrugCostTrait(ctx, &x.DrugCostTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDrugCostTrait(ctx jsonld.Context, x *schema.DrugCostTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDrugCostTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDrugCostTrait(ctx jsonld.Context, x *schema.DrugCostTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDrugCostTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}