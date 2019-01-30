package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDepartmentStore strings.Replacer
var strconvDepartmentStore strconv.NumError

var basicDepartmentStoreTraitMapping = map[string]func(*schema.DepartmentStoreTrait, []string){}
var customDepartmentStoreTraitMapping = map[string]func(*schema.DepartmentStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DepartmentStore", func(ctx jsonld.Context) (interface{}) {
		return NewDepartmentStoreFromContext(ctx)
	})

}

func NewDepartmentStoreFromContext(ctx jsonld.Context) (x schema.DepartmentStore) {
	x.Type = "http://schema.org/DepartmentStore"
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


	MapBasicToDepartmentStoreTrait(ctx, &x.DepartmentStoreTrait)
	MapCustomToDepartmentStoreTrait(ctx, &x.DepartmentStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDepartmentStoreTrait(ctx jsonld.Context, x *schema.DepartmentStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDepartmentStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDepartmentStoreTrait(ctx jsonld.Context, x *schema.DepartmentStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDepartmentStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}