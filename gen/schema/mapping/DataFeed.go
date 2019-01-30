package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDataFeed strings.Replacer
var strconvDataFeed strconv.NumError

var basicDataFeedTraitMapping = map[string]func(*schema.DataFeedTrait, []string){}
var customDataFeedTraitMapping = map[string]func(*schema.DataFeedTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DataFeed", func(ctx jsonld.Context) (interface{}) {
		return NewDataFeedFromContext(ctx)
	})



	customDataFeedTraitMapping["http://schema.org/dataFeedElement"] = func(x *schema.DataFeedTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DataFeedElement, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DataFeedElement = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DataFeedElement = s
		}
	}
}

func NewDataFeedFromContext(ctx jsonld.Context) (x schema.DataFeed) {
	x.Type = "http://schema.org/DataFeed"
	MapBasicToDatasetTrait(ctx, &x.DatasetTrait)
	MapCustomToDatasetTrait(ctx, &x.DatasetTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDataFeedTrait(ctx, &x.DataFeedTrait)
	MapCustomToDataFeedTrait(ctx, &x.DataFeedTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDataFeedTrait(ctx jsonld.Context, x *schema.DataFeedTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDataFeedTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDataFeedTrait(ctx jsonld.Context, x *schema.DataFeedTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDataFeedTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}