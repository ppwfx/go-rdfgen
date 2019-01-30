package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRadiationTherapy strings.Replacer
var strconvRadiationTherapy strconv.NumError

var basicRadiationTherapyTraitMapping = map[string]func(*schema.RadiationTherapyTrait, []string){}
var customRadiationTherapyTraitMapping = map[string]func(*schema.RadiationTherapyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RadiationTherapy", func(ctx jsonld.Context) (interface{}) {
		return NewRadiationTherapyFromContext(ctx)
	})

}

func NewRadiationTherapyFromContext(ctx jsonld.Context) (x schema.RadiationTherapy) {
	x.Type = "http://schema.org/RadiationTherapy"
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


	MapBasicToRadiationTherapyTrait(ctx, &x.RadiationTherapyTrait)
	MapCustomToRadiationTherapyTrait(ctx, &x.RadiationTherapyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRadiationTherapyTrait(ctx jsonld.Context, x *schema.RadiationTherapyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRadiationTherapyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRadiationTherapyTrait(ctx jsonld.Context, x *schema.RadiationTherapyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRadiationTherapyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}