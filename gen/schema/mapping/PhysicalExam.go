package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPhysicalExam strings.Replacer
var strconvPhysicalExam strconv.NumError

var basicPhysicalExamTraitMapping = map[string]func(*schema.PhysicalExamTrait, []string){}
var customPhysicalExamTraitMapping = map[string]func(*schema.PhysicalExamTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PhysicalExam", func(ctx jsonld.Context) (interface{}) {
		return NewPhysicalExamFromContext(ctx)
	})

}

func NewPhysicalExamFromContext(ctx jsonld.Context) (x schema.PhysicalExam) {
	x.Type = "http://schema.org/PhysicalExam"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)
	MapCustomToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)


	MapBasicToPhysicalExamTrait(ctx, &x.PhysicalExamTrait)
	MapCustomToPhysicalExamTrait(ctx, &x.PhysicalExamTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPhysicalExamTrait(ctx jsonld.Context, x *schema.PhysicalExamTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPhysicalExamTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPhysicalExamTrait(ctx jsonld.Context, x *schema.PhysicalExamTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPhysicalExamTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}