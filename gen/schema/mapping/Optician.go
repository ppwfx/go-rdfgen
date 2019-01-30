package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOptician strings.Replacer
var strconvOptician strconv.NumError

var basicOpticianTraitMapping = map[string]func(*schema.OpticianTrait, []string){}
var customOpticianTraitMapping = map[string]func(*schema.OpticianTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Optician", func(ctx jsonld.Context) (interface{}) {
		return NewOpticianFromContext(ctx)
	})

}

func NewOpticianFromContext(ctx jsonld.Context) (x schema.Optician) {
	x.Type = "http://schema.org/Optician"
	MapBasicToMedicalBusinessTrait(ctx, &x.MedicalBusinessTrait)
	MapCustomToMedicalBusinessTrait(ctx, &x.MedicalBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToOpticianTrait(ctx, &x.OpticianTrait)
	MapCustomToOpticianTrait(ctx, &x.OpticianTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOpticianTrait(ctx jsonld.Context, x *schema.OpticianTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOpticianTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOpticianTrait(ctx jsonld.Context, x *schema.OpticianTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOpticianTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}