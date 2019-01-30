package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCollection strings.Replacer
var strconvCollection strconv.NumError

var basicCollectionTraitMapping = map[string]func(*schema.CollectionTrait, []string){}
var customCollectionTraitMapping = map[string]func(*schema.CollectionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Collection", func(ctx jsonld.Context) (interface{}) {
		return NewCollectionFromContext(ctx)
	})

}

func NewCollectionFromContext(ctx jsonld.Context) (x schema.Collection) {
	x.Type = "http://schema.org/Collection"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCollectionTrait(ctx, &x.CollectionTrait)
	MapCustomToCollectionTrait(ctx, &x.CollectionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCollectionTrait(ctx jsonld.Context, x *schema.CollectionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCollectionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCollectionTrait(ctx jsonld.Context, x *schema.CollectionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCollectionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}