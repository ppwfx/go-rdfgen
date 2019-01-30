package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBoardingPolicyType strings.Replacer
var strconvBoardingPolicyType strconv.NumError

var basicBoardingPolicyTypeTraitMapping = map[string]func(*schema.BoardingPolicyTypeTrait, []string){}
var customBoardingPolicyTypeTraitMapping = map[string]func(*schema.BoardingPolicyTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BoardingPolicyType", func(ctx jsonld.Context) (interface{}) {
		return NewBoardingPolicyTypeFromContext(ctx)
	})

}

func NewBoardingPolicyTypeFromContext(ctx jsonld.Context) (x schema.BoardingPolicyType) {
	x.Type = "http://schema.org/BoardingPolicyType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBoardingPolicyTypeTrait(ctx, &x.BoardingPolicyTypeTrait)
	MapCustomToBoardingPolicyTypeTrait(ctx, &x.BoardingPolicyTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBoardingPolicyTypeTrait(ctx jsonld.Context, x *schema.BoardingPolicyTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBoardingPolicyTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBoardingPolicyTypeTrait(ctx jsonld.Context, x *schema.BoardingPolicyTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBoardingPolicyTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}