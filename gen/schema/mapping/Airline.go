package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAirline strings.Replacer
var strconvAirline strconv.NumError

var basicAirlineTraitMapping = map[string]func(*schema.AirlineTrait, []string){}
var customAirlineTraitMapping = map[string]func(*schema.AirlineTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Airline", func(ctx jsonld.Context) (interface{}) {
		return NewAirlineFromContext(ctx)
	})


	basicAirlineTraitMapping["http://schema.org/iataCode"] = func(x *schema.AirlineTrait, s []string) {
		x.IataCode = s[0]
	}



	customAirlineTraitMapping["http://schema.org/boardingPolicy"] = func(x *schema.AirlineTrait, ctx jsonld.Context, s string){
		var y schema.BoardingPolicyType
		if strings.HasPrefix(s, "_:") {
			y = NewBoardingPolicyTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBoardingPolicyType()
			y.Id = s
		}

		x.BoardingPolicy = &y

		return
	}
}

func NewAirlineFromContext(ctx jsonld.Context) (x schema.Airline) {
	x.Type = "http://schema.org/Airline"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAirlineTrait(ctx, &x.AirlineTrait)
	MapCustomToAirlineTrait(ctx, &x.AirlineTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAirlineTrait(ctx jsonld.Context, x *schema.AirlineTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAirlineTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAirlineTrait(ctx jsonld.Context, x *schema.AirlineTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAirlineTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}