package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSuite strings.Replacer
var strconvSuite strconv.NumError

var basicSuiteTraitMapping = map[string]func(*schema.SuiteTrait, []string){}
var customSuiteTraitMapping = map[string]func(*schema.SuiteTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Suite", func(ctx jsonld.Context) (interface{}) {
		return NewSuiteFromContext(ctx)
	})





	customSuiteTraitMapping["http://schema.org/bed"] = func(x *schema.SuiteTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Bed, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Bed = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Bed = s
		}
	}

	customSuiteTraitMapping["http://schema.org/occupancy"] = func(x *schema.SuiteTrait, ctx jsonld.Context, s string){
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

	customSuiteTraitMapping["http://schema.org/numberOfRooms"] = func(x *schema.SuiteTrait, ctx jsonld.Context, s string){
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

func NewSuiteFromContext(ctx jsonld.Context) (x schema.Suite) {
	x.Type = "http://schema.org/Suite"
	MapBasicToAccommodationTrait(ctx, &x.AccommodationTrait)
	MapCustomToAccommodationTrait(ctx, &x.AccommodationTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSuiteTrait(ctx, &x.SuiteTrait)
	MapCustomToSuiteTrait(ctx, &x.SuiteTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSuiteTrait(ctx jsonld.Context, x *schema.SuiteTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSuiteTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSuiteTrait(ctx jsonld.Context, x *schema.SuiteTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSuiteTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}