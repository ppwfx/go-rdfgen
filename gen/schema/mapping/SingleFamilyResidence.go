package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSingleFamilyResidence strings.Replacer
var strconvSingleFamilyResidence strconv.NumError

var basicSingleFamilyResidenceTraitMapping = map[string]func(*schema.SingleFamilyResidenceTrait, []string){}
var customSingleFamilyResidenceTraitMapping = map[string]func(*schema.SingleFamilyResidenceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SingleFamilyResidence", func(ctx jsonld.Context) (interface{}) {
		return NewSingleFamilyResidenceFromContext(ctx)
	})




	customSingleFamilyResidenceTraitMapping["http://schema.org/occupancy"] = func(x *schema.SingleFamilyResidenceTrait, ctx jsonld.Context, s string){
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

	customSingleFamilyResidenceTraitMapping["http://schema.org/numberOfRooms"] = func(x *schema.SingleFamilyResidenceTrait, ctx jsonld.Context, s string){
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

func NewSingleFamilyResidenceFromContext(ctx jsonld.Context) (x schema.SingleFamilyResidence) {
	x.Type = "http://schema.org/SingleFamilyResidence"
	MapBasicToHouseTrait(ctx, &x.HouseTrait)
	MapCustomToHouseTrait(ctx, &x.HouseTrait)

	MapBasicToAccommodationTrait(ctx, &x.AccommodationTrait)
	MapCustomToAccommodationTrait(ctx, &x.AccommodationTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSingleFamilyResidenceTrait(ctx, &x.SingleFamilyResidenceTrait)
	MapCustomToSingleFamilyResidenceTrait(ctx, &x.SingleFamilyResidenceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSingleFamilyResidenceTrait(ctx jsonld.Context, x *schema.SingleFamilyResidenceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSingleFamilyResidenceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSingleFamilyResidenceTrait(ctx jsonld.Context, x *schema.SingleFamilyResidenceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSingleFamilyResidenceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}