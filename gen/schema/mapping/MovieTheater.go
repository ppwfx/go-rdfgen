package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMovieTheater strings.Replacer
var strconvMovieTheater strconv.NumError

var basicMovieTheaterTraitMapping = map[string]func(*schema.MovieTheaterTrait, []string){}
var customMovieTheaterTraitMapping = map[string]func(*schema.MovieTheaterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MovieTheater", func(ctx jsonld.Context) (interface{}) {
		return NewMovieTheaterFromContext(ctx)
	})


	basicMovieTheaterTraitMapping["http://schema.org/screenCount"] = func(x *schema.MovieTheaterTrait, s []string) {
		var err error
		x.ScreenCount, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}

}

func NewMovieTheaterFromContext(ctx jsonld.Context) (x schema.MovieTheater) {
	x.Type = "http://schema.org/MovieTheater"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToEntertainmentBusinessTrait(ctx, &x.EntertainmentBusinessTrait)
	MapCustomToEntertainmentBusinessTrait(ctx, &x.EntertainmentBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToMovieTheaterTrait(ctx, &x.MovieTheaterTrait)
	MapCustomToMovieTheaterTrait(ctx, &x.MovieTheaterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMovieTheaterTrait(ctx jsonld.Context, x *schema.MovieTheaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMovieTheaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMovieTheaterTrait(ctx jsonld.Context, x *schema.MovieTheaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMovieTheaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}