package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTime strings.Replacer
var strconvTime strconv.NumError

var basicTimeTraitMapping = map[string]func(*schema.TimeTrait, []string){}
var customTimeTraitMapping = map[string]func(*schema.TimeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Time", func(ctx jsonld.Context) (interface{}) {
		return NewTimeFromContext(ctx)
	})

}

func NewTimeFromContext(ctx jsonld.Context) (x schema.Time) {
	x.Type = "http://schema.org/Time"

	MapBasicToTimeTrait(ctx, &x.TimeTrait)
	MapCustomToTimeTrait(ctx, &x.TimeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTimeTrait(ctx jsonld.Context, x *schema.TimeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTimeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTimeTrait(ctx jsonld.Context, x *schema.TimeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTimeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}