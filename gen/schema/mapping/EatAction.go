package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEatAction strings.Replacer
var strconvEatAction strconv.NumError

var basicEatActionTraitMapping = map[string]func(*schema.EatActionTrait, []string){}
var customEatActionTraitMapping = map[string]func(*schema.EatActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EatAction", func(ctx jsonld.Context) (interface{}) {
		return NewEatActionFromContext(ctx)
	})

}

func NewEatActionFromContext(ctx jsonld.Context) (x schema.EatAction) {
	x.Type = "http://schema.org/EatAction"
	MapBasicToConsumeActionTrait(ctx, &x.ConsumeActionTrait)
	MapCustomToConsumeActionTrait(ctx, &x.ConsumeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEatActionTrait(ctx, &x.EatActionTrait)
	MapCustomToEatActionTrait(ctx, &x.EatActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEatActionTrait(ctx jsonld.Context, x *schema.EatActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEatActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEatActionTrait(ctx jsonld.Context, x *schema.EatActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEatActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}