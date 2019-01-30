package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCategoryCodeSet strings.Replacer
var strconvCategoryCodeSet strconv.NumError

var basicCategoryCodeSetTraitMapping = map[string]func(*schema.CategoryCodeSetTrait, []string){}
var customCategoryCodeSetTraitMapping = map[string]func(*schema.CategoryCodeSetTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CategoryCodeSet", func(ctx jsonld.Context) (interface{}) {
		return NewCategoryCodeSetFromContext(ctx)
	})



	customCategoryCodeSetTraitMapping["http://schema.org/hasCategoryCode"] = func(x *schema.CategoryCodeSetTrait, ctx jsonld.Context, s string){
		var y schema.CategoryCode
		if strings.HasPrefix(s, "_:") {
			y = NewCategoryCodeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCategoryCode()
			y.Id = s
		}

		x.HasCategoryCode = &y

		return
	}
}

func NewCategoryCodeSetFromContext(ctx jsonld.Context) (x schema.CategoryCodeSet) {
	x.Type = "http://schema.org/CategoryCodeSet"
	MapBasicToDefinedTermSetTrait(ctx, &x.DefinedTermSetTrait)
	MapCustomToDefinedTermSetTrait(ctx, &x.DefinedTermSetTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCategoryCodeSetTrait(ctx, &x.CategoryCodeSetTrait)
	MapCustomToCategoryCodeSetTrait(ctx, &x.CategoryCodeSetTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCategoryCodeSetTrait(ctx jsonld.Context, x *schema.CategoryCodeSetTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCategoryCodeSetTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCategoryCodeSetTrait(ctx jsonld.Context, x *schema.CategoryCodeSetTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCategoryCodeSetTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}