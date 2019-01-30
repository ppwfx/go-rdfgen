package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNoteDigitalDocument strings.Replacer
var strconvNoteDigitalDocument strconv.NumError

var basicNoteDigitalDocumentTraitMapping = map[string]func(*schema.NoteDigitalDocumentTrait, []string){}
var customNoteDigitalDocumentTraitMapping = map[string]func(*schema.NoteDigitalDocumentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/NoteDigitalDocument", func(ctx jsonld.Context) (interface{}) {
		return NewNoteDigitalDocumentFromContext(ctx)
	})

}

func NewNoteDigitalDocumentFromContext(ctx jsonld.Context) (x schema.NoteDigitalDocument) {
	x.Type = "http://schema.org/NoteDigitalDocument"
	MapBasicToDigitalDocumentTrait(ctx, &x.DigitalDocumentTrait)
	MapCustomToDigitalDocumentTrait(ctx, &x.DigitalDocumentTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToNoteDigitalDocumentTrait(ctx, &x.NoteDigitalDocumentTrait)
	MapCustomToNoteDigitalDocumentTrait(ctx, &x.NoteDigitalDocumentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNoteDigitalDocumentTrait(ctx jsonld.Context, x *schema.NoteDigitalDocumentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNoteDigitalDocumentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNoteDigitalDocumentTrait(ctx jsonld.Context, x *schema.NoteDigitalDocumentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNoteDigitalDocumentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}