package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTreatmentIndication strings.Replacer
var strconvTreatmentIndication strconv.NumError

var basicTreatmentIndicationTraitMapping = map[string]func(*schema.TreatmentIndicationTrait, []string){}
var customTreatmentIndicationTraitMapping = map[string]func(*schema.TreatmentIndicationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TreatmentIndication", func(ctx jsonld.Context) (interface{}) {
		return NewTreatmentIndicationFromContext(ctx)
	})

}

func NewTreatmentIndicationFromContext(ctx jsonld.Context) (x schema.TreatmentIndication) {
	x.Type = "http://schema.org/TreatmentIndication"
	MapBasicToMedicalIndicationTrait(ctx, &x.MedicalIndicationTrait)
	MapCustomToMedicalIndicationTrait(ctx, &x.MedicalIndicationTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTreatmentIndicationTrait(ctx, &x.TreatmentIndicationTrait)
	MapCustomToTreatmentIndicationTrait(ctx, &x.TreatmentIndicationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTreatmentIndicationTrait(ctx jsonld.Context, x *schema.TreatmentIndicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTreatmentIndicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTreatmentIndicationTrait(ctx jsonld.Context, x *schema.TreatmentIndicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTreatmentIndicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}