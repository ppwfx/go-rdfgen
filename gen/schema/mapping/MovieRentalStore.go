package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMovieRentalStore strings.Replacer
var strconvMovieRentalStore strconv.NumError

var basicMovieRentalStoreTraitMapping = map[string]func(*schema.MovieRentalStoreTrait, []string){}
var customMovieRentalStoreTraitMapping = map[string]func(*schema.MovieRentalStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MovieRentalStore", func(ctx jsonld.Context) (interface{}) {
		return NewMovieRentalStoreFromContext(ctx)
	})

}

func NewMovieRentalStoreFromContext(ctx jsonld.Context) (x schema.MovieRentalStore) {
	x.Type = "http://schema.org/MovieRentalStore"
	MapBasicToStoreTrait(ctx, &x.StoreTrait)
	MapCustomToStoreTrait(ctx, &x.StoreTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToMovieRentalStoreTrait(ctx, &x.MovieRentalStoreTrait)
	MapCustomToMovieRentalStoreTrait(ctx, &x.MovieRentalStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMovieRentalStoreTrait(ctx jsonld.Context, x *schema.MovieRentalStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMovieRentalStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMovieRentalStoreTrait(ctx jsonld.Context, x *schema.MovieRentalStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMovieRentalStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}