package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsApartment strings.Replacer
var strconvApartment strconv.NumError

var basicApartmentTraitMapping = map[string]func(*schema.ApartmentTrait, []string){}
var customApartmentTraitMapping = map[string]func(*schema.ApartmentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Apartment", func(ctx jsonld.Context) (interface{}) {
		return NewApartmentFromContext(ctx)
	})




	customApartmentTraitMapping["http://schema.org/occupancy"] = func(x *schema.ApartmentTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.Occupancy = &y

		return
	}

	customApartmentTraitMapping["http://schema.org/numberOfRooms"] = func(x *schema.ApartmentTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NumberOfRooms, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NumberOfRooms = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NumberOfRooms = s
		}
	}
}

func NewApartmentFromContext(ctx jsonld.Context) (x schema.Apartment) {
	x.Type = "http://schema.org/Apartment"
	MapBasicToAccommodationTrait(ctx, &x.AccommodationTrait)
	MapCustomToAccommodationTrait(ctx, &x.AccommodationTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToApartmentTrait(ctx, &x.ApartmentTrait)
	MapCustomToApartmentTrait(ctx, &x.ApartmentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToApartmentTrait(ctx jsonld.Context, x *schema.ApartmentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicApartmentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToApartmentTrait(ctx jsonld.Context, x *schema.ApartmentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customApartmentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}