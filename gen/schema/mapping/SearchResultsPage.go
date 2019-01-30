package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSearchResultsPage strings.Replacer
var strconvSearchResultsPage strconv.NumError

var basicSearchResultsPageTraitMapping = map[string]func(*schema.SearchResultsPageTrait, []string){}
var customSearchResultsPageTraitMapping = map[string]func(*schema.SearchResultsPageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SearchResultsPage", func(ctx jsonld.Context) (interface{}) {
		return NewSearchResultsPageFromContext(ctx)
	})

}

func NewSearchResultsPageFromContext(ctx jsonld.Context) (x schema.SearchResultsPage) {
	x.Type = "http://schema.org/SearchResultsPage"
	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSearchResultsPageTrait(ctx, &x.SearchResultsPageTrait)
	MapCustomToSearchResultsPageTrait(ctx, &x.SearchResultsPageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSearchResultsPageTrait(ctx jsonld.Context, x *schema.SearchResultsPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSearchResultsPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSearchResultsPageTrait(ctx jsonld.Context, x *schema.SearchResultsPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSearchResultsPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}