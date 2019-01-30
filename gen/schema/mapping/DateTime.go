package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDateTime strings.Replacer
var strconvDateTime strconv.NumError

var basicDateTimeTraitMapping = map[string]func(*schema.DateTimeTrait, []string){}
var customDateTimeTraitMapping = map[string]func(*schema.DateTimeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DateTime", func(ctx jsonld.Context) (interface{}) {
		return NewDateTimeFromContext(ctx)
	})

}

func NewDateTimeFromContext(ctx jsonld.Context) (x schema.DateTime) {
	x.Type = "http://schema.org/DateTime"

	MapBasicToDateTimeTrait(ctx, &x.DateTimeTrait)
	MapCustomToDateTimeTrait(ctx, &x.DateTimeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDateTimeTrait(ctx jsonld.Context, x *schema.DateTimeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDateTimeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDateTimeTrait(ctx jsonld.Context, x *schema.DateTimeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDateTimeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}