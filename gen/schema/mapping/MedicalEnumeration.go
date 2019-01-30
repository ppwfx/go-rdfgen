package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalEnumeration strings.Replacer
var strconvMedicalEnumeration strconv.NumError

var basicMedicalEnumerationTraitMapping = map[string]func(*schema.MedicalEnumerationTrait, []string){}
var customMedicalEnumerationTraitMapping = map[string]func(*schema.MedicalEnumerationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalEnumeration", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalEnumerationFromContext(ctx)
	})

}

func NewMedicalEnumerationFromContext(ctx jsonld.Context) (x schema.MedicalEnumeration) {
	x.Type = "http://schema.org/MedicalEnumeration"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalEnumerationTrait(ctx jsonld.Context, x *schema.MedicalEnumerationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalEnumerationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalEnumerationTrait(ctx jsonld.Context, x *schema.MedicalEnumerationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalEnumerationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}