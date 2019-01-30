package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLocalBusiness strings.Replacer
var strconvLocalBusiness strconv.NumError

var basicLocalBusinessTraitMapping = map[string]func(*schema.LocalBusinessTrait, []string){}
var customLocalBusinessTraitMapping = map[string]func(*schema.LocalBusinessTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LocalBusiness", func(ctx jsonld.Context) (interface{}) {
		return NewLocalBusinessFromContext(ctx)
	})



	basicLocalBusinessTraitMapping["http://schema.org/currenciesAccepted"] = func(x *schema.LocalBusinessTrait, s []string) {
		x.CurrenciesAccepted = s[0]
	}


	basicLocalBusinessTraitMapping["http://schema.org/openingHours"] = func(x *schema.LocalBusinessTrait, s []string) {
		x.OpeningHours = s[0]
	}


	basicLocalBusinessTraitMapping["http://schema.org/paymentAccepted"] = func(x *schema.LocalBusinessTrait, s []string) {
		x.PaymentAccepted = s[0]
	}


	basicLocalBusinessTraitMapping["http://schema.org/priceRange"] = func(x *schema.LocalBusinessTrait, s []string) {
		x.PriceRange = s[0]
	}


	customLocalBusinessTraitMapping["http://schema.org/branchOf"] = func(x *schema.LocalBusinessTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.BranchOf = &y

		return
	}
}

func NewLocalBusinessFromContext(ctx jsonld.Context) (x schema.LocalBusiness) {
	x.Type = "http://schema.org/LocalBusiness"
	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLocalBusinessTrait(ctx jsonld.Context, x *schema.LocalBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLocalBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLocalBusinessTrait(ctx jsonld.Context, x *schema.LocalBusinessTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLocalBusinessTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}