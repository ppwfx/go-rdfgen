package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSomeProducts strings.Replacer
var strconvSomeProducts strconv.NumError

var basicSomeProductsTraitMapping = map[string]func(*schema.SomeProductsTrait, []string){}
var customSomeProductsTraitMapping = map[string]func(*schema.SomeProductsTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SomeProducts", func(ctx jsonld.Context) (interface{}) {
		return NewSomeProductsFromContext(ctx)
	})



	customSomeProductsTraitMapping["http://schema.org/inventoryLevel"] = func(x *schema.SomeProductsTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.InventoryLevel = &y

		return
	}
}

func NewSomeProductsFromContext(ctx jsonld.Context) (x schema.SomeProducts) {
	x.Type = "http://schema.org/SomeProducts"
	MapBasicToProductTrait(ctx, &x.ProductTrait)
	MapCustomToProductTrait(ctx, &x.ProductTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSomeProductsTrait(ctx, &x.SomeProductsTrait)
	MapCustomToSomeProductsTrait(ctx, &x.SomeProductsTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSomeProductsTrait(ctx jsonld.Context, x *schema.SomeProductsTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSomeProductsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSomeProductsTrait(ctx jsonld.Context, x *schema.SomeProductsTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSomeProductsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}