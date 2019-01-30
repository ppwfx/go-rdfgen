package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBookmarkAction strings.Replacer
var strconvBookmarkAction strconv.NumError

var basicBookmarkActionTraitMapping = map[string]func(*schema.BookmarkActionTrait, []string){}
var customBookmarkActionTraitMapping = map[string]func(*schema.BookmarkActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BookmarkAction", func(ctx jsonld.Context) (interface{}) {
		return NewBookmarkActionFromContext(ctx)
	})

}

func NewBookmarkActionFromContext(ctx jsonld.Context) (x schema.BookmarkAction) {
	x.Type = "http://schema.org/BookmarkAction"
	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBookmarkActionTrait(ctx, &x.BookmarkActionTrait)
	MapCustomToBookmarkActionTrait(ctx, &x.BookmarkActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBookmarkActionTrait(ctx jsonld.Context, x *schema.BookmarkActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBookmarkActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBookmarkActionTrait(ctx jsonld.Context, x *schema.BookmarkActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBookmarkActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}