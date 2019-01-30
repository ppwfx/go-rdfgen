package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCompleteDataFeed strings.Replacer
var strconvCompleteDataFeed strconv.NumError

var basicCompleteDataFeedTraitMapping = map[string]func(*schema.CompleteDataFeedTrait, []string){}
var customCompleteDataFeedTraitMapping = map[string]func(*schema.CompleteDataFeedTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CompleteDataFeed", func(ctx jsonld.Context) (interface{}) {
		return NewCompleteDataFeedFromContext(ctx)
	})

}

func NewCompleteDataFeedFromContext(ctx jsonld.Context) (x schema.CompleteDataFeed) {
	x.Type = "http://schema.org/CompleteDataFeed"
	MapBasicToDataFeedTrait(ctx, &x.DataFeedTrait)
	MapCustomToDataFeedTrait(ctx, &x.DataFeedTrait)

	MapBasicToDatasetTrait(ctx, &x.DatasetTrait)
	MapCustomToDatasetTrait(ctx, &x.DatasetTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCompleteDataFeedTrait(ctx, &x.CompleteDataFeedTrait)
	MapCustomToCompleteDataFeedTrait(ctx, &x.CompleteDataFeedTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCompleteDataFeedTrait(ctx jsonld.Context, x *schema.CompleteDataFeedTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCompleteDataFeedTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCompleteDataFeedTrait(ctx jsonld.Context, x *schema.CompleteDataFeedTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCompleteDataFeedTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}