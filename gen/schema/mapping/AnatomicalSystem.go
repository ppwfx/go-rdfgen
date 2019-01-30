package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAnatomicalSystem strings.Replacer
var strconvAnatomicalSystem strconv.NumError

var basicAnatomicalSystemTraitMapping = map[string]func(*schema.AnatomicalSystemTrait, []string){}
var customAnatomicalSystemTraitMapping = map[string]func(*schema.AnatomicalSystemTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AnatomicalSystem", func(ctx jsonld.Context) (interface{}) {
		return NewAnatomicalSystemFromContext(ctx)
	})


	basicAnatomicalSystemTraitMapping["http://schema.org/associatedPathophysiology"] = func(x *schema.AnatomicalSystemTrait, s []string) {
		x.AssociatedPathophysiology = s[0]
	}






	customAnatomicalSystemTraitMapping["http://schema.org/relatedCondition"] = func(x *schema.AnatomicalSystemTrait, ctx jsonld.Context, s string){
		var y schema.MedicalCondition
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalConditionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalCondition()
			y.Id = s
		}

		x.RelatedCondition = &y

		return
	}

	customAnatomicalSystemTraitMapping["http://schema.org/relatedTherapy"] = func(x *schema.AnatomicalSystemTrait, ctx jsonld.Context, s string){
		var y schema.MedicalTherapy
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalTherapyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalTherapy()
			y.Id = s
		}

		x.RelatedTherapy = &y

		return
	}

	customAnatomicalSystemTraitMapping["http://schema.org/comprisedOf"] = func(x *schema.AnatomicalSystemTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ComprisedOf, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ComprisedOf = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ComprisedOf = s
		}
	}

	customAnatomicalSystemTraitMapping["http://schema.org/relatedStructure"] = func(x *schema.AnatomicalSystemTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalStructure
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalStructure()
			y.Id = s
		}

		x.RelatedStructure = &y

		return
	}
}

func NewAnatomicalSystemFromContext(ctx jsonld.Context) (x schema.AnatomicalSystem) {
	x.Type = "http://schema.org/AnatomicalSystem"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAnatomicalSystemTrait(ctx, &x.AnatomicalSystemTrait)
	MapCustomToAnatomicalSystemTrait(ctx, &x.AnatomicalSystemTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAnatomicalSystemTrait(ctx jsonld.Context, x *schema.AnatomicalSystemTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAnatomicalSystemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAnatomicalSystemTrait(ctx jsonld.Context, x *schema.AnatomicalSystemTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAnatomicalSystemTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}