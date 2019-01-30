package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHardwareStore strings.Replacer
var strconvHardwareStore strconv.NumError

var basicHardwareStoreTraitMapping = map[string]func(*schema.HardwareStoreTrait, []string){}
var customHardwareStoreTraitMapping = map[string]func(*schema.HardwareStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HardwareStore", func(ctx jsonld.Context) (interface{}) {
		return NewHardwareStoreFromContext(ctx)
	})

}

func NewHardwareStoreFromContext(ctx jsonld.Context) (x schema.HardwareStore) {
	x.Type = "http://schema.org/HardwareStore"
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


	MapBasicToHardwareStoreTrait(ctx, &x.HardwareStoreTrait)
	MapCustomToHardwareStoreTrait(ctx, &x.HardwareStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHardwareStoreTrait(ctx jsonld.Context, x *schema.HardwareStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHardwareStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHardwareStoreTrait(ctx jsonld.Context, x *schema.HardwareStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHardwareStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}