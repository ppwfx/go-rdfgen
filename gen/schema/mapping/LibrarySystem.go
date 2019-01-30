package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLibrarySystem strings.Replacer
var strconvLibrarySystem strconv.NumError

var basicLibrarySystemTraitMapping = map[string]func(*schema.LibrarySystemTrait, []string){}
var customLibrarySystemTraitMapping = map[string]func(*schema.LibrarySystemTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LibrarySystem", func(ctx jsonld.Context) (interface{}) {
		return NewLibrarySystemFromContext(ctx)
	})

}

func NewLibrarySystemFromContext(ctx jsonld.Context) (x schema.LibrarySystem) {
	x.Type = "http://schema.org/LibrarySystem"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLibrarySystemTrait(ctx, &x.LibrarySystemTrait)
	MapCustomToLibrarySystemTrait(ctx, &x.LibrarySystemTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLibrarySystemTrait(ctx jsonld.Context, x *schema.LibrarySystemTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLibrarySystemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLibrarySystemTrait(ctx jsonld.Context, x *schema.LibrarySystemTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLibrarySystemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}