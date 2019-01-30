package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGovernmentPermit strings.Replacer
var strconvGovernmentPermit strconv.NumError

var basicGovernmentPermitTraitMapping = map[string]func(*schema.GovernmentPermitTrait, []string){}
var customGovernmentPermitTraitMapping = map[string]func(*schema.GovernmentPermitTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GovernmentPermit", func(ctx jsonld.Context) (interface{}) {
		return NewGovernmentPermitFromContext(ctx)
	})

}

func NewGovernmentPermitFromContext(ctx jsonld.Context) (x schema.GovernmentPermit) {
	x.Type = "http://schema.org/GovernmentPermit"
	MapBasicToPermitTrait(ctx, &x.PermitTrait)
	MapCustomToPermitTrait(ctx, &x.PermitTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGovernmentPermitTrait(ctx, &x.GovernmentPermitTrait)
	MapCustomToGovernmentPermitTrait(ctx, &x.GovernmentPermitTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGovernmentPermitTrait(ctx jsonld.Context, x *schema.GovernmentPermitTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGovernmentPermitTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGovernmentPermitTrait(ctx jsonld.Context, x *schema.GovernmentPermitTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGovernmentPermitTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}