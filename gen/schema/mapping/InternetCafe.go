package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInternetCafe strings.Replacer
var strconvInternetCafe strconv.NumError

var basicInternetCafeTraitMapping = map[string]func(*schema.InternetCafeTrait, []string){}
var customInternetCafeTraitMapping = map[string]func(*schema.InternetCafeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InternetCafe", func(ctx jsonld.Context) (interface{}) {
		return NewInternetCafeFromContext(ctx)
	})

}

func NewInternetCafeFromContext(ctx jsonld.Context) (x schema.InternetCafe) {
	x.Type = "http://schema.org/InternetCafe"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToInternetCafeTrait(ctx, &x.InternetCafeTrait)
	MapCustomToInternetCafeTrait(ctx, &x.InternetCafeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInternetCafeTrait(ctx jsonld.Context, x *schema.InternetCafeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInternetCafeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInternetCafeTrait(ctx jsonld.Context, x *schema.InternetCafeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInternetCafeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}