package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOfficeEquipmentStore strings.Replacer
var strconvOfficeEquipmentStore strconv.NumError

var basicOfficeEquipmentStoreTraitMapping = map[string]func(*schema.OfficeEquipmentStoreTrait, []string){}
var customOfficeEquipmentStoreTraitMapping = map[string]func(*schema.OfficeEquipmentStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OfficeEquipmentStore", func(ctx jsonld.Context) (interface{}) {
		return NewOfficeEquipmentStoreFromContext(ctx)
	})

}

func NewOfficeEquipmentStoreFromContext(ctx jsonld.Context) (x schema.OfficeEquipmentStore) {
	x.Type = "http://schema.org/OfficeEquipmentStore"
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


	MapBasicToOfficeEquipmentStoreTrait(ctx, &x.OfficeEquipmentStoreTrait)
	MapCustomToOfficeEquipmentStoreTrait(ctx, &x.OfficeEquipmentStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOfficeEquipmentStoreTrait(ctx jsonld.Context, x *schema.OfficeEquipmentStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOfficeEquipmentStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOfficeEquipmentStoreTrait(ctx jsonld.Context, x *schema.OfficeEquipmentStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOfficeEquipmentStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}