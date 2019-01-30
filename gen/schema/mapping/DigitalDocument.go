package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDigitalDocument strings.Replacer
var strconvDigitalDocument strconv.NumError

var basicDigitalDocumentTraitMapping = map[string]func(*schema.DigitalDocumentTrait, []string){}
var customDigitalDocumentTraitMapping = map[string]func(*schema.DigitalDocumentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DigitalDocument", func(ctx jsonld.Context) (interface{}) {
		return NewDigitalDocumentFromContext(ctx)
	})



	customDigitalDocumentTraitMapping["http://schema.org/hasDigitalDocumentPermission"] = func(x *schema.DigitalDocumentTrait, ctx jsonld.Context, s string){
		var y schema.DigitalDocumentPermission
		if strings.HasPrefix(s, "_:") {
			y = NewDigitalDocumentPermissionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDigitalDocumentPermission()
			y.Id = s
		}

		x.HasDigitalDocumentPermission = &y

		return
	}
}

func NewDigitalDocumentFromContext(ctx jsonld.Context) (x schema.DigitalDocument) {
	x.Type = "http://schema.org/DigitalDocument"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDigitalDocumentTrait(ctx, &x.DigitalDocumentTrait)
	MapCustomToDigitalDocumentTrait(ctx, &x.DigitalDocumentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDigitalDocumentTrait(ctx jsonld.Context, x *schema.DigitalDocumentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDigitalDocumentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDigitalDocumentTrait(ctx jsonld.Context, x *schema.DigitalDocumentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDigitalDocumentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}