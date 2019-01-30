package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTVClip strings.Replacer
var strconvTVClip strconv.NumError

var basicTVClipTraitMapping = map[string]func(*schema.TVClipTrait, []string){}
var customTVClipTraitMapping = map[string]func(*schema.TVClipTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TVClip", func(ctx jsonld.Context) (interface{}) {
		return NewTVClipFromContext(ctx)
	})



	customTVClipTraitMapping["http://schema.org/partOfTVSeries"] = func(x *schema.TVClipTrait, ctx jsonld.Context, s string){
		var y schema.TVSeries
		if strings.HasPrefix(s, "_:") {
			y = NewTVSeriesFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewTVSeries()
			y.Id = s
		}

		x.PartOfTVSeries = &y

		return
	}
}

func NewTVClipFromContext(ctx jsonld.Context) (x schema.TVClip) {
	x.Type = "http://schema.org/TVClip"
	MapBasicToClipTrait(ctx, &x.ClipTrait)
	MapCustomToClipTrait(ctx, &x.ClipTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTVClipTrait(ctx, &x.TVClipTrait)
	MapCustomToTVClipTrait(ctx, &x.TVClipTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTVClipTrait(ctx jsonld.Context, x *schema.TVClipTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTVClipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTVClipTrait(ctx jsonld.Context, x *schema.TVClipTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTVClipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}