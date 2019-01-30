package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDiagnosticProcedure strings.Replacer
var strconvDiagnosticProcedure strconv.NumError

var basicDiagnosticProcedureTraitMapping = map[string]func(*schema.DiagnosticProcedureTrait, []string){}
var customDiagnosticProcedureTraitMapping = map[string]func(*schema.DiagnosticProcedureTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DiagnosticProcedure", func(ctx jsonld.Context) (interface{}) {
		return NewDiagnosticProcedureFromContext(ctx)
	})

}

func NewDiagnosticProcedureFromContext(ctx jsonld.Context) (x schema.DiagnosticProcedure) {
	x.Type = "http://schema.org/DiagnosticProcedure"
	MapBasicToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)
	MapCustomToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDiagnosticProcedureTrait(ctx, &x.DiagnosticProcedureTrait)
	MapCustomToDiagnosticProcedureTrait(ctx, &x.DiagnosticProcedureTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDiagnosticProcedureTrait(ctx jsonld.Context, x *schema.DiagnosticProcedureTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDiagnosticProcedureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDiagnosticProcedureTrait(ctx jsonld.Context, x *schema.DiagnosticProcedureTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDiagnosticProcedureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}