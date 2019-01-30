package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBloodTest strings.Replacer
var strconvBloodTest strconv.NumError

var basicBloodTestTraitMapping = map[string]func(*schema.BloodTestTrait, []string){}
var customBloodTestTraitMapping = map[string]func(*schema.BloodTestTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BloodTest", func(ctx jsonld.Context) (interface{}) {
		return NewBloodTestFromContext(ctx)
	})

}

func NewBloodTestFromContext(ctx jsonld.Context) (x schema.BloodTest) {
	x.Type = "http://schema.org/BloodTest"
	MapBasicToMedicalTestTrait(ctx, &x.MedicalTestTrait)
	MapCustomToMedicalTestTrait(ctx, &x.MedicalTestTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBloodTestTrait(ctx, &x.BloodTestTrait)
	MapCustomToBloodTestTrait(ctx, &x.BloodTestTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBloodTestTrait(ctx jsonld.Context, x *schema.BloodTestTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBloodTestTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBloodTestTrait(ctx jsonld.Context, x *schema.BloodTestTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBloodTestTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}