package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCode strings.Replacer
var strconvCode strconv.NumError

var basicCodeTraitMapping = map[string]func(*schema.CodeTrait, []string){}
var customCodeTraitMapping = map[string]func(*schema.CodeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Code", func(ctx jsonld.Context) (interface{}) {
		return NewCodeFromContext(ctx)
	})

}

func NewCodeFromContext(ctx jsonld.Context) (x schema.Code) {
	x.Type = "http://schema.org/Code"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCodeTrait(ctx, &x.CodeTrait)
	MapCustomToCodeTrait(ctx, &x.CodeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCodeTrait(ctx jsonld.Context, x *schema.CodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCodeTrait(ctx jsonld.Context, x *schema.CodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}