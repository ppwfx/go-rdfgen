package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAccommodation strings.Replacer
var strconvAccommodation strconv.NumError

var basicAccommodationTraitMapping = map[string]func(*schema.AccommodationTrait, []string){}
var customAccommodationTraitMapping = map[string]func(*schema.AccommodationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Accommodation", func(ctx jsonld.Context) (interface{}) {
		return NewAccommodationFromContext(ctx)
	})




	basicAccommodationTraitMapping["http://schema.org/permittedUsage"] = func(x *schema.AccommodationTrait, s []string) {
		x.PermittedUsage = s[0]
	}




	customAccommodationTraitMapping["http://schema.org/amenityFeature"] = func(x *schema.AccommodationTrait, ctx jsonld.Context, s string){
		var y schema.LocationFeatureSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewLocationFeatureSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewLocationFeatureSpecification()
			y.Id = s
		}

		x.AmenityFeature = &y

		return
	}

	customAccommodationTraitMapping["http://schema.org/petsAllowed"] = func(x *schema.AccommodationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PetsAllowed, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PetsAllowed = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PetsAllowed = s
		}
	}

	customAccommodationTraitMapping["http://schema.org/numberOfRooms"] = func(x *schema.AccommodationTrait, ctx jsonld.Context, s string){
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

	customAccommodationTraitMapping["http://schema.org/floorSize"] = func(x *schema.AccommodationTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.FloorSize = &y

		return
	}
}

func NewAccommodationFromContext(ctx jsonld.Context) (x schema.Accommodation) {
	x.Type = "http://schema.org/Accommodation"
	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAccommodationTrait(ctx, &x.AccommodationTrait)
	MapCustomToAccommodationTrait(ctx, &x.AccommodationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAccommodationTrait(ctx jsonld.Context, x *schema.AccommodationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAccommodationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAccommodationTrait(ctx jsonld.Context, x *schema.AccommodationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAccommodationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}