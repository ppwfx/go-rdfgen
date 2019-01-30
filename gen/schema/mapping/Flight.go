package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFlight strings.Replacer
var strconvFlight strconv.NumError

var basicFlightTraitMapping = map[string]func(*schema.FlightTrait, []string){}
var customFlightTraitMapping = map[string]func(*schema.FlightTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Flight", func(ctx jsonld.Context) (interface{}) {
		return NewFlightFromContext(ctx)
	})




	basicFlightTraitMapping["http://schema.org/flightNumber"] = func(x *schema.FlightTrait, s []string) {
		x.FlightNumber = s[0]
	}




	basicFlightTraitMapping["http://schema.org/departureGate"] = func(x *schema.FlightTrait, s []string) {
		x.DepartureGate = s[0]
	}


	basicFlightTraitMapping["http://schema.org/arrivalTerminal"] = func(x *schema.FlightTrait, s []string) {
		x.ArrivalTerminal = s[0]
	}



	basicFlightTraitMapping["http://schema.org/departureTerminal"] = func(x *schema.FlightTrait, s []string) {
		x.DepartureTerminal = s[0]
	}


	basicFlightTraitMapping["http://schema.org/arrivalGate"] = func(x *schema.FlightTrait, s []string) {
		x.ArrivalGate = s[0]
	}


	basicFlightTraitMapping["http://schema.org/mealService"] = func(x *schema.FlightTrait, s []string) {
		x.MealService = s[0]
	}






	customFlightTraitMapping["http://schema.org/seller"] = func(x *schema.FlightTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Seller, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Seller = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Seller = s
		}
	}

	customFlightTraitMapping["http://schema.org/carrier"] = func(x *schema.FlightTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.Carrier = &y

		return
	}

	customFlightTraitMapping["http://schema.org/webCheckinTime"] = func(x *schema.FlightTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.WebCheckinTime = &y

		return
	}

	customFlightTraitMapping["http://schema.org/flightDistance"] = func(x *schema.FlightTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.FlightDistance, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.FlightDistance = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.FlightDistance = s
		}
	}

	customFlightTraitMapping["http://schema.org/aircraft"] = func(x *schema.FlightTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Aircraft, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Aircraft = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Aircraft = s
		}
	}

	customFlightTraitMapping["http://schema.org/estimatedFlightDuration"] = func(x *schema.FlightTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EstimatedFlightDuration, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EstimatedFlightDuration = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EstimatedFlightDuration = s
		}
	}

	customFlightTraitMapping["http://schema.org/arrivalAirport"] = func(x *schema.FlightTrait, ctx jsonld.Context, s string){
		var y schema.Airport
		if strings.HasPrefix(s, "_:") {
			y = NewAirportFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAirport()
			y.Id = s
		}

		x.ArrivalAirport = &y

		return
	}

	customFlightTraitMapping["http://schema.org/boardingPolicy"] = func(x *schema.FlightTrait, ctx jsonld.Context, s string){
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

	customFlightTraitMapping["http://schema.org/departureAirport"] = func(x *schema.FlightTrait, ctx jsonld.Context, s string){
		var y schema.Airport
		if strings.HasPrefix(s, "_:") {
			y = NewAirportFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAirport()
			y.Id = s
		}

		x.DepartureAirport = &y

		return
	}
}

func NewFlightFromContext(ctx jsonld.Context) (x schema.Flight) {
	x.Type = "http://schema.org/Flight"
	MapBasicToTripTrait(ctx, &x.TripTrait)
	MapCustomToTripTrait(ctx, &x.TripTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFlightTrait(ctx, &x.FlightTrait)
	MapCustomToFlightTrait(ctx, &x.FlightTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFlightTrait(ctx jsonld.Context, x *schema.FlightTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFlightTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFlightTrait(ctx jsonld.Context, x *schema.FlightTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFlightTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}