package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsItemListOrderType strings.Replacer
var strconvItemListOrderType strconv.NumError

var basicItemListOrderTypeTraitMapping = map[string]func(*schema.ItemListOrderTypeTrait, []string){}
var customItemListOrderTypeTraitMapping = map[string]func(*schema.ItemListOrderTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ItemListOrderType", func(ctx jsonld.Context) (interface{}) {
		return NewItemListOrderTypeFromContext(ctx)
	})

}

func NewItemListOrderTypeFromContext(ctx jsonld.Context) (x schema.ItemListOrderType) {
	x.Type = "http://schema.org/ItemListOrderType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToItemListOrderTypeTrait(ctx, &x.ItemListOrderTypeTrait)
	MapCustomToItemListOrderTypeTrait(ctx, &x.ItemListOrderTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToItemListOrderTypeTrait(ctx jsonld.Context, x *schema.ItemListOrderTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicItemListOrderTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToItemListOrderTypeTrait(ctx jsonld.Context, x *schema.ItemListOrderTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customItemListOrderTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}