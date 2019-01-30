package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBrainStructure strings.Replacer
var strconvBrainStructure strconv.NumError

var basicBrainStructureTraitMapping = map[string]func(*schema.BrainStructureTrait, []string){}
var customBrainStructureTraitMapping = map[string]func(*schema.BrainStructureTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BrainStructure", func(ctx jsonld.Context) (interface{}) {
		return NewBrainStructureFromContext(ctx)
	})

}

func NewBrainStructureFromContext(ctx jsonld.Context) (x schema.BrainStructure) {
	x.Type = "http://schema.org/BrainStructure"
	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBrainStructureTrait(ctx, &x.BrainStructureTrait)
	MapCustomToBrainStructureTrait(ctx, &x.BrainStructureTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBrainStructureTrait(ctx jsonld.Context, x *schema.BrainStructureTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBrainStructureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBrainStructureTrait(ctx jsonld.Context, x *schema.BrainStructureTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBrainStructureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}