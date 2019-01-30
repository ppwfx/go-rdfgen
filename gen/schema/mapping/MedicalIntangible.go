package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalIntangible strings.Replacer
var strconvMedicalIntangible strconv.NumError

var basicMedicalIntangibleTraitMapping = map[string]func(*schema.MedicalIntangibleTrait, []string){}
var customMedicalIntangibleTraitMapping = map[string]func(*schema.MedicalIntangibleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalIntangible", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalIntangibleFromContext(ctx)
	})

}

func NewMedicalIntangibleFromContext(ctx jsonld.Context) (x schema.MedicalIntangible) {
	x.Type = "http://schema.org/MedicalIntangible"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)
	MapCustomToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalIntangibleTrait(ctx jsonld.Context, x *schema.MedicalIntangibleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalIntangibleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalIntangibleTrait(ctx jsonld.Context, x *schema.MedicalIntangibleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalIntangibleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}