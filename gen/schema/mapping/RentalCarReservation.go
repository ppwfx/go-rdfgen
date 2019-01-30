package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRentalCarReservation strings.Replacer
var strconvRentalCarReservation strconv.NumError

var basicRentalCarReservationTraitMapping = map[string]func(*schema.RentalCarReservationTrait, []string){}
var customRentalCarReservationTraitMapping = map[string]func(*schema.RentalCarReservationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RentalCarReservation", func(ctx jsonld.Context) (interface{}) {
		return NewRentalCarReservationFromContext(ctx)
	})






	customRentalCarReservationTraitMapping["http://schema.org/pickupTime"] = func(x *schema.RentalCarReservationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.PickupTime = &y

		return
	}

	customRentalCarReservationTraitMapping["http://schema.org/dropoffTime"] = func(x *schema.RentalCarReservationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.DropoffTime = &y

		return
	}

	customRentalCarReservationTraitMapping["http://schema.org/pickupLocation"] = func(x *schema.RentalCarReservationTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.PickupLocation = &y

		return
	}

	customRentalCarReservationTraitMapping["http://schema.org/dropoffLocation"] = func(x *schema.RentalCarReservationTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.DropoffLocation = &y

		return
	}
}

func NewRentalCarReservationFromContext(ctx jsonld.Context) (x schema.RentalCarReservation) {
	x.Type = "http://schema.org/RentalCarReservation"
	MapBasicToReservationTrait(ctx, &x.ReservationTrait)
	MapCustomToReservationTrait(ctx, &x.ReservationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRentalCarReservationTrait(ctx, &x.RentalCarReservationTrait)
	MapCustomToRentalCarReservationTrait(ctx, &x.RentalCarReservationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRentalCarReservationTrait(ctx jsonld.Context, x *schema.RentalCarReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRentalCarReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRentalCarReservationTrait(ctx jsonld.Context, x *schema.RentalCarReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRentalCarReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}