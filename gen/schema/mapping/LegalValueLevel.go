package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLegalValueLevel strings.Replacer
var strconvLegalValueLevel strconv.NumError

var basicLegalValueLevelTraitMapping = map[string]func(*schema.LegalValueLevelTrait, []string){}
var customLegalValueLevelTraitMapping = map[string]func(*schema.LegalValueLevelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LegalValueLevel", func(ctx jsonld.Context) (interface{}) {
		return NewLegalValueLevelFromContext(ctx)
	})

}

func NewLegalValueLevelFromContext(ctx jsonld.Context) (x schema.LegalValueLevel) {
	x.Type = "http://schema.org/LegalValueLevel"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLegalValueLevelTrait(ctx, &x.LegalValueLevelTrait)
	MapCustomToLegalValueLevelTrait(ctx, &x.LegalValueLevelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLegalValueLevelTrait(ctx jsonld.Context, x *schema.LegalValueLevelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLegalValueLevelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLegalValueLevelTrait(ctx jsonld.Context, x *schema.LegalValueLevelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLegalValueLevelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}