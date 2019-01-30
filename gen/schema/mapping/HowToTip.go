package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHowToTip strings.Replacer
var strconvHowToTip strconv.NumError

var basicHowToTipTraitMapping = map[string]func(*schema.HowToTipTrait, []string){}
var customHowToTipTraitMapping = map[string]func(*schema.HowToTipTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HowToTip", func(ctx jsonld.Context) (interface{}) {
		return NewHowToTipFromContext(ctx)
	})

}

func NewHowToTipFromContext(ctx jsonld.Context) (x schema.HowToTip) {
	x.Type = "http://schema.org/HowToTip"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToListItemTrait(ctx, &x.ListItemTrait)
	MapCustomToListItemTrait(ctx, &x.ListItemTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)


	MapBasicToHowToTipTrait(ctx, &x.HowToTipTrait)
	MapCustomToHowToTipTrait(ctx, &x.HowToTipTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHowToTipTrait(ctx jsonld.Context, x *schema.HowToTipTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHowToTipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHowToTipTrait(ctx jsonld.Context, x *schema.HowToTipTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHowToTipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}