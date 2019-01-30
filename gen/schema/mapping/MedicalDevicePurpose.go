package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalDevicePurpose strings.Replacer
var strconvMedicalDevicePurpose strconv.NumError

var basicMedicalDevicePurposeTraitMapping = map[string]func(*schema.MedicalDevicePurposeTrait, []string){}
var customMedicalDevicePurposeTraitMapping = map[string]func(*schema.MedicalDevicePurposeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalDevicePurpose", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalDevicePurposeFromContext(ctx)
	})

}

func NewMedicalDevicePurposeFromContext(ctx jsonld.Context) (x schema.MedicalDevicePurpose) {
	x.Type = "http://schema.org/MedicalDevicePurpose"
	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalDevicePurposeTrait(ctx, &x.MedicalDevicePurposeTrait)
	MapCustomToMedicalDevicePurposeTrait(ctx, &x.MedicalDevicePurposeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalDevicePurposeTrait(ctx jsonld.Context, x *schema.MedicalDevicePurposeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalDevicePurposeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalDevicePurposeTrait(ctx jsonld.Context, x *schema.MedicalDevicePurposeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalDevicePurposeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}