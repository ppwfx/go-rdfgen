package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUseAction strings.Replacer
var strconvUseAction strconv.NumError

var basicUseActionTraitMapping = map[string]func(*schema.UseActionTrait, []string){}
var customUseActionTraitMapping = map[string]func(*schema.UseActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UseAction", func(ctx jsonld.Context) (interface{}) {
		return NewUseActionFromContext(ctx)
	})

}

func NewUseActionFromContext(ctx jsonld.Context) (x schema.UseAction) {
	x.Type = "http://schema.org/UseAction"
	MapBasicToConsumeActionTrait(ctx, &x.ConsumeActionTrait)
	MapCustomToConsumeActionTrait(ctx, &x.ConsumeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUseActionTrait(ctx, &x.UseActionTrait)
	MapCustomToUseActionTrait(ctx, &x.UseActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUseActionTrait(ctx jsonld.Context, x *schema.UseActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUseActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUseActionTrait(ctx jsonld.Context, x *schema.UseActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUseActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}