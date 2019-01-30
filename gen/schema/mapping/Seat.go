package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSeat strings.Replacer
var strconvSeat strconv.NumError

var basicSeatTraitMapping = map[string]func(*schema.SeatTrait, []string){}
var customSeatTraitMapping = map[string]func(*schema.SeatTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Seat", func(ctx jsonld.Context) (interface{}) {
		return NewSeatFromContext(ctx)
	})


	basicSeatTraitMapping["http://schema.org/seatNumber"] = func(x *schema.SeatTrait, s []string) {
		x.SeatNumber = s[0]
	}


	basicSeatTraitMapping["http://schema.org/seatRow"] = func(x *schema.SeatTrait, s []string) {
		x.SeatRow = s[0]
	}


	basicSeatTraitMapping["http://schema.org/seatSection"] = func(x *schema.SeatTrait, s []string) {
		x.SeatSection = s[0]
	}



	customSeatTraitMapping["http://schema.org/seatingType"] = func(x *schema.SeatTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SeatingType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SeatingType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SeatingType = s
		}
	}
}

func NewSeatFromContext(ctx jsonld.Context) (x schema.Seat) {
	x.Type = "http://schema.org/Seat"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSeatTrait(ctx, &x.SeatTrait)
	MapCustomToSeatTrait(ctx, &x.SeatTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSeatTrait(ctx jsonld.Context, x *schema.SeatTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSeatTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSeatTrait(ctx jsonld.Context, x *schema.SeatTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSeatTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}