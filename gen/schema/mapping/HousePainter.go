package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHousePainter strings.Replacer
var strconvHousePainter strconv.NumError

var basicHousePainterTraitMapping = map[string]func(*schema.HousePainterTrait, []string){}
var customHousePainterTraitMapping = map[string]func(*schema.HousePainterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HousePainter", func(ctx jsonld.Context) (interface{}) {
		return NewHousePainterFromContext(ctx)
	})

}

func NewHousePainterFromContext(ctx jsonld.Context) (x schema.HousePainter) {
	x.Type = "http://schema.org/HousePainter"
	MapBasicToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)
	MapCustomToHomeAndConstructionBusinessTrait(ctx, &x.HomeAndConstructionBusinessTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToHousePainterTrait(ctx, &x.HousePainterTrait)
	MapCustomToHousePainterTrait(ctx, &x.HousePainterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHousePainterTrait(ctx jsonld.Context, x *schema.HousePainterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHousePainterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHousePainterTrait(ctx jsonld.Context, x *schema.HousePainterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHousePainterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}