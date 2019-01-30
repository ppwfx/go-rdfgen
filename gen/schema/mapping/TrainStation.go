package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTrainStation strings.Replacer
var strconvTrainStation strconv.NumError

var basicTrainStationTraitMapping = map[string]func(*schema.TrainStationTrait, []string){}
var customTrainStationTraitMapping = map[string]func(*schema.TrainStationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TrainStation", func(ctx jsonld.Context) (interface{}) {
		return NewTrainStationFromContext(ctx)
	})

}

func NewTrainStationFromContext(ctx jsonld.Context) (x schema.TrainStation) {
	x.Type = "http://schema.org/TrainStation"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTrainStationTrait(ctx, &x.TrainStationTrait)
	MapCustomToTrainStationTrait(ctx, &x.TrainStationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTrainStationTrait(ctx jsonld.Context, x *schema.TrainStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTrainStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTrainStationTrait(ctx jsonld.Context, x *schema.TrainStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTrainStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}