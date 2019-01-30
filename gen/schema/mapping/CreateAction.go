package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCreateAction strings.Replacer
var strconvCreateAction strconv.NumError

var basicCreateActionTraitMapping = map[string]func(*schema.CreateActionTrait, []string){}
var customCreateActionTraitMapping = map[string]func(*schema.CreateActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CreateAction", func(ctx jsonld.Context) (interface{}) {
		return NewCreateActionFromContext(ctx)
	})

}

func NewCreateActionFromContext(ctx jsonld.Context) (x schema.CreateAction) {
	x.Type = "http://schema.org/CreateAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCreateActionTrait(ctx, &x.CreateActionTrait)
	MapCustomToCreateActionTrait(ctx, &x.CreateActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCreateActionTrait(ctx jsonld.Context, x *schema.CreateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCreateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCreateActionTrait(ctx jsonld.Context, x *schema.CreateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCreateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}