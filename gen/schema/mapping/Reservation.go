package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReservation strings.Replacer
var strconvReservation strconv.NumError

var basicReservationTraitMapping = map[string]func(*schema.ReservationTrait, []string){}
var customReservationTraitMapping = map[string]func(*schema.ReservationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Reservation", func(ctx jsonld.Context) (interface{}) {
		return NewReservationFromContext(ctx)
	})



	basicReservationTraitMapping["http://schema.org/priceCurrency"] = func(x *schema.ReservationTrait, s []string) {
		x.PriceCurrency = s[0]
	}








	basicReservationTraitMapping["http://schema.org/reservationId"] = func(x *schema.ReservationTrait, s []string) {
		x.ReservationId = s[0]
	}






	customReservationTraitMapping["http://schema.org/provider"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Provider, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Provider = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Provider = s
		}
	}

	customReservationTraitMapping["http://schema.org/broker"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Broker, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Broker = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Broker = s
		}
	}

	customReservationTraitMapping["http://schema.org/bookingAgent"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BookingAgent, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BookingAgent = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BookingAgent = s
		}
	}

	customReservationTraitMapping["http://schema.org/bookingTime"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.BookingTime = &y

		return
	}

	customReservationTraitMapping["http://schema.org/modifiedTime"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ModifiedTime = &y

		return
	}

	customReservationTraitMapping["http://schema.org/programMembershipUsed"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		var y schema.ProgramMembership
		if strings.HasPrefix(s, "_:") {
			y = NewProgramMembershipFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewProgramMembership()
			y.Id = s
		}

		x.ProgramMembershipUsed = &y

		return
	}

	customReservationTraitMapping["http://schema.org/reservationFor"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.ReservationFor = &y

		return
	}

	customReservationTraitMapping["http://schema.org/reservationStatus"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		var y schema.ReservationStatusType
		if strings.HasPrefix(s, "_:") {
			y = NewReservationStatusTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewReservationStatusType()
			y.Id = s
		}

		x.ReservationStatus = &y

		return
	}

	customReservationTraitMapping["http://schema.org/reservedTicket"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		var y schema.Ticket
		if strings.HasPrefix(s, "_:") {
			y = NewTicketFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewTicket()
			y.Id = s
		}

		x.ReservedTicket = &y

		return
	}

	customReservationTraitMapping["http://schema.org/totalPrice"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TotalPrice, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TotalPrice = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TotalPrice = s
		}
	}

	customReservationTraitMapping["http://schema.org/underName"] = func(x *schema.ReservationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.UnderName, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.UnderName = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.UnderName = s
		}
	}
}

func NewReservationFromContext(ctx jsonld.Context) (x schema.Reservation) {
	x.Type = "http://schema.org/Reservation"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReservationTrait(ctx, &x.ReservationTrait)
	MapCustomToReservationTrait(ctx, &x.ReservationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReservationTrait(ctx jsonld.Context, x *schema.ReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReservationTrait(ctx jsonld.Context, x *schema.ReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}