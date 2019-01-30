package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGeoCircle strings.Replacer
var strconvGeoCircle strconv.NumError

var basicGeoCircleTraitMapping = map[string]func(*schema.GeoCircleTrait, []string){}
var customGeoCircleTraitMapping = map[string]func(*schema.GeoCircleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GeoCircle", func(ctx jsonld.Context) (interface{}) {
		return NewGeoCircleFromContext(ctx)
	})




	customGeoCircleTraitMapping["http://schema.org/geoRadius"] = func(x *schema.GeoCircleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeoRadius, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeoRadius = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeoRadius = s
		}
	}

	customGeoCircleTraitMapping["http://schema.org/geoMidpoint"] = func(x *schema.GeoCircleTrait, ctx jsonld.Context, s string){
		var y schema.GeoCoordinates
		if strings.HasPrefix(s, "_:") {
			y = NewGeoCoordinatesFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewGeoCoordinates()
			y.Id = s
		}

		x.GeoMidpoint = &y

		return
	}
}

func NewGeoCircleFromContext(ctx jsonld.Context) (x schema.GeoCircle) {
	x.Type = "http://schema.org/GeoCircle"
	MapBasicToGeoShapeTrait(ctx, &x.GeoShapeTrait)
	MapCustomToGeoShapeTrait(ctx, &x.GeoShapeTrait)

	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGeoCircleTrait(ctx, &x.GeoCircleTrait)
	MapCustomToGeoCircleTrait(ctx, &x.GeoCircleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGeoCircleTrait(ctx jsonld.Context, x *schema.GeoCircleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGeoCircleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGeoCircleTrait(ctx jsonld.Context, x *schema.GeoCircleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGeoCircleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}