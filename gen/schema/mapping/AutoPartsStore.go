package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAutoPartsStore strings.Replacer
var strconvAutoPartsStore strconv.NumError

var basicAutoPartsStoreTraitMapping = map[string]func(*schema.AutoPartsStoreTrait, []string){}
var customAutoPartsStoreTraitMapping = map[string]func(*schema.AutoPartsStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AutoPartsStore", func(ctx jsonld.Context) (interface{}) {
		return NewAutoPartsStoreFromContext(ctx)
	})

}

func NewAutoPartsStoreFromContext(ctx jsonld.Context) (x schema.AutoPartsStore) {
	x.Type = "http://schema.org/AutoPartsStore"
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

	MapBasicToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)
	MapCustomToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)


	MapBasicToAutoPartsStoreTrait(ctx, &x.AutoPartsStoreTrait)
	MapCustomToAutoPartsStoreTrait(ctx, &x.AutoPartsStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAutoPartsStoreTrait(ctx jsonld.Context, x *schema.AutoPartsStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAutoPartsStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAutoPartsStoreTrait(ctx jsonld.Context, x *schema.AutoPartsStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAutoPartsStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}