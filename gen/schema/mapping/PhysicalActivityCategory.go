package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPhysicalActivityCategory strings.Replacer
var strconvPhysicalActivityCategory strconv.NumError

var basicPhysicalActivityCategoryTraitMapping = map[string]func(*schema.PhysicalActivityCategoryTrait, []string){}
var customPhysicalActivityCategoryTraitMapping = map[string]func(*schema.PhysicalActivityCategoryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PhysicalActivityCategory", func(ctx jsonld.Context) (interface{}) {
		return NewPhysicalActivityCategoryFromContext(ctx)
	})

}

func NewPhysicalActivityCategoryFromContext(ctx jsonld.Context) (x schema.PhysicalActivityCategory) {
	x.Type = "http://schema.org/PhysicalActivityCategory"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPhysicalActivityCategoryTrait(ctx, &x.PhysicalActivityCategoryTrait)
	MapCustomToPhysicalActivityCategoryTrait(ctx, &x.PhysicalActivityCategoryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPhysicalActivityCategoryTrait(ctx jsonld.Context, x *schema.PhysicalActivityCategoryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPhysicalActivityCategoryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPhysicalActivityCategoryTrait(ctx jsonld.Context, x *schema.PhysicalActivityCategoryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPhysicalActivityCategoryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}