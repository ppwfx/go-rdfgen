package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsImagingTest strings.Replacer
var strconvImagingTest strconv.NumError

var basicImagingTestTraitMapping = map[string]func(*schema.ImagingTestTrait, []string){}
var customImagingTestTraitMapping = map[string]func(*schema.ImagingTestTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ImagingTest", func(ctx jsonld.Context) (interface{}) {
		return NewImagingTestFromContext(ctx)
	})



	customImagingTestTraitMapping["http://schema.org/imagingTechnique"] = func(x *schema.ImagingTestTrait, ctx jsonld.Context, s string){
		var y schema.MedicalImagingTechnique
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalImagingTechniqueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalImagingTechnique()
			y.Id = s
		}

		x.ImagingTechnique = &y

		return
	}
}

func NewImagingTestFromContext(ctx jsonld.Context) (x schema.ImagingTest) {
	x.Type = "http://schema.org/ImagingTest"
	MapBasicToMedicalTestTrait(ctx, &x.MedicalTestTrait)
	MapCustomToMedicalTestTrait(ctx, &x.MedicalTestTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToImagingTestTrait(ctx, &x.ImagingTestTrait)
	MapCustomToImagingTestTrait(ctx, &x.ImagingTestTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToImagingTestTrait(ctx jsonld.Context, x *schema.ImagingTestTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicImagingTestTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToImagingTestTrait(ctx jsonld.Context, x *schema.ImagingTestTrait) () {
	for k, v := range ctx.Current {
		f, ok := customImagingTestTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}