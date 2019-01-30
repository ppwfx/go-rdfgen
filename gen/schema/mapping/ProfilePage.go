package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsProfilePage strings.Replacer
var strconvProfilePage strconv.NumError

var basicProfilePageTraitMapping = map[string]func(*schema.ProfilePageTrait, []string){}
var customProfilePageTraitMapping = map[string]func(*schema.ProfilePageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ProfilePage", func(ctx jsonld.Context) (interface{}) {
		return NewProfilePageFromContext(ctx)
	})

}

func NewProfilePageFromContext(ctx jsonld.Context) (x schema.ProfilePage) {
	x.Type = "http://schema.org/ProfilePage"
	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToProfilePageTrait(ctx, &x.ProfilePageTrait)
	MapCustomToProfilePageTrait(ctx, &x.ProfilePageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToProfilePageTrait(ctx jsonld.Context, x *schema.ProfilePageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicProfilePageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToProfilePageTrait(ctx jsonld.Context, x *schema.ProfilePageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customProfilePageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}