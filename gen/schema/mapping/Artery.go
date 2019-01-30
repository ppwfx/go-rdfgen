package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsArtery strings.Replacer
var strconvArtery strconv.NumError

var basicArteryTraitMapping = map[string]func(*schema.ArteryTrait, []string){}
var customArteryTraitMapping = map[string]func(*schema.ArteryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Artery", func(ctx jsonld.Context) (interface{}) {
		return NewArteryFromContext(ctx)
	})





	customArteryTraitMapping["http://schema.org/arterialBranch"] = func(x *schema.ArteryTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalStructure
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalStructure()
			y.Id = s
		}

		x.ArterialBranch = &y

		return
	}

	customArteryTraitMapping["http://schema.org/supplyTo"] = func(x *schema.ArteryTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalStructure
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalStructure()
			y.Id = s
		}

		x.SupplyTo = &y

		return
	}

	customArteryTraitMapping["http://schema.org/source"] = func(x *schema.ArteryTrait, ctx jsonld.Context, s string){
		var y schema.AnatomicalStructure
		if strings.HasPrefix(s, "_:") {
			y = NewAnatomicalStructureFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAnatomicalStructure()
			y.Id = s
		}

		x.Source = &y

		return
	}
}

func NewArteryFromContext(ctx jsonld.Context) (x schema.Artery) {
	x.Type = "http://schema.org/Artery"
	MapBasicToVesselTrait(ctx, &x.VesselTrait)
	MapCustomToVesselTrait(ctx, &x.VesselTrait)

	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToArteryTrait(ctx, &x.ArteryTrait)
	MapCustomToArteryTrait(ctx, &x.ArteryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToArteryTrait(ctx jsonld.Context, x *schema.ArteryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicArteryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToArteryTrait(ctx jsonld.Context, x *schema.ArteryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customArteryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}