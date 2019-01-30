package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDate strings.Replacer
var strconvDate strconv.NumError

var basicDateTraitMapping = map[string]func(*schema.DateTrait, []string){}
var customDateTraitMapping = map[string]func(*schema.DateTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Date", func(ctx jsonld.Context) (interface{}) {
		return NewDateFromContext(ctx)
	})

}

func NewDateFromContext(ctx jsonld.Context) (x schema.Date) {
	x.Type = "http://schema.org/Date"

	MapBasicToDateTrait(ctx, &x.DateTrait)
	MapCustomToDateTrait(ctx, &x.DateTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDateTrait(ctx jsonld.Context, x *schema.DateTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDateTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDateTrait(ctx jsonld.Context, x *schema.DateTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDateTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}