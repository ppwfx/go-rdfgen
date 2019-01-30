package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWarrantyScope strings.Replacer
var strconvWarrantyScope strconv.NumError

var basicWarrantyScopeTraitMapping = map[string]func(*schema.WarrantyScopeTrait, []string){}
var customWarrantyScopeTraitMapping = map[string]func(*schema.WarrantyScopeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WarrantyScope", func(ctx jsonld.Context) (interface{}) {
		return NewWarrantyScopeFromContext(ctx)
	})

}

func NewWarrantyScopeFromContext(ctx jsonld.Context) (x schema.WarrantyScope) {
	x.Type = "http://schema.org/WarrantyScope"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWarrantyScopeTrait(ctx, &x.WarrantyScopeTrait)
	MapCustomToWarrantyScopeTrait(ctx, &x.WarrantyScopeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWarrantyScopeTrait(ctx jsonld.Context, x *schema.WarrantyScopeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWarrantyScopeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWarrantyScopeTrait(ctx jsonld.Context, x *schema.WarrantyScopeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWarrantyScopeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}