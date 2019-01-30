package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTouristDestination strings.Replacer
var strconvTouristDestination strconv.NumError

var basicTouristDestinationTraitMapping = map[string]func(*schema.TouristDestinationTrait, []string){}
var customTouristDestinationTraitMapping = map[string]func(*schema.TouristDestinationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TouristDestination", func(ctx jsonld.Context) (interface{}) {
		return NewTouristDestinationFromContext(ctx)
	})




	customTouristDestinationTraitMapping["http://schema.org/touristType"] = func(x *schema.TouristDestinationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TouristType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TouristType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TouristType = s
		}
	}

	customTouristDestinationTraitMapping["http://schema.org/includesAttraction"] = func(x *schema.TouristDestinationTrait, ctx jsonld.Context, s string){
		var y schema.TouristAttraction
		if strings.HasPrefix(s, "_:") {
			y = NewTouristAttractionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewTouristAttraction()
			y.Id = s
		}

		x.IncludesAttraction = &y

		return
	}
}

func NewTouristDestinationFromContext(ctx jsonld.Context) (x schema.TouristDestination) {
	x.Type = "http://schema.org/TouristDestination"
	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTouristDestinationTrait(ctx, &x.TouristDestinationTrait)
	MapCustomToTouristDestinationTrait(ctx, &x.TouristDestinationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTouristDestinationTrait(ctx jsonld.Context, x *schema.TouristDestinationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTouristDestinationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTouristDestinationTrait(ctx jsonld.Context, x *schema.TouristDestinationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTouristDestinationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}