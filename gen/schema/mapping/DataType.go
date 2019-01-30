package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDataType strings.Replacer
var strconvDataType strconv.NumError

var basicDataTypeTraitMapping = map[string]func(*schema.DataTypeTrait, []string){}
var customDataTypeTraitMapping = map[string]func(*schema.DataTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DataType", func(ctx jsonld.Context) (interface{}) {
		return NewDataTypeFromContext(ctx)
	})

}

func NewDataTypeFromContext(ctx jsonld.Context) (x schema.DataType) {
	x.Type = "http://schema.org/DataType"

	MapBasicToDataTypeTrait(ctx, &x.DataTypeTrait)
	MapCustomToDataTypeTrait(ctx, &x.DataTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDataTypeTrait(ctx jsonld.Context, x *schema.DataTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDataTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDataTypeTrait(ctx jsonld.Context, x *schema.DataTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDataTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}