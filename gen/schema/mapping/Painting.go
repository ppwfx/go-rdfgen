package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPainting strings.Replacer
var strconvPainting strconv.NumError

var basicPaintingTraitMapping = map[string]func(*schema.PaintingTrait, []string){}
var customPaintingTraitMapping = map[string]func(*schema.PaintingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Painting", func(ctx jsonld.Context) (interface{}) {
		return NewPaintingFromContext(ctx)
	})

}

func NewPaintingFromContext(ctx jsonld.Context) (x schema.Painting) {
	x.Type = "http://schema.org/Painting"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPaintingTrait(ctx, &x.PaintingTrait)
	MapCustomToPaintingTrait(ctx, &x.PaintingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPaintingTrait(ctx jsonld.Context, x *schema.PaintingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPaintingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPaintingTrait(ctx jsonld.Context, x *schema.PaintingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPaintingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}