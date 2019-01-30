package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSpecialty strings.Replacer
var strconvSpecialty strconv.NumError

var basicSpecialtyTraitMapping = map[string]func(*schema.SpecialtyTrait, []string){}
var customSpecialtyTraitMapping = map[string]func(*schema.SpecialtyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Specialty", func(ctx jsonld.Context) (interface{}) {
		return NewSpecialtyFromContext(ctx)
	})

}

func NewSpecialtyFromContext(ctx jsonld.Context) (x schema.Specialty) {
	x.Type = "http://schema.org/Specialty"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSpecialtyTrait(ctx, &x.SpecialtyTrait)
	MapCustomToSpecialtyTrait(ctx, &x.SpecialtyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSpecialtyTrait(ctx jsonld.Context, x *schema.SpecialtyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSpecialtyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSpecialtyTrait(ctx jsonld.Context, x *schema.SpecialtyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSpecialtyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}