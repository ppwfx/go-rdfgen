package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBookFormatType strings.Replacer
var strconvBookFormatType strconv.NumError

var basicBookFormatTypeTraitMapping = map[string]func(*schema.BookFormatTypeTrait, []string){}
var customBookFormatTypeTraitMapping = map[string]func(*schema.BookFormatTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BookFormatType", func(ctx jsonld.Context) (interface{}) {
		return NewBookFormatTypeFromContext(ctx)
	})

}

func NewBookFormatTypeFromContext(ctx jsonld.Context) (x schema.BookFormatType) {
	x.Type = "http://schema.org/BookFormatType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBookFormatTypeTrait(ctx, &x.BookFormatTypeTrait)
	MapCustomToBookFormatTypeTrait(ctx, &x.BookFormatTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBookFormatTypeTrait(ctx jsonld.Context, x *schema.BookFormatTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBookFormatTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBookFormatTypeTrait(ctx jsonld.Context, x *schema.BookFormatTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBookFormatTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}