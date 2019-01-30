package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLymphaticVessel strings.Replacer
var strconvLymphaticVessel strconv.NumError

var basicLymphaticVesselTraitMapping = map[string]func(*schema.LymphaticVesselTrait, []string){}
var customLymphaticVesselTraitMapping = map[string]func(*schema.LymphaticVesselTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LymphaticVessel", func(ctx jsonld.Context) (interface{}) {
		return NewLymphaticVesselFromContext(ctx)
	})





	customLymphaticVesselTraitMapping["http://schema.org/regionDrained"] = func(x *schema.LymphaticVesselTrait, ctx jsonld.Context, s string){
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

	customLymphaticVesselTraitMapping["http://schema.org/originatesFrom"] = func(x *schema.LymphaticVesselTrait, ctx jsonld.Context, s string){
		var y schema.Vessel
		if strings.HasPrefix(s, "_:") {
			y = NewVesselFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewVessel()
			y.Id = s
		}

		x.OriginatesFrom = &y

		return
	}

	customLymphaticVesselTraitMapping["http://schema.org/runsTo"] = func(x *schema.LymphaticVesselTrait, ctx jsonld.Context, s string){
		var y schema.Vessel
		if strings.HasPrefix(s, "_:") {
			y = NewVesselFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewVessel()
			y.Id = s
		}

		x.RunsTo = &y

		return
	}
}

func NewLymphaticVesselFromContext(ctx jsonld.Context) (x schema.LymphaticVessel) {
	x.Type = "http://schema.org/LymphaticVessel"
	MapBasicToVesselTrait(ctx, &x.VesselTrait)
	MapCustomToVesselTrait(ctx, &x.VesselTrait)

	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLymphaticVesselTrait(ctx, &x.LymphaticVesselTrait)
	MapCustomToLymphaticVesselTrait(ctx, &x.LymphaticVesselTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLymphaticVesselTrait(ctx jsonld.Context, x *schema.LymphaticVesselTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLymphaticVesselTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLymphaticVesselTrait(ctx jsonld.Context, x *schema.LymphaticVesselTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLymphaticVesselTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}