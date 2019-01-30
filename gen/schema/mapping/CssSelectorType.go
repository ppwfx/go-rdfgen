package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCssSelectorType strings.Replacer
var strconvCssSelectorType strconv.NumError

var basicCssSelectorTypeTraitMapping = map[string]func(*schema.CssSelectorTypeTrait, []string){}
var customCssSelectorTypeTraitMapping = map[string]func(*schema.CssSelectorTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CssSelectorType", func(ctx jsonld.Context) (interface{}) {
		return NewCssSelectorTypeFromContext(ctx)
	})

}

func NewCssSelectorTypeFromContext(ctx jsonld.Context) (x schema.CssSelectorType) {
	x.Type = "http://schema.org/CssSelectorType"
	MapBasicToTextTrait(ctx, &x.TextTrait)
	MapCustomToTextTrait(ctx, &x.TextTrait)


	MapBasicToCssSelectorTypeTrait(ctx, &x.CssSelectorTypeTrait)
	MapCustomToCssSelectorTypeTrait(ctx, &x.CssSelectorTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCssSelectorTypeTrait(ctx jsonld.Context, x *schema.CssSelectorTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCssSelectorTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCssSelectorTypeTrait(ctx jsonld.Context, x *schema.CssSelectorTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCssSelectorTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}