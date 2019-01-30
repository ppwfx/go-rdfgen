package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBusinessEntityType strings.Replacer
var strconvBusinessEntityType strconv.NumError

var basicBusinessEntityTypeTraitMapping = map[string]func(*schema.BusinessEntityTypeTrait, []string){}
var customBusinessEntityTypeTraitMapping = map[string]func(*schema.BusinessEntityTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BusinessEntityType", func(ctx jsonld.Context) (interface{}) {
		return NewBusinessEntityTypeFromContext(ctx)
	})

}

func NewBusinessEntityTypeFromContext(ctx jsonld.Context) (x schema.BusinessEntityType) {
	x.Type = "http://schema.org/BusinessEntityType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBusinessEntityTypeTrait(ctx, &x.BusinessEntityTypeTrait)
	MapCustomToBusinessEntityTypeTrait(ctx, &x.BusinessEntityTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBusinessEntityTypeTrait(ctx jsonld.Context, x *schema.BusinessEntityTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBusinessEntityTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBusinessEntityTypeTrait(ctx jsonld.Context, x *schema.BusinessEntityTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBusinessEntityTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}