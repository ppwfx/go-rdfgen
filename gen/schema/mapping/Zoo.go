package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsZoo strings.Replacer
var strconvZoo strconv.NumError

var basicZooTraitMapping = map[string]func(*schema.ZooTrait, []string){}
var customZooTraitMapping = map[string]func(*schema.ZooTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Zoo", func(ctx jsonld.Context) (interface{}) {
		return NewZooFromContext(ctx)
	})

}

func NewZooFromContext(ctx jsonld.Context) (x schema.Zoo) {
	x.Type = "http://schema.org/Zoo"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToZooTrait(ctx, &x.ZooTrait)
	MapCustomToZooTrait(ctx, &x.ZooTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToZooTrait(ctx jsonld.Context, x *schema.ZooTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicZooTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToZooTrait(ctx jsonld.Context, x *schema.ZooTrait) () {
	for k, v := range ctx.Current {
		f, ok := customZooTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}