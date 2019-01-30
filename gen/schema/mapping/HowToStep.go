package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHowToStep strings.Replacer
var strconvHowToStep strconv.NumError

var basicHowToStepTraitMapping = map[string]func(*schema.HowToStepTrait, []string){}
var customHowToStepTraitMapping = map[string]func(*schema.HowToStepTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HowToStep", func(ctx jsonld.Context) (interface{}) {
		return NewHowToStepFromContext(ctx)
	})

}

func NewHowToStepFromContext(ctx jsonld.Context) (x schema.HowToStep) {
	x.Type = "http://schema.org/HowToStep"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToListItemTrait(ctx, &x.ListItemTrait)
	MapCustomToListItemTrait(ctx, &x.ListItemTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToItemListTrait(ctx, &x.ItemListTrait)
	MapCustomToItemListTrait(ctx, &x.ItemListTrait)


	MapBasicToHowToStepTrait(ctx, &x.HowToStepTrait)
	MapCustomToHowToStepTrait(ctx, &x.HowToStepTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHowToStepTrait(ctx jsonld.Context, x *schema.HowToStepTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHowToStepTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHowToStepTrait(ctx jsonld.Context, x *schema.HowToStepTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHowToStepTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}