package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPaintAction strings.Replacer
var strconvPaintAction strconv.NumError

var basicPaintActionTraitMapping = map[string]func(*schema.PaintActionTrait, []string){}
var customPaintActionTraitMapping = map[string]func(*schema.PaintActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PaintAction", func(ctx jsonld.Context) (interface{}) {
		return NewPaintActionFromContext(ctx)
	})

}

func NewPaintActionFromContext(ctx jsonld.Context) (x schema.PaintAction) {
	x.Type = "http://schema.org/PaintAction"
	MapBasicToCreateActionTrait(ctx, &x.CreateActionTrait)
	MapCustomToCreateActionTrait(ctx, &x.CreateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPaintActionTrait(ctx, &x.PaintActionTrait)
	MapCustomToPaintActionTrait(ctx, &x.PaintActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPaintActionTrait(ctx jsonld.Context, x *schema.PaintActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPaintActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPaintActionTrait(ctx jsonld.Context, x *schema.PaintActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPaintActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}