package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLifestyleModification strings.Replacer
var strconvLifestyleModification strconv.NumError

var basicLifestyleModificationTraitMapping = map[string]func(*schema.LifestyleModificationTrait, []string){}
var customLifestyleModificationTraitMapping = map[string]func(*schema.LifestyleModificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LifestyleModification", func(ctx jsonld.Context) (interface{}) {
		return NewLifestyleModificationFromContext(ctx)
	})

}

func NewLifestyleModificationFromContext(ctx jsonld.Context) (x schema.LifestyleModification) {
	x.Type = "http://schema.org/LifestyleModification"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLifestyleModificationTrait(ctx, &x.LifestyleModificationTrait)
	MapCustomToLifestyleModificationTrait(ctx, &x.LifestyleModificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLifestyleModificationTrait(ctx jsonld.Context, x *schema.LifestyleModificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLifestyleModificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLifestyleModificationTrait(ctx jsonld.Context, x *schema.LifestyleModificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLifestyleModificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}