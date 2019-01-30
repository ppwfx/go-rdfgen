package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPostOffice strings.Replacer
var strconvPostOffice strconv.NumError

var basicPostOfficeTraitMapping = map[string]func(*schema.PostOfficeTrait, []string){}
var customPostOfficeTraitMapping = map[string]func(*schema.PostOfficeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PostOffice", func(ctx jsonld.Context) (interface{}) {
		return NewPostOfficeFromContext(ctx)
	})

}

func NewPostOfficeFromContext(ctx jsonld.Context) (x schema.PostOffice) {
	x.Type = "http://schema.org/PostOffice"
	MapBasicToGovernmentOfficeTrait(ctx, &x.GovernmentOfficeTrait)
	MapCustomToGovernmentOfficeTrait(ctx, &x.GovernmentOfficeTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToPostOfficeTrait(ctx, &x.PostOfficeTrait)
	MapCustomToPostOfficeTrait(ctx, &x.PostOfficeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPostOfficeTrait(ctx jsonld.Context, x *schema.PostOfficeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPostOfficeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPostOfficeTrait(ctx jsonld.Context, x *schema.PostOfficeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPostOfficeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}