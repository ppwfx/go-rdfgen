package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPalliativeProcedure strings.Replacer
var strconvPalliativeProcedure strconv.NumError

var basicPalliativeProcedureTraitMapping = map[string]func(*schema.PalliativeProcedureTrait, []string){}
var customPalliativeProcedureTraitMapping = map[string]func(*schema.PalliativeProcedureTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PalliativeProcedure", func(ctx jsonld.Context) (interface{}) {
		return NewPalliativeProcedureFromContext(ctx)
	})

}

func NewPalliativeProcedureFromContext(ctx jsonld.Context) (x schema.PalliativeProcedure) {
	x.Type = "http://schema.org/PalliativeProcedure"
	MapBasicToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)
	MapCustomToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToMedicalTherapyTrait(ctx, &x.MedicalTherapyTrait)
	MapCustomToMedicalTherapyTrait(ctx, &x.MedicalTherapyTrait)

	MapBasicToTherapeuticProcedureTrait(ctx, &x.TherapeuticProcedureTrait)
	MapCustomToTherapeuticProcedureTrait(ctx, &x.TherapeuticProcedureTrait)


	MapBasicToPalliativeProcedureTrait(ctx, &x.PalliativeProcedureTrait)
	MapCustomToPalliativeProcedureTrait(ctx, &x.PalliativeProcedureTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPalliativeProcedureTrait(ctx jsonld.Context, x *schema.PalliativeProcedureTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPalliativeProcedureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPalliativeProcedureTrait(ctx jsonld.Context, x *schema.PalliativeProcedureTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPalliativeProcedureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}