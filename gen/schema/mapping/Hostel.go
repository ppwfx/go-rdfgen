package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHostel strings.Replacer
var strconvHostel strconv.NumError

var basicHostelTraitMapping = map[string]func(*schema.HostelTrait, []string){}
var customHostelTraitMapping = map[string]func(*schema.HostelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Hostel", func(ctx jsonld.Context) (interface{}) {
		return NewHostelFromContext(ctx)
	})

}

func NewHostelFromContext(ctx jsonld.Context) (x schema.Hostel) {
	x.Type = "http://schema.org/Hostel"
	MapBasicToLodgingBusinessTrait(ctx, &x.LodgingBusinessTrait)
	MapCustomToLodgingBusinessTrait(ctx, &x.LodgingBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToHostelTrait(ctx, &x.HostelTrait)
	MapCustomToHostelTrait(ctx, &x.HostelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHostelTrait(ctx jsonld.Context, x *schema.HostelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHostelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHostelTrait(ctx jsonld.Context, x *schema.HostelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHostelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}