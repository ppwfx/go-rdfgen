package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMotorcycleDealer strings.Replacer
var strconvMotorcycleDealer strconv.NumError

var basicMotorcycleDealerTraitMapping = map[string]func(*schema.MotorcycleDealerTrait, []string){}
var customMotorcycleDealerTraitMapping = map[string]func(*schema.MotorcycleDealerTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MotorcycleDealer", func(ctx jsonld.Context) (interface{}) {
		return NewMotorcycleDealerFromContext(ctx)
	})

}

func NewMotorcycleDealerFromContext(ctx jsonld.Context) (x schema.MotorcycleDealer) {
	x.Type = "http://schema.org/MotorcycleDealer"
	MapBasicToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)
	MapCustomToAutomotiveBusinessTrait(ctx, &x.AutomotiveBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToMotorcycleDealerTrait(ctx, &x.MotorcycleDealerTrait)
	MapCustomToMotorcycleDealerTrait(ctx, &x.MotorcycleDealerTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMotorcycleDealerTrait(ctx jsonld.Context, x *schema.MotorcycleDealerTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMotorcycleDealerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMotorcycleDealerTrait(ctx jsonld.Context, x *schema.MotorcycleDealerTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMotorcycleDealerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}