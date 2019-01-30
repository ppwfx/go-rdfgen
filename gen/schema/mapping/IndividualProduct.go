package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsIndividualProduct strings.Replacer
var strconvIndividualProduct strconv.NumError

var basicIndividualProductTraitMapping = map[string]func(*schema.IndividualProductTrait, []string){}
var customIndividualProductTraitMapping = map[string]func(*schema.IndividualProductTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/IndividualProduct", func(ctx jsonld.Context) (interface{}) {
		return NewIndividualProductFromContext(ctx)
	})


	basicIndividualProductTraitMapping["http://schema.org/serialNumber"] = func(x *schema.IndividualProductTrait, s []string) {
		x.SerialNumber = s[0]
	}

}

func NewIndividualProductFromContext(ctx jsonld.Context) (x schema.IndividualProduct) {
	x.Type = "http://schema.org/IndividualProduct"
	MapBasicToProductTrait(ctx, &x.ProductTrait)
	MapCustomToProductTrait(ctx, &x.ProductTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToIndividualProductTrait(ctx, &x.IndividualProductTrait)
	MapCustomToIndividualProductTrait(ctx, &x.IndividualProductTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToIndividualProductTrait(ctx jsonld.Context, x *schema.IndividualProductTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicIndividualProductTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToIndividualProductTrait(ctx jsonld.Context, x *schema.IndividualProductTrait) () {
	for k, v := range ctx.Current {
		f, ok := customIndividualProductTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}