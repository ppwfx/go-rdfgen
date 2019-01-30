package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMediaSubscription strings.Replacer
var strconvMediaSubscription strconv.NumError

var basicMediaSubscriptionTraitMapping = map[string]func(*schema.MediaSubscriptionTrait, []string){}
var customMediaSubscriptionTraitMapping = map[string]func(*schema.MediaSubscriptionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MediaSubscription", func(ctx jsonld.Context) (interface{}) {
		return NewMediaSubscriptionFromContext(ctx)
	})




	customMediaSubscriptionTraitMapping["http://schema.org/expectsAcceptanceOf"] = func(x *schema.MediaSubscriptionTrait, ctx jsonld.Context, s string){
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

	customMediaSubscriptionTraitMapping["http://schema.org/authenticator"] = func(x *schema.MediaSubscriptionTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.Authenticator = &y

		return
	}
}

func NewMediaSubscriptionFromContext(ctx jsonld.Context) (x schema.MediaSubscription) {
	x.Type = "http://schema.org/MediaSubscription"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMediaSubscriptionTrait(ctx, &x.MediaSubscriptionTrait)
	MapCustomToMediaSubscriptionTrait(ctx, &x.MediaSubscriptionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMediaSubscriptionTrait(ctx jsonld.Context, x *schema.MediaSubscriptionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMediaSubscriptionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMediaSubscriptionTrait(ctx jsonld.Context, x *schema.MediaSubscriptionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMediaSubscriptionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}