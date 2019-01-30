package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSpreadsheetDigitalDocument strings.Replacer
var strconvSpreadsheetDigitalDocument strconv.NumError

var basicSpreadsheetDigitalDocumentTraitMapping = map[string]func(*schema.SpreadsheetDigitalDocumentTrait, []string){}
var customSpreadsheetDigitalDocumentTraitMapping = map[string]func(*schema.SpreadsheetDigitalDocumentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SpreadsheetDigitalDocument", func(ctx jsonld.Context) (interface{}) {
		return NewSpreadsheetDigitalDocumentFromContext(ctx)
	})

}

func NewSpreadsheetDigitalDocumentFromContext(ctx jsonld.Context) (x schema.SpreadsheetDigitalDocument) {
	x.Type = "http://schema.org/SpreadsheetDigitalDocument"
	MapBasicToDigitalDocumentTrait(ctx, &x.DigitalDocumentTrait)
	MapCustomToDigitalDocumentTrait(ctx, &x.DigitalDocumentTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSpreadsheetDigitalDocumentTrait(ctx, &x.SpreadsheetDigitalDocumentTrait)
	MapCustomToSpreadsheetDigitalDocumentTrait(ctx, &x.SpreadsheetDigitalDocumentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSpreadsheetDigitalDocumentTrait(ctx jsonld.Context, x *schema.SpreadsheetDigitalDocumentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSpreadsheetDigitalDocumentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSpreadsheetDigitalDocumentTrait(ctx jsonld.Context, x *schema.SpreadsheetDigitalDocumentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSpreadsheetDigitalDocumentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}