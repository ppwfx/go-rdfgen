package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGeoShape strings.Replacer
var strconvGeoShape strconv.NumError

var basicGeoShapeTraitMapping = map[string]func(*schema.GeoShapeTrait, []string){}
var customGeoShapeTraitMapping = map[string]func(*schema.GeoShapeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GeoShape", func(ctx jsonld.Context) (interface{}) {
		return NewGeoShapeFromContext(ctx)
	})





	basicGeoShapeTraitMapping["http://schema.org/postalCode"] = func(x *schema.GeoShapeTrait, s []string) {
		x.PostalCode = s[0]
	}


	basicGeoShapeTraitMapping["http://schema.org/box"] = func(x *schema.GeoShapeTrait, s []string) {
		x.Box = s[0]
	}


	basicGeoShapeTraitMapping["http://schema.org/circle"] = func(x *schema.GeoShapeTrait, s []string) {
		x.Circle = s[0]
	}


	basicGeoShapeTraitMapping["http://schema.org/line"] = func(x *schema.GeoShapeTrait, s []string) {
		x.Line = s[0]
	}


	basicGeoShapeTraitMapping["http://schema.org/polygon"] = func(x *schema.GeoShapeTrait, s []string) {
		x.Polygon = s[0]
	}


	customGeoShapeTraitMapping["http://schema.org/address"] = func(x *schema.GeoShapeTrait, ctx jsonld.Context, s string){
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

	customGeoShapeTraitMapping["http://schema.org/addressCountry"] = func(x *schema.GeoShapeTrait, ctx jsonld.Context, s string){
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

	customGeoShapeTraitMapping["http://schema.org/elevation"] = func(x *schema.GeoShapeTrait, ctx jsonld.Context, s string){
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
}

func NewGeoShapeFromContext(ctx jsonld.Context) (x schema.GeoShape) {
	x.Type = "http://schema.org/GeoShape"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGeoShapeTrait(ctx, &x.GeoShapeTrait)
	MapCustomToGeoShapeTrait(ctx, &x.GeoShapeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGeoShapeTrait(ctx jsonld.Context, x *schema.GeoShapeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGeoShapeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGeoShapeTrait(ctx jsonld.Context, x *schema.GeoShapeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGeoShapeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}