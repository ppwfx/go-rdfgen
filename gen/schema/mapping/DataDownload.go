package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDataDownload strings.Replacer
var strconvDataDownload strconv.NumError

var basicDataDownloadTraitMapping = map[string]func(*schema.DataDownloadTrait, []string){}
var customDataDownloadTraitMapping = map[string]func(*schema.DataDownloadTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DataDownload", func(ctx jsonld.Context) (interface{}) {
		return NewDataDownloadFromContext(ctx)
	})



	customDataDownloadTraitMapping["http://schema.org/measurementTechnique"] = func(x *schema.DataDownloadTrait, ctx jsonld.Context, s string){
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
}

func NewDataDownloadFromContext(ctx jsonld.Context) (x schema.DataDownload) {
	x.Type = "http://schema.org/DataDownload"
	MapBasicToMediaObjectTrait(ctx, &x.MediaObjectTrait)
	MapCustomToMediaObjectTrait(ctx, &x.MediaObjectTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDataDownloadTrait(ctx, &x.DataDownloadTrait)
	MapCustomToDataDownloadTrait(ctx, &x.DataDownloadTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDataDownloadTrait(ctx jsonld.Context, x *schema.DataDownloadTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDataDownloadTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDataDownloadTrait(ctx jsonld.Context, x *schema.DataDownloadTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDataDownloadTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}