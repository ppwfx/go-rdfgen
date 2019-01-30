package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSearchAction strings.Replacer
var strconvSearchAction strconv.NumError

var basicSearchActionTraitMapping = map[string]func(*schema.SearchActionTrait, []string){}
var customSearchActionTraitMapping = map[string]func(*schema.SearchActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SearchAction", func(ctx jsonld.Context) (interface{}) {
		return NewSearchActionFromContext(ctx)
	})


	basicSearchActionTraitMapping["http://schema.org/query"] = func(x *schema.SearchActionTrait, s []string) {
		x.Query = s[0]
	}

}

func NewSearchActionFromContext(ctx jsonld.Context) (x schema.SearchAction) {
	x.Type = "http://schema.org/SearchAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSearchActionTrait(ctx, &x.SearchActionTrait)
	MapCustomToSearchActionTrait(ctx, &x.SearchActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSearchActionTrait(ctx jsonld.Context, x *schema.SearchActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSearchActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSearchActionTrait(ctx jsonld.Context, x *schema.SearchActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSearchActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}