package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRsvpResponseType strings.Replacer
var strconvRsvpResponseType strconv.NumError

var basicRsvpResponseTypeTraitMapping = map[string]func(*schema.RsvpResponseTypeTrait, []string){}
var customRsvpResponseTypeTraitMapping = map[string]func(*schema.RsvpResponseTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RsvpResponseType", func(ctx jsonld.Context) (interface{}) {
		return NewRsvpResponseTypeFromContext(ctx)
	})

}

func NewRsvpResponseTypeFromContext(ctx jsonld.Context) (x schema.RsvpResponseType) {
	x.Type = "http://schema.org/RsvpResponseType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRsvpResponseTypeTrait(ctx, &x.RsvpResponseTypeTrait)
	MapCustomToRsvpResponseTypeTrait(ctx, &x.RsvpResponseTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRsvpResponseTypeTrait(ctx jsonld.Context, x *schema.RsvpResponseTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRsvpResponseTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRsvpResponseTypeTrait(ctx jsonld.Context, x *schema.RsvpResponseTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRsvpResponseTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}