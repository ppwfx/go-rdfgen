package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPresentationDigitalDocument strings.Replacer
var strconvPresentationDigitalDocument strconv.NumError

var basicPresentationDigitalDocumentTraitMapping = map[string]func(*schema.PresentationDigitalDocumentTrait, []string){}
var customPresentationDigitalDocumentTraitMapping = map[string]func(*schema.PresentationDigitalDocumentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PresentationDigitalDocument", func(ctx jsonld.Context) (interface{}) {
		return NewPresentationDigitalDocumentFromContext(ctx)
	})

}

func NewPresentationDigitalDocumentFromContext(ctx jsonld.Context) (x schema.PresentationDigitalDocument) {
	x.Type = "http://schema.org/PresentationDigitalDocument"
	MapBasicToDigitalDocumentTrait(ctx, &x.DigitalDocumentTrait)
	MapCustomToDigitalDocumentTrait(ctx, &x.DigitalDocumentTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPresentationDigitalDocumentTrait(ctx, &x.PresentationDigitalDocumentTrait)
	MapCustomToPresentationDigitalDocumentTrait(ctx, &x.PresentationDigitalDocumentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPresentationDigitalDocumentTrait(ctx jsonld.Context, x *schema.PresentationDigitalDocumentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPresentationDigitalDocumentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPresentationDigitalDocumentTrait(ctx jsonld.Context, x *schema.PresentationDigitalDocumentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPresentationDigitalDocumentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}