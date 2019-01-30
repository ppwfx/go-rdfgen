package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalContraindication strings.Replacer
var strconvMedicalContraindication strconv.NumError

var basicMedicalContraindicationTraitMapping = map[string]func(*schema.MedicalContraindicationTrait, []string){}
var customMedicalContraindicationTraitMapping = map[string]func(*schema.MedicalContraindicationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalContraindication", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalContraindicationFromContext(ctx)
	})

}

func NewMedicalContraindicationFromContext(ctx jsonld.Context) (x schema.MedicalContraindication) {
	x.Type = "http://schema.org/MedicalContraindication"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalContraindicationTrait(ctx, &x.MedicalContraindicationTrait)
	MapCustomToMedicalContraindicationTrait(ctx, &x.MedicalContraindicationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalContraindicationTrait(ctx jsonld.Context, x *schema.MedicalContraindicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalContraindicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalContraindicationTrait(ctx jsonld.Context, x *schema.MedicalContraindicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalContraindicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}