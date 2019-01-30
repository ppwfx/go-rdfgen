package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsActionStatusType strings.Replacer
var strconvActionStatusType strconv.NumError

var basicActionStatusTypeTraitMapping = map[string]func(*schema.ActionStatusTypeTrait, []string){}
var customActionStatusTypeTraitMapping = map[string]func(*schema.ActionStatusTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ActionStatusType", func(ctx jsonld.Context) (interface{}) {
		return NewActionStatusTypeFromContext(ctx)
	})

}

func NewActionStatusTypeFromContext(ctx jsonld.Context) (x schema.ActionStatusType) {
	x.Type = "http://schema.org/ActionStatusType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToActionStatusTypeTrait(ctx, &x.ActionStatusTypeTrait)
	MapCustomToActionStatusTypeTrait(ctx, &x.ActionStatusTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToActionStatusTypeTrait(ctx jsonld.Context, x *schema.ActionStatusTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicActionStatusTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToActionStatusTypeTrait(ctx jsonld.Context, x *schema.ActionStatusTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customActionStatusTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}