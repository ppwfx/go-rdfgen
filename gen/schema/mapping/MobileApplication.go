package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMobileApplication strings.Replacer
var strconvMobileApplication strconv.NumError

var basicMobileApplicationTraitMapping = map[string]func(*schema.MobileApplicationTrait, []string){}
var customMobileApplicationTraitMapping = map[string]func(*schema.MobileApplicationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MobileApplication", func(ctx jsonld.Context) (interface{}) {
		return NewMobileApplicationFromContext(ctx)
	})


	basicMobileApplicationTraitMapping["http://schema.org/carrierRequirements"] = func(x *schema.MobileApplicationTrait, s []string) {
		x.CarrierRequirements = s[0]
	}

}

func NewMobileApplicationFromContext(ctx jsonld.Context) (x schema.MobileApplication) {
	x.Type = "http://schema.org/MobileApplication"
	MapBasicToSoftwareApplicationTrait(ctx, &x.SoftwareApplicationTrait)
	MapCustomToSoftwareApplicationTrait(ctx, &x.SoftwareApplicationTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMobileApplicationTrait(ctx, &x.MobileApplicationTrait)
	MapCustomToMobileApplicationTrait(ctx, &x.MobileApplicationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMobileApplicationTrait(ctx jsonld.Context, x *schema.MobileApplicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMobileApplicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMobileApplicationTrait(ctx jsonld.Context, x *schema.MobileApplicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMobileApplicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}