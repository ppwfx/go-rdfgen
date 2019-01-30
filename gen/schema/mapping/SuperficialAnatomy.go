package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSuperficialAnatomy strings.Replacer
var strconvSuperficialAnatomy strconv.NumError

var basicSuperficialAnatomyTraitMapping = map[string]func(*schema.SuperficialAnatomyTrait, []string){}
var customSuperficialAnatomyTraitMapping = map[string]func(*schema.SuperficialAnatomyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SuperficialAnatomy", func(ctx jsonld.Context) (interface{}) {
		return NewSuperficialAnatomyFromContext(ctx)
	})


	basicSuperficialAnatomyTraitMapping["http://schema.org/associatedPathophysiology"] = func(x *schema.SuperficialAnatomyTrait, s []string) {
		x.AssociatedPathophysiology = s[0]
	}


	basicSuperficialAnatomyTraitMapping["http://schema.org/significance"] = func(x *schema.SuperficialAnatomyTrait, s []string) {
		x.Significance = s[0]
	}





	customSuperficialAnatomyTraitMapping["http://schema.org/relatedCondition"] = func(x *schema.SuperficialAnatomyTrait, ctx jsonld.Context, s string){
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

	customSuperficialAnatomyTraitMapping["http://schema.org/relatedTherapy"] = func(x *schema.SuperficialAnatomyTrait, ctx jsonld.Context, s string){
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

	customSuperficialAnatomyTraitMapping["http://schema.org/relatedAnatomy"] = func(x *schema.SuperficialAnatomyTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RelatedAnatomy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RelatedAnatomy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RelatedAnatomy = s
		}
	}
}

func NewSuperficialAnatomyFromContext(ctx jsonld.Context) (x schema.SuperficialAnatomy) {
	x.Type = "http://schema.org/SuperficialAnatomy"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSuperficialAnatomyTrait(ctx, &x.SuperficialAnatomyTrait)
	MapCustomToSuperficialAnatomyTrait(ctx, &x.SuperficialAnatomyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSuperficialAnatomyTrait(ctx jsonld.Context, x *schema.SuperficialAnatomyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSuperficialAnatomyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSuperficialAnatomyTrait(ctx jsonld.Context, x *schema.SuperficialAnatomyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSuperficialAnatomyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}