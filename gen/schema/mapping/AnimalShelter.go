package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAnimalShelter strings.Replacer
var strconvAnimalShelter strconv.NumError

var basicAnimalShelterTraitMapping = map[string]func(*schema.AnimalShelterTrait, []string){}
var customAnimalShelterTraitMapping = map[string]func(*schema.AnimalShelterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AnimalShelter", func(ctx jsonld.Context) (interface{}) {
		return NewAnimalShelterFromContext(ctx)
	})

}

func NewAnimalShelterFromContext(ctx jsonld.Context) (x schema.AnimalShelter) {
	x.Type = "http://schema.org/AnimalShelter"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToAnimalShelterTrait(ctx, &x.AnimalShelterTrait)
	MapCustomToAnimalShelterTrait(ctx, &x.AnimalShelterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAnimalShelterTrait(ctx jsonld.Context, x *schema.AnimalShelterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAnimalShelterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAnimalShelterTrait(ctx jsonld.Context, x *schema.AnimalShelterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAnimalShelterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}