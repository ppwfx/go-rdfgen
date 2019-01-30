package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalIndication strings.Replacer
var strconvMedicalIndication strconv.NumError

var basicMedicalIndicationTraitMapping = map[string]func(*schema.MedicalIndicationTrait, []string){}
var customMedicalIndicationTraitMapping = map[string]func(*schema.MedicalIndicationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalIndication", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalIndicationFromContext(ctx)
	})

}

func NewMedicalIndicationFromContext(ctx jsonld.Context) (x schema.MedicalIndication) {
	x.Type = "http://schema.org/MedicalIndication"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalIndicationTrait(ctx, &x.MedicalIndicationTrait)
	MapCustomToMedicalIndicationTrait(ctx, &x.MedicalIndicationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalIndicationTrait(ctx jsonld.Context, x *schema.MedicalIndicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalIndicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalIndicationTrait(ctx jsonld.Context, x *schema.MedicalIndicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalIndicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}