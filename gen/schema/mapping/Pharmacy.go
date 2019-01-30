package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPharmacy strings.Replacer
var strconvPharmacy strconv.NumError

var basicPharmacyTraitMapping = map[string]func(*schema.PharmacyTrait, []string){}
var customPharmacyTraitMapping = map[string]func(*schema.PharmacyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Pharmacy", func(ctx jsonld.Context) (interface{}) {
		return NewPharmacyFromContext(ctx)
	})

}

func NewPharmacyFromContext(ctx jsonld.Context) (x schema.Pharmacy) {
	x.Type = "http://schema.org/Pharmacy"
	MapBasicToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)
	MapCustomToMedicalOrganizationTrait(ctx, &x.MedicalOrganizationTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToMedicalBusinessTrait(ctx, &x.MedicalBusinessTrait)
	MapCustomToMedicalBusinessTrait(ctx, &x.MedicalBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)


	MapBasicToPharmacyTrait(ctx, &x.PharmacyTrait)
	MapCustomToPharmacyTrait(ctx, &x.PharmacyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPharmacyTrait(ctx jsonld.Context, x *schema.PharmacyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPharmacyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPharmacyTrait(ctx jsonld.Context, x *schema.PharmacyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPharmacyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}