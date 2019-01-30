package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsComputerStore strings.Replacer
var strconvComputerStore strconv.NumError

var basicComputerStoreTraitMapping = map[string]func(*schema.ComputerStoreTrait, []string){}
var customComputerStoreTraitMapping = map[string]func(*schema.ComputerStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ComputerStore", func(ctx jsonld.Context) (interface{}) {
		return NewComputerStoreFromContext(ctx)
	})

}

func NewComputerStoreFromContext(ctx jsonld.Context) (x schema.ComputerStore) {
	x.Type = "http://schema.org/ComputerStore"
	MapBasicToStoreTrait(ctx, &x.StoreTrait)
	MapCustomToStoreTrait(ctx, &x.StoreTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToComputerStoreTrait(ctx, &x.ComputerStoreTrait)
	MapCustomToComputerStoreTrait(ctx, &x.ComputerStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToComputerStoreTrait(ctx jsonld.Context, x *schema.ComputerStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicComputerStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToComputerStoreTrait(ctx jsonld.Context, x *schema.ComputerStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customComputerStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}