package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFlightReservation strings.Replacer
var strconvFlightReservation strconv.NumError

var basicFlightReservationTraitMapping = map[string]func(*schema.FlightReservationTrait, []string){}
var customFlightReservationTraitMapping = map[string]func(*schema.FlightReservationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FlightReservation", func(ctx jsonld.Context) (interface{}) {
		return NewFlightReservationFromContext(ctx)
	})



	basicFlightReservationTraitMapping["http://schema.org/passengerSequenceNumber"] = func(x *schema.FlightReservationTrait, s []string) {
		x.PassengerSequenceNumber = s[0]
	}


	basicFlightReservationTraitMapping["http://schema.org/securityScreening"] = func(x *schema.FlightReservationTrait, s []string) {
		x.SecurityScreening = s[0]
	}


	basicFlightReservationTraitMapping["http://schema.org/boardingGroup"] = func(x *schema.FlightReservationTrait, s []string) {
		x.BoardingGroup = s[0]
	}


	customFlightReservationTraitMapping["http://schema.org/passengerPriorityStatus"] = func(x *schema.FlightReservationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PassengerPriorityStatus, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PassengerPriorityStatus = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PassengerPriorityStatus = s
		}
	}
}

func NewFlightReservationFromContext(ctx jsonld.Context) (x schema.FlightReservation) {
	x.Type = "http://schema.org/FlightReservation"
	MapBasicToReservationTrait(ctx, &x.ReservationTrait)
	MapCustomToReservationTrait(ctx, &x.ReservationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFlightReservationTrait(ctx, &x.FlightReservationTrait)
	MapCustomToFlightReservationTrait(ctx, &x.FlightReservationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFlightReservationTrait(ctx jsonld.Context, x *schema.FlightReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFlightReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFlightReservationTrait(ctx jsonld.Context, x *schema.FlightReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFlightReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}