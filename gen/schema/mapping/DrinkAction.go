package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDrinkAction strings.Replacer
var strconvDrinkAction strconv.NumError

var basicDrinkActionTraitMapping = map[string]func(*schema.DrinkActionTrait, []string){}
var customDrinkActionTraitMapping = map[string]func(*schema.DrinkActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DrinkAction", func(ctx jsonld.Context) (interface{}) {
		return NewDrinkActionFromContext(ctx)
	})

}

func NewDrinkActionFromContext(ctx jsonld.Context) (x schema.DrinkAction) {
	x.Type = "http://schema.org/DrinkAction"
	MapBasicToConsumeActionTrait(ctx, &x.ConsumeActionTrait)
	MapCustomToConsumeActionTrait(ctx, &x.ConsumeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDrinkActionTrait(ctx, &x.DrinkActionTrait)
	MapCustomToDrinkActionTrait(ctx, &x.DrinkActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDrinkActionTrait(ctx jsonld.Context, x *schema.DrinkActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDrinkActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDrinkActionTrait(ctx jsonld.Context, x *schema.DrinkActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDrinkActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}