package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBone strings.Replacer
var strconvBone strconv.NumError

var basicBoneTraitMapping = map[string]func(*schema.BoneTrait, []string){}
var customBoneTraitMapping = map[string]func(*schema.BoneTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Bone", func(ctx jsonld.Context) (interface{}) {
		return NewBoneFromContext(ctx)
	})

}

func NewBoneFromContext(ctx jsonld.Context) (x schema.Bone) {
	x.Type = "http://schema.org/Bone"
	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBoneTrait(ctx, &x.BoneTrait)
	MapCustomToBoneTrait(ctx, &x.BoneTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBoneTrait(ctx jsonld.Context, x *schema.BoneTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBoneTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBoneTrait(ctx jsonld.Context, x *schema.BoneTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBoneTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}