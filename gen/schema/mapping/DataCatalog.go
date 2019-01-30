package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDataCatalog strings.Replacer
var strconvDataCatalog strconv.NumError

var basicDataCatalogTraitMapping = map[string]func(*schema.DataCatalogTrait, []string){}
var customDataCatalogTraitMapping = map[string]func(*schema.DataCatalogTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DataCatalog", func(ctx jsonld.Context) (interface{}) {
		return NewDataCatalogFromContext(ctx)
	})




	customDataCatalogTraitMapping["http://schema.org/measurementTechnique"] = func(x *schema.DataCatalogTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MeasurementTechnique, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MeasurementTechnique = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MeasurementTechnique = s
		}
	}

	customDataCatalogTraitMapping["http://schema.org/dataset"] = func(x *schema.DataCatalogTrait, ctx jsonld.Context, s string){
		var y schema.Dataset
		if strings.HasPrefix(s, "_:") {
			y = NewDatasetFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDataset()
			y.Id = s
		}

		x.Dataset = &y

		return
	}
}

func NewDataCatalogFromContext(ctx jsonld.Context) (x schema.DataCatalog) {
	x.Type = "http://schema.org/DataCatalog"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDataCatalogTrait(ctx, &x.DataCatalogTrait)
	MapCustomToDataCatalogTrait(ctx, &x.DataCatalogTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDataCatalogTrait(ctx jsonld.Context, x *schema.DataCatalogTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDataCatalogTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDataCatalogTrait(ctx jsonld.Context, x *schema.DataCatalogTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDataCatalogTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}