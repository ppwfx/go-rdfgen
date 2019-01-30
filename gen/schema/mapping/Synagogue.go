package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSynagogue strings.Replacer
var strconvSynagogue strconv.NumError

var basicSynagogueTraitMapping = map[string]func(*schema.SynagogueTrait, []string){}
var customSynagogueTraitMapping = map[string]func(*schema.SynagogueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Synagogue", func(ctx jsonld.Context) (interface{}) {
		return NewSynagogueFromContext(ctx)
	})

}

func NewSynagogueFromContext(ctx jsonld.Context) (x schema.Synagogue) {
	x.Type = "http://schema.org/Synagogue"
	MapBasicToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)
	MapCustomToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSynagogueTrait(ctx, &x.SynagogueTrait)
	MapCustomToSynagogueTrait(ctx, &x.SynagogueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSynagogueTrait(ctx jsonld.Context, x *schema.SynagogueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSynagogueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSynagogueTrait(ctx jsonld.Context, x *schema.SynagogueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSynagogueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}