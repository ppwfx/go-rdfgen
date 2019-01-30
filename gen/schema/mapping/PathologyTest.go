package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPathologyTest strings.Replacer
var strconvPathologyTest strconv.NumError

var basicPathologyTestTraitMapping = map[string]func(*schema.PathologyTestTrait, []string){}
var customPathologyTestTraitMapping = map[string]func(*schema.PathologyTestTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PathologyTest", func(ctx jsonld.Context) (interface{}) {
		return NewPathologyTestFromContext(ctx)
	})


	basicPathologyTestTraitMapping["http://schema.org/tissueSample"] = func(x *schema.PathologyTestTrait, s []string) {
		x.TissueSample = s[0]
	}

}

func NewPathologyTestFromContext(ctx jsonld.Context) (x schema.PathologyTest) {
	x.Type = "http://schema.org/PathologyTest"
	MapBasicToMedicalTestTrait(ctx, &x.MedicalTestTrait)
	MapCustomToMedicalTestTrait(ctx, &x.MedicalTestTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPathologyTestTrait(ctx, &x.PathologyTestTrait)
	MapCustomToPathologyTestTrait(ctx, &x.PathologyTestTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPathologyTestTrait(ctx jsonld.Context, x *schema.PathologyTestTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPathologyTestTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPathologyTestTrait(ctx jsonld.Context, x *schema.PathologyTestTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPathologyTestTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}