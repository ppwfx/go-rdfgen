package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalWebPage strings.Replacer
var strconvMedicalWebPage strconv.NumError

var basicMedicalWebPageTraitMapping = map[string]func(*schema.MedicalWebPageTrait, []string){}
var customMedicalWebPageTraitMapping = map[string]func(*schema.MedicalWebPageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalWebPage", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalWebPageFromContext(ctx)
	})


	basicMedicalWebPageTraitMapping["http://schema.org/aspect"] = func(x *schema.MedicalWebPageTrait, s []string) {
		x.Aspect = s[0]
	}

}

func NewMedicalWebPageFromContext(ctx jsonld.Context) (x schema.MedicalWebPage) {
	x.Type = "http://schema.org/MedicalWebPage"
	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalWebPageTrait(ctx, &x.MedicalWebPageTrait)
	MapCustomToMedicalWebPageTrait(ctx, &x.MedicalWebPageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalWebPageTrait(ctx jsonld.Context, x *schema.MedicalWebPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalWebPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalWebPageTrait(ctx jsonld.Context, x *schema.MedicalWebPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalWebPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}