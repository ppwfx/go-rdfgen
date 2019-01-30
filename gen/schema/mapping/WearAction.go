package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWearAction strings.Replacer
var strconvWearAction strconv.NumError

var basicWearActionTraitMapping = map[string]func(*schema.WearActionTrait, []string){}
var customWearActionTraitMapping = map[string]func(*schema.WearActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WearAction", func(ctx jsonld.Context) (interface{}) {
		return NewWearActionFromContext(ctx)
	})

}

func NewWearActionFromContext(ctx jsonld.Context) (x schema.WearAction) {
	x.Type = "http://schema.org/WearAction"
	MapBasicToUseActionTrait(ctx, &x.UseActionTrait)
	MapCustomToUseActionTrait(ctx, &x.UseActionTrait)

	MapBasicToConsumeActionTrait(ctx, &x.ConsumeActionTrait)
	MapCustomToConsumeActionTrait(ctx, &x.ConsumeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWearActionTrait(ctx, &x.WearActionTrait)
	MapCustomToWearActionTrait(ctx, &x.WearActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWearActionTrait(ctx jsonld.Context, x *schema.WearActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWearActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWearActionTrait(ctx jsonld.Context, x *schema.WearActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWearActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}