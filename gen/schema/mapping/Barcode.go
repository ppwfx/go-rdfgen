package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBarcode strings.Replacer
var strconvBarcode strconv.NumError

var basicBarcodeTraitMapping = map[string]func(*schema.BarcodeTrait, []string){}
var customBarcodeTraitMapping = map[string]func(*schema.BarcodeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Barcode", func(ctx jsonld.Context) (interface{}) {
		return NewBarcodeFromContext(ctx)
	})

}

func NewBarcodeFromContext(ctx jsonld.Context) (x schema.Barcode) {
	x.Type = "http://schema.org/Barcode"
	MapBasicToImageObjectTrait(ctx, &x.ImageObjectTrait)
	MapCustomToImageObjectTrait(ctx, &x.ImageObjectTrait)

	MapBasicToMediaObjectTrait(ctx, &x.MediaObjectTrait)
	MapCustomToMediaObjectTrait(ctx, &x.MediaObjectTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBarcodeTrait(ctx, &x.BarcodeTrait)
	MapCustomToBarcodeTrait(ctx, &x.BarcodeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBarcodeTrait(ctx jsonld.Context, x *schema.BarcodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBarcodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBarcodeTrait(ctx jsonld.Context, x *schema.BarcodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBarcodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}