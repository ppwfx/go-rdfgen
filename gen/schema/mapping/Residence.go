package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsResidence strings.Replacer
var strconvResidence strconv.NumError

var basicResidenceTraitMapping = map[string]func(*schema.ResidenceTrait, []string){}
var customResidenceTraitMapping = map[string]func(*schema.ResidenceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Residence", func(ctx jsonld.Context) (interface{}) {
		return NewResidenceFromContext(ctx)
	})

}

func NewResidenceFromContext(ctx jsonld.Context) (x schema.Residence) {
	x.Type = "http://schema.org/Residence"
	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToResidenceTrait(ctx, &x.ResidenceTrait)
	MapCustomToResidenceTrait(ctx, &x.ResidenceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToResidenceTrait(ctx jsonld.Context, x *schema.ResidenceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicResidenceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToResidenceTrait(ctx jsonld.Context, x *schema.ResidenceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customResidenceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}