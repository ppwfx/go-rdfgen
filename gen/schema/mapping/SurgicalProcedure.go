package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSurgicalProcedure strings.Replacer
var strconvSurgicalProcedure strconv.NumError

var basicSurgicalProcedureTraitMapping = map[string]func(*schema.SurgicalProcedureTrait, []string){}
var customSurgicalProcedureTraitMapping = map[string]func(*schema.SurgicalProcedureTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SurgicalProcedure", func(ctx jsonld.Context) (interface{}) {
		return NewSurgicalProcedureFromContext(ctx)
	})

}

func NewSurgicalProcedureFromContext(ctx jsonld.Context) (x schema.SurgicalProcedure) {
	x.Type = "http://schema.org/SurgicalProcedure"
	MapBasicToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)
	MapCustomToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSurgicalProcedureTrait(ctx, &x.SurgicalProcedureTrait)
	MapCustomToSurgicalProcedureTrait(ctx, &x.SurgicalProcedureTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSurgicalProcedureTrait(ctx jsonld.Context, x *schema.SurgicalProcedureTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSurgicalProcedureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSurgicalProcedureTrait(ctx jsonld.Context, x *schema.SurgicalProcedureTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSurgicalProcedureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}