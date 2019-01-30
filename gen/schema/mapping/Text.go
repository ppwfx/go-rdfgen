package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsText strings.Replacer
var strconvText strconv.NumError

var basicTextTraitMapping = map[string]func(*schema.TextTrait, []string){}
var customTextTraitMapping = map[string]func(*schema.TextTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Text", func(ctx jsonld.Context) (interface{}) {
		return NewTextFromContext(ctx)
	})

}

func NewTextFromContext(ctx jsonld.Context) (x schema.Text) {
	x.Type = "http://schema.org/Text"

	MapBasicToTextTrait(ctx, &x.TextTrait)
	MapCustomToTextTrait(ctx, &x.TextTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTextTrait(ctx jsonld.Context, x *schema.TextTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTextTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTextTrait(ctx jsonld.Context, x *schema.TextTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTextTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}