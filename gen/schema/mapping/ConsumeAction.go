package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsConsumeAction strings.Replacer
var strconvConsumeAction strconv.NumError

var basicConsumeActionTraitMapping = map[string]func(*schema.ConsumeActionTrait, []string){}
var customConsumeActionTraitMapping = map[string]func(*schema.ConsumeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ConsumeAction", func(ctx jsonld.Context) (interface{}) {
		return NewConsumeActionFromContext(ctx)
	})




	customConsumeActionTraitMapping["http://schema.org/expectsAcceptanceOf"] = func(x *schema.ConsumeActionTrait, ctx jsonld.Context, s string){
		var y schema.Offer
		if strings.HasPrefix(s, "_:") {
			y = NewOfferFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOffer()
			y.Id = s
		}

		x.ExpectsAcceptanceOf = &y

		return
	}

	customConsumeActionTraitMapping["http://schema.org/actionAccessibilityRequirement"] = func(x *schema.ConsumeActionTrait, ctx jsonld.Context, s string){
		var y schema.ActionAccessSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewActionAccessSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewActionAccessSpecification()
			y.Id = s
		}

		x.ActionAccessibilityRequirement = &y

		return
	}
}

func NewConsumeActionFromContext(ctx jsonld.Context) (x schema.ConsumeAction) {
	x.Type = "http://schema.org/ConsumeAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToConsumeActionTrait(ctx, &x.ConsumeActionTrait)
	MapCustomToConsumeActionTrait(ctx, &x.ConsumeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToConsumeActionTrait(ctx jsonld.Context, x *schema.ConsumeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicConsumeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToConsumeActionTrait(ctx jsonld.Context, x *schema.ConsumeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customConsumeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}