package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsClaim strings.Replacer
var strconvClaim strconv.NumError

var basicClaimTraitMapping = map[string]func(*schema.ClaimTrait, []string){}
var customClaimTraitMapping = map[string]func(*schema.ClaimTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Claim", func(ctx jsonld.Context) (interface{}) {
		return NewClaimFromContext(ctx)
	})




	customClaimTraitMapping["http://schema.org/appearance"] = func(x *schema.ClaimTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.Appearance = &y

		return
	}

	customClaimTraitMapping["http://schema.org/firstAppearance"] = func(x *schema.ClaimTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.FirstAppearance = &y

		return
	}
}

func NewClaimFromContext(ctx jsonld.Context) (x schema.Claim) {
	x.Type = "http://schema.org/Claim"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToClaimTrait(ctx, &x.ClaimTrait)
	MapCustomToClaimTrait(ctx, &x.ClaimTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToClaimTrait(ctx jsonld.Context, x *schema.ClaimTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicClaimTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToClaimTrait(ctx jsonld.Context, x *schema.ClaimTrait) () {
	for k, v := range ctx.Current {
		f, ok := customClaimTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}