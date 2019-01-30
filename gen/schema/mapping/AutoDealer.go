package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAutoDealer strings.Replacer
var strconvAutoDealer strconv.NumError

var basicAutoDealerTraitMapping = map[string]func(*schema.AutoDealerTrait, []string){}
var customAutoDealerTraitMapping = map[string]func(*schema.AutoDealerTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AutoDealer", func(ctx jsonld.Context) (interface{}) {
		return NewAutoDealerFromContext(ctx)
	})

}

func NewAutoDealerFromContext(ctx jsonld.Context) (x schema.AutoDealer) {
	x.Type = "http://schema.org/AutoDealer"
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


	MapBasicToAutoDealerTrait(ctx, &x.AutoDealerTrait)
	MapCustomToAutoDealerTrait(ctx, &x.AutoDealerTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAutoDealerTrait(ctx jsonld.Context, x *schema.AutoDealerTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAutoDealerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAutoDealerTrait(ctx jsonld.Context, x *schema.AutoDealerTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAutoDealerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}