package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMobilePhoneStore strings.Replacer
var strconvMobilePhoneStore strconv.NumError

var basicMobilePhoneStoreTraitMapping = map[string]func(*schema.MobilePhoneStoreTrait, []string){}
var customMobilePhoneStoreTraitMapping = map[string]func(*schema.MobilePhoneStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MobilePhoneStore", func(ctx jsonld.Context) (interface{}) {
		return NewMobilePhoneStoreFromContext(ctx)
	})

}

func NewMobilePhoneStoreFromContext(ctx jsonld.Context) (x schema.MobilePhoneStore) {
	x.Type = "http://schema.org/MobilePhoneStore"
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


	MapBasicToMobilePhoneStoreTrait(ctx, &x.MobilePhoneStoreTrait)
	MapCustomToMobilePhoneStoreTrait(ctx, &x.MobilePhoneStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMobilePhoneStoreTrait(ctx jsonld.Context, x *schema.MobilePhoneStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMobilePhoneStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMobilePhoneStoreTrait(ctx jsonld.Context, x *schema.MobilePhoneStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMobilePhoneStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}