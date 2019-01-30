package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPhotograph strings.Replacer
var strconvPhotograph strconv.NumError

var basicPhotographTraitMapping = map[string]func(*schema.PhotographTrait, []string){}
var customPhotographTraitMapping = map[string]func(*schema.PhotographTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Photograph", func(ctx jsonld.Context) (interface{}) {
		return NewPhotographFromContext(ctx)
	})

}

func NewPhotographFromContext(ctx jsonld.Context) (x schema.Photograph) {
	x.Type = "http://schema.org/Photograph"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPhotographTrait(ctx, &x.PhotographTrait)
	MapCustomToPhotographTrait(ctx, &x.PhotographTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPhotographTrait(ctx jsonld.Context, x *schema.PhotographTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPhotographTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPhotographTrait(ctx jsonld.Context, x *schema.PhotographTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPhotographTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}