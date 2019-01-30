package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVitalSign strings.Replacer
var strconvVitalSign strconv.NumError

var basicVitalSignTraitMapping = map[string]func(*schema.VitalSignTrait, []string){}
var customVitalSignTraitMapping = map[string]func(*schema.VitalSignTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/VitalSign", func(ctx jsonld.Context) (interface{}) {
		return NewVitalSignFromContext(ctx)
	})

}

func NewVitalSignFromContext(ctx jsonld.Context) (x schema.VitalSign) {
	x.Type = "http://schema.org/VitalSign"
	MapBasicToMedicalSignTrait(ctx, &x.MedicalSignTrait)
	MapCustomToMedicalSignTrait(ctx, &x.MedicalSignTrait)

	MapBasicToMedicalSignOrSymptomTrait(ctx, &x.MedicalSignOrSymptomTrait)
	MapCustomToMedicalSignOrSymptomTrait(ctx, &x.MedicalSignOrSymptomTrait)

	MapBasicToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)
	MapCustomToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVitalSignTrait(ctx, &x.VitalSignTrait)
	MapCustomToVitalSignTrait(ctx, &x.VitalSignTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVitalSignTrait(ctx jsonld.Context, x *schema.VitalSignTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVitalSignTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVitalSignTrait(ctx jsonld.Context, x *schema.VitalSignTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVitalSignTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}