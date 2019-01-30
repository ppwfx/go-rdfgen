package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVein strings.Replacer
var strconvVein strconv.NumError

var basicVeinTraitMapping = map[string]func(*schema.VeinTrait, []string){}
var customVeinTraitMapping = map[string]func(*schema.VeinTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Vein", func(ctx jsonld.Context) (interface{}) {
		return NewVeinFromContext(ctx)
	})





	customVeinTraitMapping["http://schema.org/tributary"] = func(x *schema.VeinTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalStructure
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalStructure()
			y.Id = s
		}

		x.Tributary = &y

		return
	}

	customVeinTraitMapping["http://schema.org/regionDrained"] = func(x *schema.VeinTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RegionDrained, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RegionDrained = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RegionDrained = s
		}
	}

	customVeinTraitMapping["http://schema.org/drainsTo"] = func(x *schema.VeinTrait, ctx jsonld.Context, s string){
		var y schema.Vessel
		if strings.HasPrefix(s, "_:") {
			y = NewVesselFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewVessel()
			y.Id = s
		}

		x.DrainsTo = &y

		return
	}
}

func NewVeinFromContext(ctx jsonld.Context) (x schema.Vein) {
	x.Type = "http://schema.org/Vein"
	MapBasicToVesselTrait(ctx, &x.VesselTrait)
	MapCustomToVesselTrait(ctx, &x.VesselTrait)

	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVeinTrait(ctx, &x.VeinTrait)
	MapCustomToVeinTrait(ctx, &x.VeinTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVeinTrait(ctx jsonld.Context, x *schema.VeinTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVeinTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVeinTrait(ctx jsonld.Context, x *schema.VeinTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVeinTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}