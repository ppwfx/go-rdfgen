package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReservationPackage strings.Replacer
var strconvReservationPackage strconv.NumError

var basicReservationPackageTraitMapping = map[string]func(*schema.ReservationPackageTrait, []string){}
var customReservationPackageTraitMapping = map[string]func(*schema.ReservationPackageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReservationPackage", func(ctx jsonld.Context) (interface{}) {
		return NewReservationPackageFromContext(ctx)
	})



	customReservationPackageTraitMapping["http://schema.org/subReservation"] = func(x *schema.ReservationPackageTrait, ctx jsonld.Context, s string){
		var y schema.Reservation
		if strings.HasPrefix(s, "_:") {
			y = NewReservationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewReservation()
			y.Id = s
		}

		x.SubReservation = &y

		return
	}
}

func NewReservationPackageFromContext(ctx jsonld.Context) (x schema.ReservationPackage) {
	x.Type = "http://schema.org/ReservationPackage"
	MapBasicToReservationTrait(ctx, &x.ReservationTrait)
	MapCustomToReservationTrait(ctx, &x.ReservationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReservationPackageTrait(ctx, &x.ReservationPackageTrait)
	MapCustomToReservationPackageTrait(ctx, &x.ReservationPackageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReservationPackageTrait(ctx jsonld.Context, x *schema.ReservationPackageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReservationPackageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReservationPackageTrait(ctx jsonld.Context, x *schema.ReservationPackageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReservationPackageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}