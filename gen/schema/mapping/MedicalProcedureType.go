package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalProcedureType strings.Replacer
var strconvMedicalProcedureType strconv.NumError

var basicMedicalProcedureTypeTraitMapping = map[string]func(*schema.MedicalProcedureTypeTrait, []string){}
var customMedicalProcedureTypeTraitMapping = map[string]func(*schema.MedicalProcedureTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalProcedureType", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalProcedureTypeFromContext(ctx)
	})

}

func NewMedicalProcedureTypeFromContext(ctx jsonld.Context) (x schema.MedicalProcedureType) {
	x.Type = "http://schema.org/MedicalProcedureType"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalProcedureTypeTrait(ctx, &x.MedicalProcedureTypeTrait)
	MapCustomToMedicalProcedureTypeTrait(ctx, &x.MedicalProcedureTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalProcedureTypeTrait(ctx jsonld.Context, x *schema.MedicalProcedureTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalProcedureTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalProcedureTypeTrait(ctx jsonld.Context, x *schema.MedicalProcedureTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalProcedureTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}