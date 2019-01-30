package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGeoCoordinates strings.Replacer
var strconvGeoCoordinates strconv.NumError

var basicGeoCoordinatesTraitMapping = map[string]func(*schema.GeoCoordinatesTrait, []string){}
var customGeoCoordinatesTraitMapping = map[string]func(*schema.GeoCoordinatesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GeoCoordinates", func(ctx jsonld.Context) (interface{}) {
		return NewGeoCoordinatesFromContext(ctx)
	})







	basicGeoCoordinatesTraitMapping["http://schema.org/postalCode"] = func(x *schema.GeoCoordinatesTrait, s []string) {
		x.PostalCode = s[0]
	}


	customGeoCoordinatesTraitMapping["http://schema.org/address"] = func(x *schema.GeoCoordinatesTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Address, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Address = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Address = s
		}
	}

	customGeoCoordinatesTraitMapping["http://schema.org/addressCountry"] = func(x *schema.GeoCoordinatesTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AddressCountry, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AddressCountry = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AddressCountry = s
		}
	}

	customGeoCoordinatesTraitMapping["http://schema.org/elevation"] = func(x *schema.GeoCoordinatesTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Elevation, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Elevation = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Elevation = s
		}
	}

	customGeoCoordinatesTraitMapping["http://schema.org/latitude"] = func(x *schema.GeoCoordinatesTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Latitude, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Latitude = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Latitude = s
		}
	}

	customGeoCoordinatesTraitMapping["http://schema.org/longitude"] = func(x *schema.GeoCoordinatesTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Longitude, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Longitude = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Longitude = s
		}
	}
}

func NewGeoCoordinatesFromContext(ctx jsonld.Context) (x schema.GeoCoordinates) {
	x.Type = "http://schema.org/GeoCoordinates"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGeoCoordinatesTrait(ctx, &x.GeoCoordinatesTrait)
	MapCustomToGeoCoordinatesTrait(ctx, &x.GeoCoordinatesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGeoCoordinatesTrait(ctx jsonld.Context, x *schema.GeoCoordinatesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGeoCoordinatesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGeoCoordinatesTrait(ctx jsonld.Context, x *schema.GeoCoordinatesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGeoCoordinatesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}