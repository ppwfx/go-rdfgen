package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGeospatialGeometry strings.Replacer
var strconvGeospatialGeometry strconv.NumError

var basicGeospatialGeometryTraitMapping = map[string]func(*schema.GeospatialGeometryTrait, []string){}
var customGeospatialGeometryTraitMapping = map[string]func(*schema.GeospatialGeometryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GeospatialGeometry", func(ctx jsonld.Context) (interface{}) {
		return NewGeospatialGeometryFromContext(ctx)
	})












	customGeospatialGeometryTraitMapping["http://schema.org/geospatiallyContains"] = func(x *schema.GeospatialGeometryTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyContains, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyContains = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyContains = s
		}
	}

	customGeospatialGeometryTraitMapping["http://schema.org/geospatiallyCoveredBy"] = func(x *schema.GeospatialGeometryTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyCoveredBy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyCoveredBy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyCoveredBy = s
		}
	}

	customGeospatialGeometryTraitMapping["http://schema.org/geospatiallyCovers"] = func(x *schema.GeospatialGeometryTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyCovers, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyCovers = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyCovers = s
		}
	}

	customGeospatialGeometryTraitMapping["http://schema.org/geospatiallyCrosses"] = func(x *schema.GeospatialGeometryTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyCrosses, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyCrosses = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyCrosses = s
		}
	}

	customGeospatialGeometryTraitMapping["http://schema.org/geospatiallyDisjoint"] = func(x *schema.GeospatialGeometryTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyDisjoint, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyDisjoint = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyDisjoint = s
		}
	}

	customGeospatialGeometryTraitMapping["http://schema.org/geospatiallyEquals"] = func(x *schema.GeospatialGeometryTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyEquals, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyEquals = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyEquals = s
		}
	}

	customGeospatialGeometryTraitMapping["http://schema.org/geospatiallyIntersects"] = func(x *schema.GeospatialGeometryTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyIntersects, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyIntersects = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyIntersects = s
		}
	}

	customGeospatialGeometryTraitMapping["http://schema.org/geospatiallyOverlaps"] = func(x *schema.GeospatialGeometryTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyOverlaps, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyOverlaps = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyOverlaps = s
		}
	}

	customGeospatialGeometryTraitMapping["http://schema.org/geospatiallyTouches"] = func(x *schema.GeospatialGeometryTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyTouches, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyTouches = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyTouches = s
		}
	}

	customGeospatialGeometryTraitMapping["http://schema.org/geospatiallyWithin"] = func(x *schema.GeospatialGeometryTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.GeospatiallyWithin, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.GeospatiallyWithin = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.GeospatiallyWithin = s
		}
	}
}

func NewGeospatialGeometryFromContext(ctx jsonld.Context) (x schema.GeospatialGeometry) {
	x.Type = "http://schema.org/GeospatialGeometry"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGeospatialGeometryTrait(ctx, &x.GeospatialGeometryTrait)
	MapCustomToGeospatialGeometryTrait(ctx, &x.GeospatialGeometryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGeospatialGeometryTrait(ctx jsonld.Context, x *schema.GeospatialGeometryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGeospatialGeometryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGeospatialGeometryTrait(ctx jsonld.Context, x *schema.GeospatialGeometryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGeospatialGeometryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}