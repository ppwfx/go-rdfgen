package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTrainTrip strings.Replacer
var strconvTrainTrip strconv.NumError

var basicTrainTripTraitMapping = map[string]func(*schema.TrainTripTrait, []string){}
var customTrainTripTraitMapping = map[string]func(*schema.TrainTripTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TrainTrip", func(ctx jsonld.Context) (interface{}) {
		return NewTrainTripFromContext(ctx)
	})


	basicTrainTripTraitMapping["http://schema.org/departurePlatform"] = func(x *schema.TrainTripTrait, s []string) {
		x.DeparturePlatform = s[0]
	}


	basicTrainTripTraitMapping["http://schema.org/arrivalPlatform"] = func(x *schema.TrainTripTrait, s []string) {
		x.ArrivalPlatform = s[0]
	}


	basicTrainTripTraitMapping["http://schema.org/trainName"] = func(x *schema.TrainTripTrait, s []string) {
		x.TrainName = s[0]
	}


	basicTrainTripTraitMapping["http://schema.org/trainNumber"] = func(x *schema.TrainTripTrait, s []string) {
		x.TrainNumber = s[0]
	}




	customTrainTripTraitMapping["http://schema.org/arrivalStation"] = func(x *schema.TrainTripTrait, ctx jsonld.Context, s string){
		var y schema.TrainStation
		if strings.HasPrefix(s, "_:") {
			y = NewTrainStationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewTrainStation()
			y.Id = s
		}

		x.ArrivalStation = &y

		return
	}

	customTrainTripTraitMapping["http://schema.org/departureStation"] = func(x *schema.TrainTripTrait, ctx jsonld.Context, s string){
		var y schema.TrainStation
		if strings.HasPrefix(s, "_:") {
			y = NewTrainStationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewTrainStation()
			y.Id = s
		}

		x.DepartureStation = &y

		return
	}
}

func NewTrainTripFromContext(ctx jsonld.Context) (x schema.TrainTrip) {
	x.Type = "http://schema.org/TrainTrip"
	MapBasicToTripTrait(ctx, &x.TripTrait)
	MapCustomToTripTrait(ctx, &x.TripTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTrainTripTrait(ctx, &x.TrainTripTrait)
	MapCustomToTrainTripTrait(ctx, &x.TrainTripTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTrainTripTrait(ctx jsonld.Context, x *schema.TrainTripTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTrainTripTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTrainTripTrait(ctx jsonld.Context, x *schema.TrainTripTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTrainTripTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}