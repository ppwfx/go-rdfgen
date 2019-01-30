package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReadAction strings.Replacer
var strconvReadAction strconv.NumError

var basicReadActionTraitMapping = map[string]func(*schema.ReadActionTrait, []string){}
var customReadActionTraitMapping = map[string]func(*schema.ReadActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReadAction", func(ctx jsonld.Context) (interface{}) {
		return NewReadActionFromContext(ctx)
	})

}

func NewReadActionFromContext(ctx jsonld.Context) (x schema.ReadAction) {
	x.Type = "http://schema.org/ReadAction"
	MapBasicToConsumeActionTrait(ctx, &x.ConsumeActionTrait)
	MapCustomToConsumeActionTrait(ctx, &x.ConsumeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReadActionTrait(ctx, &x.ReadActionTrait)
	MapCustomToReadActionTrait(ctx, &x.ReadActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReadActionTrait(ctx jsonld.Context, x *schema.ReadActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReadActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReadActionTrait(ctx jsonld.Context, x *schema.ReadActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReadActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}