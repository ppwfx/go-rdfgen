package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAddAction strings.Replacer
var strconvAddAction strconv.NumError

var basicAddActionTraitMapping = map[string]func(*schema.AddActionTrait, []string){}
var customAddActionTraitMapping = map[string]func(*schema.AddActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AddAction", func(ctx jsonld.Context) (interface{}) {
		return NewAddActionFromContext(ctx)
	})

}

func NewAddActionFromContext(ctx jsonld.Context) (x schema.AddAction) {
	x.Type = "http://schema.org/AddAction"
	MapBasicToUpdateActionTrait(ctx, &x.UpdateActionTrait)
	MapCustomToUpdateActionTrait(ctx, &x.UpdateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAddActionTrait(ctx, &x.AddActionTrait)
	MapCustomToAddActionTrait(ctx, &x.AddActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAddActionTrait(ctx jsonld.Context, x *schema.AddActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAddActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAddActionTrait(ctx jsonld.Context, x *schema.AddActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAddActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}