package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFindAction strings.Replacer
var strconvFindAction strconv.NumError

var basicFindActionTraitMapping = map[string]func(*schema.FindActionTrait, []string){}
var customFindActionTraitMapping = map[string]func(*schema.FindActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FindAction", func(ctx jsonld.Context) (interface{}) {
		return NewFindActionFromContext(ctx)
	})

}

func NewFindActionFromContext(ctx jsonld.Context) (x schema.FindAction) {
	x.Type = "http://schema.org/FindAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFindActionTrait(ctx, &x.FindActionTrait)
	MapCustomToFindActionTrait(ctx, &x.FindActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFindActionTrait(ctx jsonld.Context, x *schema.FindActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFindActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFindActionTrait(ctx jsonld.Context, x *schema.FindActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFindActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}