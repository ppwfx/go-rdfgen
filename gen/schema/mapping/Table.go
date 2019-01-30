package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTable strings.Replacer
var strconvTable strconv.NumError

var basicTableTraitMapping = map[string]func(*schema.TableTrait, []string){}
var customTableTraitMapping = map[string]func(*schema.TableTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Table", func(ctx jsonld.Context) (interface{}) {
		return NewTableFromContext(ctx)
	})

}

func NewTableFromContext(ctx jsonld.Context) (x schema.Table) {
	x.Type = "http://schema.org/Table"
	MapBasicToWebPageElementTrait(ctx, &x.WebPageElementTrait)
	MapCustomToWebPageElementTrait(ctx, &x.WebPageElementTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTableTrait(ctx, &x.TableTrait)
	MapCustomToTableTrait(ctx, &x.TableTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTableTrait(ctx jsonld.Context, x *schema.TableTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTableTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTableTrait(ctx jsonld.Context, x *schema.TableTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTableTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}