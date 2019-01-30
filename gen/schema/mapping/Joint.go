package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsJoint strings.Replacer
var strconvJoint strconv.NumError

var basicJointTraitMapping = map[string]func(*schema.JointTrait, []string){}
var customJointTraitMapping = map[string]func(*schema.JointTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Joint", func(ctx jsonld.Context) (interface{}) {
		return NewJointFromContext(ctx)
	})


	basicJointTraitMapping["http://schema.org/biomechnicalClass"] = func(x *schema.JointTrait, s []string) {
		x.BiomechnicalClass = s[0]
	}


	basicJointTraitMapping["http://schema.org/structuralClass"] = func(x *schema.JointTrait, s []string) {
		x.StructuralClass = s[0]
	}



	customJointTraitMapping["http://schema.org/functionalClass"] = func(x *schema.JointTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.FunctionalClass, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.FunctionalClass = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.FunctionalClass = s
		}
	}
}

func NewJointFromContext(ctx jsonld.Context) (x schema.Joint) {
	x.Type = "http://schema.org/Joint"
	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToJointTrait(ctx, &x.JointTrait)
	MapCustomToJointTrait(ctx, &x.JointTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToJointTrait(ctx jsonld.Context, x *schema.JointTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicJointTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToJointTrait(ctx jsonld.Context, x *schema.JointTrait) () {
	for k, v := range ctx.Current {
		f, ok := customJointTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}