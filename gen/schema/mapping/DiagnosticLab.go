package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDiagnosticLab strings.Replacer
var strconvDiagnosticLab strconv.NumError

var basicDiagnosticLabTraitMapping = map[string]func(*schema.DiagnosticLabTrait, []string){}
var customDiagnosticLabTraitMapping = map[string]func(*schema.DiagnosticLabTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DiagnosticLab", func(ctx jsonld.Context) (interface{}) {
		return NewDiagnosticLabFromContext(ctx)
	})



	customDiagnosticLabTraitMapping["http://schema.org/availableTest"] = func(x *schema.DiagnosticLabTrait, ctx jsonld.Context, s string){
		var y schema.MedicalTest
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalTestFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalTest()
			y.Id = s
		}

		x.AvailableTest = &y

		return
	}
}

func NewDiagnosticLabFromContext(ctx jsonld.Context) (x schema.DiagnosticLab) {
	x.Type = "http://schema.org/DiagnosticLab"
	MapBasicToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)
	MapCustomToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDiagnosticLabTrait(ctx, &x.DiagnosticLabTrait)
	MapCustomToDiagnosticLabTrait(ctx, &x.DiagnosticLabTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDiagnosticLabTrait(ctx jsonld.Context, x *schema.DiagnosticLabTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDiagnosticLabTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDiagnosticLabTrait(ctx jsonld.Context, x *schema.DiagnosticLabTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDiagnosticLabTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}