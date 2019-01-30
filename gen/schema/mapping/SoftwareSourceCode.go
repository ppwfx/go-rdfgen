package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSoftwareSourceCode strings.Replacer
var strconvSoftwareSourceCode strconv.NumError

var basicSoftwareSourceCodeTraitMapping = map[string]func(*schema.SoftwareSourceCodeTrait, []string){}
var customSoftwareSourceCodeTraitMapping = map[string]func(*schema.SoftwareSourceCodeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SoftwareSourceCode", func(ctx jsonld.Context) (interface{}) {
		return NewSoftwareSourceCodeFromContext(ctx)
	})


	basicSoftwareSourceCodeTraitMapping["http://schema.org/sampleType"] = func(x *schema.SoftwareSourceCodeTrait, s []string) {
		x.SampleType = s[0]
	}


	basicSoftwareSourceCodeTraitMapping["http://schema.org/runtime"] = func(x *schema.SoftwareSourceCodeTrait, s []string) {
		x.Runtime = s[0]
	}


	basicSoftwareSourceCodeTraitMapping["http://schema.org/runtimePlatform"] = func(x *schema.SoftwareSourceCodeTrait, s []string) {
		x.RuntimePlatform = s[0]
	}



	basicSoftwareSourceCodeTraitMapping["http://schema.org/codeSampleType"] = func(x *schema.SoftwareSourceCodeTrait, s []string) {
		x.CodeSampleType = s[0]
	}


	basicSoftwareSourceCodeTraitMapping["http://schema.org/codeRepository"] = func(x *schema.SoftwareSourceCodeTrait, s []string) {
		x.CodeRepository = s[0]
	}



	customSoftwareSourceCodeTraitMapping["http://schema.org/programmingLanguage"] = func(x *schema.SoftwareSourceCodeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ProgrammingLanguage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ProgrammingLanguage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ProgrammingLanguage = s
		}
	}

	customSoftwareSourceCodeTraitMapping["http://schema.org/targetProduct"] = func(x *schema.SoftwareSourceCodeTrait, ctx jsonld.Context, s string){
		var y schema.SoftwareApplication
		if strings.HasPrefix(s, "_:") {
			y = NewSoftwareApplicationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewSoftwareApplication()
			y.Id = s
		}

		x.TargetProduct = &y

		return
	}
}

func NewSoftwareSourceCodeFromContext(ctx jsonld.Context) (x schema.SoftwareSourceCode) {
	x.Type = "http://schema.org/SoftwareSourceCode"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSoftwareSourceCodeTrait(ctx, &x.SoftwareSourceCodeTrait)
	MapCustomToSoftwareSourceCodeTrait(ctx, &x.SoftwareSourceCodeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSoftwareSourceCodeTrait(ctx jsonld.Context, x *schema.SoftwareSourceCodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSoftwareSourceCodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSoftwareSourceCodeTrait(ctx jsonld.Context, x *schema.SoftwareSourceCodeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSoftwareSourceCodeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}