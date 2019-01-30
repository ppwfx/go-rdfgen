package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLigament strings.Replacer
var strconvLigament strconv.NumError

var basicLigamentTraitMapping = map[string]func(*schema.LigamentTrait, []string){}
var customLigamentTraitMapping = map[string]func(*schema.LigamentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Ligament", func(ctx jsonld.Context) (interface{}) {
		return NewLigamentFromContext(ctx)
	})

}

func NewLigamentFromContext(ctx jsonld.Context) (x schema.Ligament) {
	x.Type = "http://schema.org/Ligament"
	MapBasicToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)
	MapCustomToAnatomicalStructureTrait(ctx, &x.AnatomicalStructureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLigamentTrait(ctx, &x.LigamentTrait)
	MapCustomToLigamentTrait(ctx, &x.LigamentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLigamentTrait(ctx jsonld.Context, x *schema.LigamentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLigamentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLigamentTrait(ctx jsonld.Context, x *schema.LigamentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLigamentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}