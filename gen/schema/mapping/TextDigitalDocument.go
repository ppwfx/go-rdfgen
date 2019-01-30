package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTextDigitalDocument strings.Replacer
var strconvTextDigitalDocument strconv.NumError

var basicTextDigitalDocumentTraitMapping = map[string]func(*schema.TextDigitalDocumentTrait, []string){}
var customTextDigitalDocumentTraitMapping = map[string]func(*schema.TextDigitalDocumentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TextDigitalDocument", func(ctx jsonld.Context) (interface{}) {
		return NewTextDigitalDocumentFromContext(ctx)
	})

}

func NewTextDigitalDocumentFromContext(ctx jsonld.Context) (x schema.TextDigitalDocument) {
	x.Type = "http://schema.org/TextDigitalDocument"
	MapBasicToDigitalDocumentTrait(ctx, &x.DigitalDocumentTrait)
	MapCustomToDigitalDocumentTrait(ctx, &x.DigitalDocumentTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTextDigitalDocumentTrait(ctx, &x.TextDigitalDocumentTrait)
	MapCustomToTextDigitalDocumentTrait(ctx, &x.TextDigitalDocumentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTextDigitalDocumentTrait(ctx jsonld.Context, x *schema.TextDigitalDocumentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTextDigitalDocumentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTextDigitalDocumentTrait(ctx jsonld.Context, x *schema.TextDigitalDocumentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTextDigitalDocumentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}