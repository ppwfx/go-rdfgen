package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalGuideline strings.Replacer
var strconvMedicalGuideline strconv.NumError

var basicMedicalGuidelineTraitMapping = map[string]func(*schema.MedicalGuidelineTrait, []string){}
var customMedicalGuidelineTraitMapping = map[string]func(*schema.MedicalGuidelineTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalGuideline", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalGuidelineFromContext(ctx)
	})


	basicMedicalGuidelineTraitMapping["http://schema.org/evidenceOrigin"] = func(x *schema.MedicalGuidelineTrait, s []string) {
		x.EvidenceOrigin = s[0]
	}





	customMedicalGuidelineTraitMapping["http://schema.org/guidelineDate"] = func(x *schema.MedicalGuidelineTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.GuidelineDate = &y

		return
	}

	customMedicalGuidelineTraitMapping["http://schema.org/evidenceLevel"] = func(x *schema.MedicalGuidelineTrait, ctx jsonld.Context, s string){
		var y schema.MedicalEvidenceLevel
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalEvidenceLevelFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalEvidenceLevel()
			y.Id = s
		}

		x.EvidenceLevel = &y

		return
	}

	customMedicalGuidelineTraitMapping["http://schema.org/guidelineSubject"] = func(x *schema.MedicalGuidelineTrait, ctx jsonld.Context, s string){
		var y schema.MedicalEntity
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalEntityFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalEntity()
			y.Id = s
		}

		x.GuidelineSubject = &y

		return
	}
}

func NewMedicalGuidelineFromContext(ctx jsonld.Context) (x schema.MedicalGuideline) {
	x.Type = "http://schema.org/MedicalGuideline"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalGuidelineTrait(ctx, &x.MedicalGuidelineTrait)
	MapCustomToMedicalGuidelineTrait(ctx, &x.MedicalGuidelineTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalGuidelineTrait(ctx jsonld.Context, x *schema.MedicalGuidelineTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalGuidelineTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalGuidelineTrait(ctx jsonld.Context, x *schema.MedicalGuidelineTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalGuidelineTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}