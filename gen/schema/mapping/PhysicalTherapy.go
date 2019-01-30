package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPhysicalTherapy strings.Replacer
var strconvPhysicalTherapy strconv.NumError

var basicPhysicalTherapyTraitMapping = map[string]func(*schema.PhysicalTherapyTrait, []string){}
var customPhysicalTherapyTraitMapping = map[string]func(*schema.PhysicalTherapyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PhysicalTherapy", func(ctx jsonld.Context) (interface{}) {
		return NewPhysicalTherapyFromContext(ctx)
	})

}

func NewPhysicalTherapyFromContext(ctx jsonld.Context) (x schema.PhysicalTherapy) {
	x.Type = "http://schema.org/PhysicalTherapy"
	MapBasicToMedicalTherapyTrait(ctx, &x.MedicalTherapyTrait)
	MapCustomToMedicalTherapyTrait(ctx, &x.MedicalTherapyTrait)

	MapBasicToTherapeuticProcedureTrait(ctx, &x.TherapeuticProcedureTrait)
	MapCustomToTherapeuticProcedureTrait(ctx, &x.TherapeuticProcedureTrait)

	MapBasicToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)
	MapCustomToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPhysicalTherapyTrait(ctx, &x.PhysicalTherapyTrait)
	MapCustomToPhysicalTherapyTrait(ctx, &x.PhysicalTherapyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPhysicalTherapyTrait(ctx jsonld.Context, x *schema.PhysicalTherapyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPhysicalTherapyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPhysicalTherapyTrait(ctx jsonld.Context, x *schema.PhysicalTherapyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPhysicalTherapyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}