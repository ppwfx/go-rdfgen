package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMovingCompany strings.Replacer
var strconvMovingCompany strconv.NumError

var basicMovingCompanyTraitMapping = map[string]func(*schema.MovingCompanyTrait, []string){}
var customMovingCompanyTraitMapping = map[string]func(*schema.MovingCompanyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MovingCompany", func(ctx jsonld.Context) (interface{}) {
		return NewMovingCompanyFromContext(ctx)
	})

}

func NewMovingCompanyFromContext(ctx jsonld.Context) (x schema.MovingCompany) {
	x.Type = "http://schema.org/MovingCompany"
	MapBasicToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)
	MapCustomToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToMovingCompanyTrait(ctx, &x.MovingCompanyTrait)
	MapCustomToMovingCompanyTrait(ctx, &x.MovingCompanyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMovingCompanyTrait(ctx jsonld.Context, x *schema.MovingCompanyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMovingCompanyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMovingCompanyTrait(ctx jsonld.Context, x *schema.MovingCompanyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMovingCompanyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}