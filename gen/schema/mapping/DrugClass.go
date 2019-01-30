package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDrugClass strings.Replacer
var strconvDrugClass strconv.NumError

var basicDrugClassTraitMapping = map[string]func(*schema.DrugClassTrait, []string){}
var customDrugClassTraitMapping = map[string]func(*schema.DrugClassTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DrugClass", func(ctx jsonld.Context) (interface{}) {
		return NewDrugClassFromContext(ctx)
	})



	customDrugClassTraitMapping["http://schema.org/drug"] = func(x *schema.DrugClassTrait, ctx jsonld.Context, s string){
		var y schema.Drug
		if strings.HasPrefix(s, "_:") {
			y = NewDrugFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDrug()
			y.Id = s
		}

		x.Drug = &y

		return
	}
}

func NewDrugClassFromContext(ctx jsonld.Context) (x schema.DrugClass) {
	x.Type = "http://schema.org/DrugClass"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDrugClassTrait(ctx, &x.DrugClassTrait)
	MapCustomToDrugClassTrait(ctx, &x.DrugClassTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDrugClassTrait(ctx jsonld.Context, x *schema.DrugClassTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDrugClassTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDrugClassTrait(ctx jsonld.Context, x *schema.DrugClassTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDrugClassTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}