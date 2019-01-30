package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsApprovedIndication strings.Replacer
var strconvApprovedIndication strconv.NumError

var basicApprovedIndicationTraitMapping = map[string]func(*schema.ApprovedIndicationTrait, []string){}
var customApprovedIndicationTraitMapping = map[string]func(*schema.ApprovedIndicationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ApprovedIndication", func(ctx jsonld.Context) (interface{}) {
		return NewApprovedIndicationFromContext(ctx)
	})

}

func NewApprovedIndicationFromContext(ctx jsonld.Context) (x schema.ApprovedIndication) {
	x.Type = "http://schema.org/ApprovedIndication"
	MapBasicToMedicalIndicationTrait(ctx, &x.MedicalIndicationTrait)
	MapCustomToMedicalIndicationTrait(ctx, &x.MedicalIndicationTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToApprovedIndicationTrait(ctx, &x.ApprovedIndicationTrait)
	MapCustomToApprovedIndicationTrait(ctx, &x.ApprovedIndicationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToApprovedIndicationTrait(ctx jsonld.Context, x *schema.ApprovedIndicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicApprovedIndicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToApprovedIndicationTrait(ctx jsonld.Context, x *schema.ApprovedIndicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customApprovedIndicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}