package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMap strings.Replacer
var strconvMap strconv.NumError

var basicMapTraitMapping = map[string]func(*schema.MapTrait, []string){}
var customMapTraitMapping = map[string]func(*schema.MapTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Map", func(ctx jsonld.Context) (interface{}) {
		return NewMapFromContext(ctx)
	})



	customMapTraitMapping["http://schema.org/mapType"] = func(x *schema.MapTrait, ctx jsonld.Context, s string){
		var y schema.MapCategoryType
		if strings.HasPrefix(s, "_:") {
			y = NewMapCategoryTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMapCategoryType()
			y.Id = s
		}

		x.MapType = &y

		return
	}
}

func NewMapFromContext(ctx jsonld.Context) (x schema.Map) {
	x.Type = "http://schema.org/Map"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMapTrait(ctx, &x.MapTrait)
	MapCustomToMapTrait(ctx, &x.MapTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMapTrait(ctx jsonld.Context, x *schema.MapTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMapTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMapTrait(ctx jsonld.Context, x *schema.MapTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMapTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}