package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHowToTool strings.Replacer
var strconvHowToTool strconv.NumError

var basicHowToToolTraitMapping = map[string]func(*schema.HowToToolTrait, []string){}
var customHowToToolTraitMapping = map[string]func(*schema.HowToToolTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HowToTool", func(ctx jsonld.Context) (interface{}) {
		return NewHowToToolFromContext(ctx)
	})

}

func NewHowToToolFromContext(ctx jsonld.Context) (x schema.HowToTool) {
	x.Type = "http://schema.org/HowToTool"
	MapBasicToHowToItemTrait(ctx, &x.HowToItemTrait)
	MapCustomToHowToItemTrait(ctx, &x.HowToItemTrait)

	MapBasicToListItemTrait(ctx, &x.ListItemTrait)
	MapCustomToListItemTrait(ctx, &x.ListItemTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHowToToolTrait(ctx, &x.HowToToolTrait)
	MapCustomToHowToToolTrait(ctx, &x.HowToToolTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHowToToolTrait(ctx jsonld.Context, x *schema.HowToToolTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHowToToolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHowToToolTrait(ctx jsonld.Context, x *schema.HowToToolTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHowToToolTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}