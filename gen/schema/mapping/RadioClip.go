package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRadioClip strings.Replacer
var strconvRadioClip strconv.NumError

var basicRadioClipTraitMapping = map[string]func(*schema.RadioClipTrait, []string){}
var customRadioClipTraitMapping = map[string]func(*schema.RadioClipTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RadioClip", func(ctx jsonld.Context) (interface{}) {
		return NewRadioClipFromContext(ctx)
	})

}

func NewRadioClipFromContext(ctx jsonld.Context) (x schema.RadioClip) {
	x.Type = "http://schema.org/RadioClip"
	MapBasicToClipTrait(ctx, &x.ClipTrait)
	MapCustomToClipTrait(ctx, &x.ClipTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRadioClipTrait(ctx, &x.RadioClipTrait)
	MapCustomToRadioClipTrait(ctx, &x.RadioClipTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRadioClipTrait(ctx jsonld.Context, x *schema.RadioClipTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRadioClipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRadioClipTrait(ctx jsonld.Context, x *schema.RadioClipTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRadioClipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}