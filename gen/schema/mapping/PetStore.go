package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPetStore strings.Replacer
var strconvPetStore strconv.NumError

var basicPetStoreTraitMapping = map[string]func(*schema.PetStoreTrait, []string){}
var customPetStoreTraitMapping = map[string]func(*schema.PetStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PetStore", func(ctx jsonld.Context) (interface{}) {
		return NewPetStoreFromContext(ctx)
	})

}

func NewPetStoreFromContext(ctx jsonld.Context) (x schema.PetStore) {
	x.Type = "http://schema.org/PetStore"
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


	MapBasicToPetStoreTrait(ctx, &x.PetStoreTrait)
	MapCustomToPetStoreTrait(ctx, &x.PetStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPetStoreTrait(ctx jsonld.Context, x *schema.PetStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPetStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPetStoreTrait(ctx jsonld.Context, x *schema.PetStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPetStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}