package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFilmAction strings.Replacer
var strconvFilmAction strconv.NumError

var basicFilmActionTraitMapping = map[string]func(*schema.FilmActionTrait, []string){}
var customFilmActionTraitMapping = map[string]func(*schema.FilmActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FilmAction", func(ctx jsonld.Context) (interface{}) {
		return NewFilmActionFromContext(ctx)
	})

}

func NewFilmActionFromContext(ctx jsonld.Context) (x schema.FilmAction) {
	x.Type = "http://schema.org/FilmAction"
	MapBasicToCreateActionTrait(ctx, &x.CreateActionTrait)
	MapCustomToCreateActionTrait(ctx, &x.CreateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFilmActionTrait(ctx, &x.FilmActionTrait)
	MapCustomToFilmActionTrait(ctx, &x.FilmActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFilmActionTrait(ctx jsonld.Context, x *schema.FilmActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFilmActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFilmActionTrait(ctx jsonld.Context, x *schema.FilmActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFilmActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}