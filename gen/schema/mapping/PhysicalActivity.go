package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPhysicalActivity strings.Replacer
var strconvPhysicalActivity strconv.NumError

var basicPhysicalActivityTraitMapping = map[string]func(*schema.PhysicalActivityTrait, []string){}
var customPhysicalActivityTraitMapping = map[string]func(*schema.PhysicalActivityTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PhysicalActivity", func(ctx jsonld.Context) (interface{}) {
		return NewPhysicalActivityFromContext(ctx)
	})



	basicPhysicalActivityTraitMapping["http://schema.org/pathophysiology"] = func(x *schema.PhysicalActivityTrait, s []string) {
		x.Pathophysiology = s[0]
	}


	basicPhysicalActivityTraitMapping["http://schema.org/epidemiology"] = func(x *schema.PhysicalActivityTrait, s []string) {
		x.Epidemiology = s[0]
	}



	customPhysicalActivityTraitMapping["http://schema.org/category"] = func(x *schema.PhysicalActivityTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Category, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Category = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Category = s
		}
	}

	customPhysicalActivityTraitMapping["http://schema.org/associatedAnatomy"] = func(x *schema.PhysicalActivityTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AssociatedAnatomy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AssociatedAnatomy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AssociatedAnatomy = s
		}
	}
}

func NewPhysicalActivityFromContext(ctx jsonld.Context) (x schema.PhysicalActivity) {
	x.Type = "http://schema.org/PhysicalActivity"
	MapBasicToLifestyleModificationTrait(ctx, &x.LifestyleModificationTrait)
	MapCustomToLifestyleModificationTrait(ctx, &x.LifestyleModificationTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPhysicalActivityTrait(ctx, &x.PhysicalActivityTrait)
	MapCustomToPhysicalActivityTrait(ctx, &x.PhysicalActivityTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPhysicalActivityTrait(ctx jsonld.Context, x *schema.PhysicalActivityTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPhysicalActivityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPhysicalActivityTrait(ctx jsonld.Context, x *schema.PhysicalActivityTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPhysicalActivityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}