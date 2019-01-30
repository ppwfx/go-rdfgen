package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAnatomicalStructure strings.Replacer
var strconvAnatomicalStructure strconv.NumError

var basicAnatomicalStructureTraitMapping = map[string]func(*schema.AnatomicalStructureTrait, []string){}
var customAnatomicalStructureTraitMapping = map[string]func(*schema.AnatomicalStructureTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AnatomicalStructure", func(ctx jsonld.Context) (interface{}) {
		return NewAnatomicalStructureFromContext(ctx)
	})



	basicAnatomicalStructureTraitMapping["http://schema.org/associatedPathophysiology"] = func(x *schema.AnatomicalStructureTrait, s []string) {
		x.AssociatedPathophysiology = s[0]
	}


	basicAnatomicalStructureTraitMapping["http://schema.org/function"] = func(x *schema.AnatomicalStructureTrait, s []string) {
		x.Function = s[0]
	}


	basicAnatomicalStructureTraitMapping["http://schema.org/bodyLocation"] = func(x *schema.AnatomicalStructureTrait, s []string) {
		x.BodyLocation = s[0]
	}







	customAnatomicalStructureTraitMapping["http://schema.org/diagram"] = func(x *schema.AnatomicalStructureTrait, ctx jsonld.Context, s string){
		var y schema.ImageObject
		if strings.HasPrefix(s, "_:") {
			y = NewImageObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewImageObject()
			y.Id = s
		}

		x.Diagram = &y

		return
	}

	customAnatomicalStructureTraitMapping["http://schema.org/connectedTo"] = func(x *schema.AnatomicalStructureTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalStructure
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalStructure()
			y.Id = s
		}

		x.ConnectedTo = &y

		return
	}

	customAnatomicalStructureTraitMapping["http://schema.org/partOfSystem"] = func(x *schema.AnatomicalStructureTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalSystem
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalSystemFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalSystem()
			y.Id = s
		}

		x.PartOfSystem = &y

		return
	}

	customAnatomicalStructureTraitMapping["http://schema.org/relatedCondition"] = func(x *schema.AnatomicalStructureTrait, ctx jsonld.Context, s string){
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

	customAnatomicalStructureTraitMapping["http://schema.org/relatedTherapy"] = func(x *schema.AnatomicalStructureTrait, ctx jsonld.Context, s string){
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

	customAnatomicalStructureTraitMapping["http://schema.org/subStructure"] = func(x *schema.AnatomicalStructureTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalStructure
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalStructure()
			y.Id = s
		}

		x.SubStructure = &y

		return
	}
}

func NewAnatomicalStructureFromContext(ctx jsonld.Context) (x schema.AnatomicalStructure) {
	x.Type = "http://schema.org/AnatomicalStructure"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAnatomicalStructureTrait(ctx jsonld.Context, x *schema.AnatomicalStructureTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAnatomicalStructureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAnatomicalStructureTrait(ctx jsonld.Context, x *schema.AnatomicalStructureTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAnatomicalStructureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}