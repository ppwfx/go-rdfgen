package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBreadcrumbList strings.Replacer
var strconvBreadcrumbList strconv.NumError

var basicBreadcrumbListTraitMapping = map[string]func(*schema.BreadcrumbListTrait, []string){}
var customBreadcrumbListTraitMapping = map[string]func(*schema.BreadcrumbListTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BreadcrumbList", func(ctx jsonld.Context) (interface{}) {
		return NewBreadcrumbListFromContext(ctx)
	})

}

func NewBreadcrumbListFromContext(ctx jsonld.Context) (x schema.BreadcrumbList) {
	x.Type = "http://schema.org/BreadcrumbList"
	MapBasicToItemListTrait(ctx, &x.ItemListTrait)
	MapCustomToItemListTrait(ctx, &x.ItemListTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBreadcrumbListTrait(ctx, &x.BreadcrumbListTrait)
	MapCustomToBreadcrumbListTrait(ctx, &x.BreadcrumbListTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBreadcrumbListTrait(ctx jsonld.Context, x *schema.BreadcrumbListTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBreadcrumbListTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBreadcrumbListTrait(ctx jsonld.Context, x *schema.BreadcrumbListTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBreadcrumbListTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}