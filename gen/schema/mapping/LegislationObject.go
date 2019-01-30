package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLegislationObject strings.Replacer
var strconvLegislationObject strconv.NumError

var basicLegislationObjectTraitMapping = map[string]func(*schema.LegislationObjectTrait, []string){}
var customLegislationObjectTraitMapping = map[string]func(*schema.LegislationObjectTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LegislationObject", func(ctx jsonld.Context) (interface{}) {
		return NewLegislationObjectFromContext(ctx)
	})



	customLegislationObjectTraitMapping["http://schema.org/legislationLegalValue"] = func(x *schema.LegislationObjectTrait, ctx jsonld.Context, s string){
		var y schema.LegalValueLevel
		if strings.HasPrefix(s, "_:") {
			y = NewLegalValueLevelFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewLegalValueLevel()
			y.Id = s
		}

		x.LegislationLegalValue = &y

		return
	}
}

func NewLegislationObjectFromContext(ctx jsonld.Context) (x schema.LegislationObject) {
	x.Type = "http://schema.org/LegislationObject"
	MapBasicToMediaObjectTrait(ctx, &x.MediaObjectTrait)
	MapCustomToMediaObjectTrait(ctx, &x.MediaObjectTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToLegislationTrait(ctx, &x.LegislationTrait)
	MapCustomToLegislationTrait(ctx, &x.LegislationTrait)


	MapBasicToLegislationObjectTrait(ctx, &x.LegislationObjectTrait)
	MapCustomToLegislationObjectTrait(ctx, &x.LegislationObjectTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLegislationObjectTrait(ctx jsonld.Context, x *schema.LegislationObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLegislationObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLegislationObjectTrait(ctx jsonld.Context, x *schema.LegislationObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLegislationObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}