package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDryCleaningOrLaundry strings.Replacer
var strconvDryCleaningOrLaundry strconv.NumError

var basicDryCleaningOrLaundryTraitMapping = map[string]func(*schema.DryCleaningOrLaundryTrait, []string){}
var customDryCleaningOrLaundryTraitMapping = map[string]func(*schema.DryCleaningOrLaundryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DryCleaningOrLaundry", func(ctx jsonld.Context) (interface{}) {
		return NewDryCleaningOrLaundryFromContext(ctx)
	})

}

func NewDryCleaningOrLaundryFromContext(ctx jsonld.Context) (x schema.DryCleaningOrLaundry) {
	x.Type = "http://schema.org/DryCleaningOrLaundry"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToDryCleaningOrLaundryTrait(ctx, &x.DryCleaningOrLaundryTrait)
	MapCustomToDryCleaningOrLaundryTrait(ctx, &x.DryCleaningOrLaundryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDryCleaningOrLaundryTrait(ctx jsonld.Context, x *schema.DryCleaningOrLaundryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDryCleaningOrLaundryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDryCleaningOrLaundryTrait(ctx jsonld.Context, x *schema.DryCleaningOrLaundryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDryCleaningOrLaundryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}