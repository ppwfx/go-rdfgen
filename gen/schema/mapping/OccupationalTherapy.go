package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOccupationalTherapy strings.Replacer
var strconvOccupationalTherapy strconv.NumError

var basicOccupationalTherapyTraitMapping = map[string]func(*schema.OccupationalTherapyTrait, []string){}
var customOccupationalTherapyTraitMapping = map[string]func(*schema.OccupationalTherapyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OccupationalTherapy", func(ctx jsonld.Context) (interface{}) {
		return NewOccupationalTherapyFromContext(ctx)
	})

}

func NewOccupationalTherapyFromContext(ctx jsonld.Context) (x schema.OccupationalTherapy) {
	x.Type = "http://schema.org/OccupationalTherapy"
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


	MapBasicToOccupationalTherapyTrait(ctx, &x.OccupationalTherapyTrait)
	MapCustomToOccupationalTherapyTrait(ctx, &x.OccupationalTherapyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOccupationalTherapyTrait(ctx jsonld.Context, x *schema.OccupationalTherapyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOccupationalTherapyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOccupationalTherapyTrait(ctx jsonld.Context, x *schema.OccupationalTherapyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOccupationalTherapyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}