package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalCode strings.Replacer
var strconvMedicalCode strconv.NumError

var basicMedicalCodeTraitMapping = map[string]func(*schema.MedicalCodeTrait, []string){}
var customMedicalCodeTraitMapping = map[string]func(*schema.MedicalCodeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalCode", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalCodeFromContext(ctx)
	})


	basicMedicalCodeTraitMapping["http://schema.org/codingSystem"] = func(x *schema.MedicalCodeTrait, s []string) {
		x.CodingSystem = s[0]
	}


	basicMedicalCodeTraitMapping["http://schema.org/codeValue"] = func(x *schema.MedicalCodeTrait, s []string) {
		x.CodeValue = s[0]
	}

}

func NewMedicalCodeFromContext(ctx jsonld.Context) (x schema.MedicalCode) {
	x.Type = "http://schema.org/MedicalCode"
	MapBasicToCategoryCodeTrait(ctx, &x.CategoryCodeTrait)
	MapCustomToCategoryCodeTrait(ctx, &x.CategoryCodeTrait)

	MapBasicToDefinedTermTrait(ctx, &x.DefinedTermTrait)
	MapCustomToDefinedTermTrait(ctx, &x.DefinedTermTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)
	MapCustomToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)


	MapBasicToMedicalCodeTrait(ctx, &x.MedicalCodeTrait)
	MapCustomToMedicalCodeTrait(ctx, &x.MedicalCodeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalCodeTrait(ctx jsonld.Context, x *schema.MedicalCodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalCodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalCodeTrait(ctx jsonld.Context, x *schema.MedicalCodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalCodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}