package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBedDetails strings.Replacer
var strconvBedDetails strconv.NumError

var basicBedDetailsTraitMapping = map[string]func(*schema.BedDetailsTrait, []string){}
var customBedDetailsTraitMapping = map[string]func(*schema.BedDetailsTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BedDetails", func(ctx jsonld.Context) (interface{}) {
		return NewBedDetailsFromContext(ctx)
	})


	basicBedDetailsTraitMapping["http://schema.org/numberOfBeds"] = func(x *schema.BedDetailsTrait, s []string) {
		var err error
		x.NumberOfBeds, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}



	customBedDetailsTraitMapping["http://schema.org/typeOfBed"] = func(x *schema.BedDetailsTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TypeOfBed, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TypeOfBed = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TypeOfBed = s
		}
	}
}

func NewBedDetailsFromContext(ctx jsonld.Context) (x schema.BedDetails) {
	x.Type = "http://schema.org/BedDetails"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBedDetailsTrait(ctx, &x.BedDetailsTrait)
	MapCustomToBedDetailsTrait(ctx, &x.BedDetailsTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBedDetailsTrait(ctx jsonld.Context, x *schema.BedDetailsTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBedDetailsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBedDetailsTrait(ctx jsonld.Context, x *schema.BedDetailsTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBedDetailsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}