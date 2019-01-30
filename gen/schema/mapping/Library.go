package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLibrary strings.Replacer
var strconvLibrary strconv.NumError

var basicLibraryTraitMapping = map[string]func(*schema.LibraryTrait, []string){}
var customLibraryTraitMapping = map[string]func(*schema.LibraryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Library", func(ctx jsonld.Context) (interface{}) {
		return NewLibraryFromContext(ctx)
	})

}

func NewLibraryFromContext(ctx jsonld.Context) (x schema.Library) {
	x.Type = "http://schema.org/Library"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToLibraryTrait(ctx, &x.LibraryTrait)
	MapCustomToLibraryTrait(ctx, &x.LibraryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLibraryTrait(ctx jsonld.Context, x *schema.LibraryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLibraryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLibraryTrait(ctx jsonld.Context, x *schema.LibraryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLibraryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}