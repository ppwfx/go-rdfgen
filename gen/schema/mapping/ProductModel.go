package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsProductModel strings.Replacer
var strconvProductModel strconv.NumError

var basicProductModelTraitMapping = map[string]func(*schema.ProductModelTrait, []string){}
var customProductModelTraitMapping = map[string]func(*schema.ProductModelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ProductModel", func(ctx jsonld.Context) (interface{}) {
		return NewProductModelFromContext(ctx)
	})





	customProductModelTraitMapping["http://schema.org/isVariantOf"] = func(x *schema.ProductModelTrait, ctx jsonld.Context, s string){
		var y schema.ProductModel
		if strings.HasPrefix(s, "_:") {
			y = NewProductModelFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewProductModel()
			y.Id = s
		}

		x.IsVariantOf = &y

		return
	}

	customProductModelTraitMapping["http://schema.org/predecessorOf"] = func(x *schema.ProductModelTrait, ctx jsonld.Context, s string){
		var y schema.ProductModel
		if strings.HasPrefix(s, "_:") {
			y = NewProductModelFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewProductModel()
			y.Id = s
		}

		x.PredecessorOf = &y

		return
	}

	customProductModelTraitMapping["http://schema.org/successorOf"] = func(x *schema.ProductModelTrait, ctx jsonld.Context, s string){
		var y schema.ProductModel
		if strings.HasPrefix(s, "_:") {
			y = NewProductModelFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewProductModel()
			y.Id = s
		}

		x.SuccessorOf = &y

		return
	}
}

func NewProductModelFromContext(ctx jsonld.Context) (x schema.ProductModel) {
	x.Type = "http://schema.org/ProductModel"
	MapBasicToProductTrait(ctx, &x.ProductTrait)
	MapCustomToProductTrait(ctx, &x.ProductTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToProductModelTrait(ctx, &x.ProductModelTrait)
	MapCustomToProductModelTrait(ctx, &x.ProductModelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToProductModelTrait(ctx jsonld.Context, x *schema.ProductModelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicProductModelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToProductModelTrait(ctx jsonld.Context, x *schema.ProductModelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customProductModelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}