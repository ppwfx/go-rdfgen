package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVolcano strings.Replacer
var strconvVolcano strconv.NumError

var basicVolcanoTraitMapping = map[string]func(*schema.VolcanoTrait, []string){}
var customVolcanoTraitMapping = map[string]func(*schema.VolcanoTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Volcano", func(ctx jsonld.Context) (interface{}) {
		return NewVolcanoFromContext(ctx)
	})

}

func NewVolcanoFromContext(ctx jsonld.Context) (x schema.Volcano) {
	x.Type = "http://schema.org/Volcano"
	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVolcanoTrait(ctx, &x.VolcanoTrait)
	MapCustomToVolcanoTrait(ctx, &x.VolcanoTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVolcanoTrait(ctx jsonld.Context, x *schema.VolcanoTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVolcanoTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVolcanoTrait(ctx jsonld.Context, x *schema.VolcanoTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVolcanoTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}