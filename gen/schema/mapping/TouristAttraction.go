package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTouristAttraction strings.Replacer
var strconvTouristAttraction strconv.NumError

var basicTouristAttractionTraitMapping = map[string]func(*schema.TouristAttractionTrait, []string){}
var customTouristAttractionTraitMapping = map[string]func(*schema.TouristAttractionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TouristAttraction", func(ctx jsonld.Context) (interface{}) {
		return NewTouristAttractionFromContext(ctx)
	})




	customTouristAttractionTraitMapping["http://schema.org/availableLanguage"] = func(x *schema.TouristAttractionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AvailableLanguage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AvailableLanguage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AvailableLanguage = s
		}
	}

	customTouristAttractionTraitMapping["http://schema.org/touristType"] = func(x *schema.TouristAttractionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TouristType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TouristType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TouristType = s
		}
	}
}

func NewTouristAttractionFromContext(ctx jsonld.Context) (x schema.TouristAttraction) {
	x.Type = "http://schema.org/TouristAttraction"
	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTouristAttractionTrait(ctx, &x.TouristAttractionTrait)
	MapCustomToTouristAttractionTrait(ctx, &x.TouristAttractionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTouristAttractionTrait(ctx jsonld.Context, x *schema.TouristAttractionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTouristAttractionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTouristAttractionTrait(ctx jsonld.Context, x *schema.TouristAttractionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTouristAttractionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}