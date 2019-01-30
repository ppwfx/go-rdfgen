package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsElectronicsStore strings.Replacer
var strconvElectronicsStore strconv.NumError

var basicElectronicsStoreTraitMapping = map[string]func(*schema.ElectronicsStoreTrait, []string){}
var customElectronicsStoreTraitMapping = map[string]func(*schema.ElectronicsStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ElectronicsStore", func(ctx jsonld.Context) (interface{}) {
		return NewElectronicsStoreFromContext(ctx)
	})

}

func NewElectronicsStoreFromContext(ctx jsonld.Context) (x schema.ElectronicsStore) {
	x.Type = "http://schema.org/ElectronicsStore"
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


	MapBasicToElectronicsStoreTrait(ctx, &x.ElectronicsStoreTrait)
	MapCustomToElectronicsStoreTrait(ctx, &x.ElectronicsStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToElectronicsStoreTrait(ctx jsonld.Context, x *schema.ElectronicsStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicElectronicsStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToElectronicsStoreTrait(ctx jsonld.Context, x *schema.ElectronicsStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customElectronicsStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}