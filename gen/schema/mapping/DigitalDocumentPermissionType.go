package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDigitalDocumentPermissionType strings.Replacer
var strconvDigitalDocumentPermissionType strconv.NumError

var basicDigitalDocumentPermissionTypeTraitMapping = map[string]func(*schema.DigitalDocumentPermissionTypeTrait, []string){}
var customDigitalDocumentPermissionTypeTraitMapping = map[string]func(*schema.DigitalDocumentPermissionTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DigitalDocumentPermissionType", func(ctx jsonld.Context) (interface{}) {
		return NewDigitalDocumentPermissionTypeFromContext(ctx)
	})

}

func NewDigitalDocumentPermissionTypeFromContext(ctx jsonld.Context) (x schema.DigitalDocumentPermissionType) {
	x.Type = "http://schema.org/DigitalDocumentPermissionType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDigitalDocumentPermissionTypeTrait(ctx, &x.DigitalDocumentPermissionTypeTrait)
	MapCustomToDigitalDocumentPermissionTypeTrait(ctx, &x.DigitalDocumentPermissionTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDigitalDocumentPermissionTypeTrait(ctx jsonld.Context, x *schema.DigitalDocumentPermissionTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDigitalDocumentPermissionTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDigitalDocumentPermissionTypeTrait(ctx jsonld.Context, x *schema.DigitalDocumentPermissionTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDigitalDocumentPermissionTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}