package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalTestPanel strings.Replacer
var strconvMedicalTestPanel strconv.NumError

var basicMedicalTestPanelTraitMapping = map[string]func(*schema.MedicalTestPanelTrait, []string){}
var customMedicalTestPanelTraitMapping = map[string]func(*schema.MedicalTestPanelTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalTestPanel", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalTestPanelFromContext(ctx)
	})



	customMedicalTestPanelTraitMapping["http://schema.org/subTest"] = func(x *schema.MedicalTestPanelTrait, ctx jsonld.Context, s string){
		var y schema.MedicalTest
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalTestFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalTest()
			y.Id = s
		}

		x.SubTest = &y

		return
	}
}

func NewMedicalTestPanelFromContext(ctx jsonld.Context) (x schema.MedicalTestPanel) {
	x.Type = "http://schema.org/MedicalTestPanel"
	MapBasicToMedicalTestTrait(ctx, &x.MedicalTestTrait)
	MapCustomToMedicalTestTrait(ctx, &x.MedicalTestTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalTestPanelTrait(ctx, &x.MedicalTestPanelTrait)
	MapCustomToMedicalTestPanelTrait(ctx, &x.MedicalTestPanelTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalTestPanelTrait(ctx jsonld.Context, x *schema.MedicalTestPanelTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalTestPanelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalTestPanelTrait(ctx jsonld.Context, x *schema.MedicalTestPanelTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalTestPanelTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}