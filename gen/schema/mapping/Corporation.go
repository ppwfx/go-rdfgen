package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCorporation strings.Replacer
var strconvCorporation strconv.NumError

var basicCorporationTraitMapping = map[string]func(*schema.CorporationTrait, []string){}
var customCorporationTraitMapping = map[string]func(*schema.CorporationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Corporation", func(ctx jsonld.Context) (interface{}) {
		return NewCorporationFromContext(ctx)
	})


	basicCorporationTraitMapping["http://schema.org/tickerSymbol"] = func(x *schema.CorporationTrait, s []string) {
		x.TickerSymbol = s[0]
	}

}

func NewCorporationFromContext(ctx jsonld.Context) (x schema.Corporation) {
	x.Type = "http://schema.org/Corporation"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCorporationTrait(ctx, &x.CorporationTrait)
	MapCustomToCorporationTrait(ctx, &x.CorporationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCorporationTrait(ctx jsonld.Context, x *schema.CorporationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCorporationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCorporationTrait(ctx jsonld.Context, x *schema.CorporationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCorporationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}