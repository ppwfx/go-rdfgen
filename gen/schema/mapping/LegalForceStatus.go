package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLegalForceStatus strings.Replacer
var strconvLegalForceStatus strconv.NumError

var basicLegalForceStatusTraitMapping = map[string]func(*schema.LegalForceStatusTrait, []string){}
var customLegalForceStatusTraitMapping = map[string]func(*schema.LegalForceStatusTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LegalForceStatus", func(ctx jsonld.Context) (interface{}) {
		return NewLegalForceStatusFromContext(ctx)
	})

}

func NewLegalForceStatusFromContext(ctx jsonld.Context) (x schema.LegalForceStatus) {
	x.Type = "http://schema.org/LegalForceStatus"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLegalForceStatusTrait(ctx, &x.LegalForceStatusTrait)
	MapCustomToLegalForceStatusTrait(ctx, &x.LegalForceStatusTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLegalForceStatusTrait(ctx jsonld.Context, x *schema.LegalForceStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLegalForceStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLegalForceStatusTrait(ctx jsonld.Context, x *schema.LegalForceStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLegalForceStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}