package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSelfStorage strings.Replacer
var strconvSelfStorage strconv.NumError

var basicSelfStorageTraitMapping = map[string]func(*schema.SelfStorageTrait, []string){}
var customSelfStorageTraitMapping = map[string]func(*schema.SelfStorageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SelfStorage", func(ctx jsonld.Context) (interface{}) {
		return NewSelfStorageFromContext(ctx)
	})

}

func NewSelfStorageFromContext(ctx jsonld.Context) (x schema.SelfStorage) {
	x.Type = "http://schema.org/SelfStorage"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToSelfStorageTrait(ctx, &x.SelfStorageTrait)
	MapCustomToSelfStorageTrait(ctx, &x.SelfStorageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSelfStorageTrait(ctx jsonld.Context, x *schema.SelfStorageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSelfStorageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSelfStorageTrait(ctx jsonld.Context, x *schema.SelfStorageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSelfStorageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}