package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsState strings.Replacer
var strconvState strconv.NumError

var basicStateTraitMapping = map[string]func(*schema.StateTrait, []string){}
var customStateTraitMapping = map[string]func(*schema.StateTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/State", func(ctx jsonld.Context) (interface{}) {
		return NewStateFromContext(ctx)
	})

}

func NewStateFromContext(ctx jsonld.Context) (x schema.State) {
	x.Type = "http://schema.org/State"
	MapBasicToAdministrativeAreaTrait(ctx, &x.AdministrativeAreaTrait)
	MapCustomToAdministrativeAreaTrait(ctx, &x.AdministrativeAreaTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToStateTrait(ctx, &x.StateTrait)
	MapCustomToStateTrait(ctx, &x.StateTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToStateTrait(ctx jsonld.Context, x *schema.StateTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicStateTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToStateTrait(ctx jsonld.Context, x *schema.StateTrait) () {
	for k, v := range ctx.Current {
		f, ok := customStateTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}