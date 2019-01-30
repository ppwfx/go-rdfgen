package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDataset strings.Replacer
var strconvDataset strconv.NumError

var basicDatasetTraitMapping = map[string]func(*schema.DatasetTrait, []string){}
var customDatasetTraitMapping = map[string]func(*schema.DatasetTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Dataset", func(ctx jsonld.Context) (interface{}) {
		return NewDatasetFromContext(ctx)
	})



	basicDatasetTraitMapping["http://schema.org/issn"] = func(x *schema.DatasetTrait, s []string) {
		x.Issn = s[0]
	}










	customDatasetTraitMapping["http://schema.org/measurementTechnique"] = func(x *schema.DatasetTrait, ctx jsonld.Context, s string){
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

	customDatasetTraitMapping["http://schema.org/datasetTimeInterval"] = func(x *schema.DatasetTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.DatasetTimeInterval = &y

		return
	}

	customDatasetTraitMapping["http://schema.org/temporal"] = func(x *schema.DatasetTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.Temporal = &y

		return
	}

	customDatasetTraitMapping["http://schema.org/variableMeasured"] = func(x *schema.DatasetTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.VariableMeasured, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.VariableMeasured = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.VariableMeasured = s
		}
	}

	customDatasetTraitMapping["http://schema.org/catalog"] = func(x *schema.DatasetTrait, ctx jsonld.Context, s string){
		var y schema.DataCatalog
		if strings.HasPrefix(s, "_:") {
			y = NewDataCatalogFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDataCatalog()
			y.Id = s
		}

		x.Catalog = &y

		return
	}

	customDatasetTraitMapping["http://schema.org/distribution"] = func(x *schema.DatasetTrait, ctx jsonld.Context, s string){
		var y schema.DataDownload
		if strings.HasPrefix(s, "_:") {
			y = NewDataDownloadFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDataDownload()
			y.Id = s
		}

		x.Distribution = &y

		return
	}

	customDatasetTraitMapping["http://schema.org/includedDataCatalog"] = func(x *schema.DatasetTrait, ctx jsonld.Context, s string){
		var y schema.DataCatalog
		if strings.HasPrefix(s, "_:") {
			y = NewDataCatalogFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDataCatalog()
			y.Id = s
		}

		x.IncludedDataCatalog = &y

		return
	}

	customDatasetTraitMapping["http://schema.org/includedInDataCatalog"] = func(x *schema.DatasetTrait, ctx jsonld.Context, s string){
		var y schema.DataCatalog
		if strings.HasPrefix(s, "_:") {
			y = NewDataCatalogFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDataCatalog()
			y.Id = s
		}

		x.IncludedInDataCatalog = &y

		return
	}

	customDatasetTraitMapping["http://schema.org/spatial"] = func(x *schema.DatasetTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.Spatial = &y

		return
	}
}

func NewDatasetFromContext(ctx jsonld.Context) (x schema.Dataset) {
	x.Type = "http://schema.org/Dataset"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDatasetTrait(ctx, &x.DatasetTrait)
	MapCustomToDatasetTrait(ctx, &x.DatasetTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDatasetTrait(ctx jsonld.Context, x *schema.DatasetTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDatasetTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDatasetTrait(ctx jsonld.Context, x *schema.DatasetTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDatasetTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}