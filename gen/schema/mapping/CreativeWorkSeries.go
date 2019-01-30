package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCreativeWorkSeries strings.Replacer
var strconvCreativeWorkSeries strconv.NumError

var basicCreativeWorkSeriesTraitMapping = map[string]func(*schema.CreativeWorkSeriesTrait, []string){}
var customCreativeWorkSeriesTraitMapping = map[string]func(*schema.CreativeWorkSeriesTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CreativeWorkSeries", func(ctx jsonld.Context) (interface{}) {
		return NewCreativeWorkSeriesFromContext(ctx)
	})




	basicCreativeWorkSeriesTraitMapping["http://schema.org/issn"] = func(x *schema.CreativeWorkSeriesTrait, s []string) {
		x.Issn = s[0]
	}


	customCreativeWorkSeriesTraitMapping["http://schema.org/endDate"] = func(x *schema.CreativeWorkSeriesTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EndDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EndDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EndDate = s
		}
	}

	customCreativeWorkSeriesTraitMapping["http://schema.org/startDate"] = func(x *schema.CreativeWorkSeriesTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.StartDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.StartDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.StartDate = s
		}
	}
}

func NewCreativeWorkSeriesFromContext(ctx jsonld.Context) (x schema.CreativeWorkSeries) {
	x.Type = "http://schema.org/CreativeWorkSeries"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToSeriesTrait(ctx, &x.SeriesTrait)
	MapCustomToSeriesTrait(ctx, &x.SeriesTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)


	MapBasicToCreativeWorkSeriesTrait(ctx, &x.CreativeWorkSeriesTrait)
	MapCustomToCreativeWorkSeriesTrait(ctx, &x.CreativeWorkSeriesTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCreativeWorkSeriesTrait(ctx jsonld.Context, x *schema.CreativeWorkSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCreativeWorkSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCreativeWorkSeriesTrait(ctx jsonld.Context, x *schema.CreativeWorkSeriesTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCreativeWorkSeriesTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}