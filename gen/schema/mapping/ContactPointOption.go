package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsContactPointOption strings.Replacer
var strconvContactPointOption strconv.NumError

var basicContactPointOptionTraitMapping = map[string]func(*schema.ContactPointOptionTrait, []string){}
var customContactPointOptionTraitMapping = map[string]func(*schema.ContactPointOptionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ContactPointOption", func(ctx jsonld.Context) (interface{}) {
		return NewContactPointOptionFromContext(ctx)
	})

}

func NewContactPointOptionFromContext(ctx jsonld.Context) (x schema.ContactPointOption) {
	x.Type = "http://schema.org/ContactPointOption"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToContactPointOptionTrait(ctx, &x.ContactPointOptionTrait)
	MapCustomToContactPointOptionTrait(ctx, &x.ContactPointOptionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToContactPointOptionTrait(ctx jsonld.Context, x *schema.ContactPointOptionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicContactPointOptionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToContactPointOptionTrait(ctx jsonld.Context, x *schema.ContactPointOptionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customContactPointOptionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}