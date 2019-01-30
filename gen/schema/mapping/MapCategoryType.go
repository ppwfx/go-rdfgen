package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMapCategoryType strings.Replacer
var strconvMapCategoryType strconv.NumError

var basicMapCategoryTypeTraitMapping = map[string]func(*schema.MapCategoryTypeTrait, []string){}
var customMapCategoryTypeTraitMapping = map[string]func(*schema.MapCategoryTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MapCategoryType", func(ctx jsonld.Context) (interface{}) {
		return NewMapCategoryTypeFromContext(ctx)
	})

}

func NewMapCategoryTypeFromContext(ctx jsonld.Context) (x schema.MapCategoryType) {
	x.Type = "http://schema.org/MapCategoryType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMapCategoryTypeTrait(ctx, &x.MapCategoryTypeTrait)
	MapCustomToMapCategoryTypeTrait(ctx, &x.MapCategoryTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMapCategoryTypeTrait(ctx jsonld.Context, x *schema.MapCategoryTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMapCategoryTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMapCategoryTypeTrait(ctx jsonld.Context, x *schema.MapCategoryTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMapCategoryTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}