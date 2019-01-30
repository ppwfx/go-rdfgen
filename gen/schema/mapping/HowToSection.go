package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHowToSection strings.Replacer
var strconvHowToSection strconv.NumError

var basicHowToSectionTraitMapping = map[string]func(*schema.HowToSectionTrait, []string){}
var customHowToSectionTraitMapping = map[string]func(*schema.HowToSectionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HowToSection", func(ctx jsonld.Context) (interface{}) {
		return NewHowToSectionFromContext(ctx)
	})



	customHowToSectionTraitMapping["http://schema.org/steps"] = func(x *schema.HowToSectionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Steps, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Steps = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Steps = s
		}
	}
}

func NewHowToSectionFromContext(ctx jsonld.Context) (x schema.HowToSection) {
	x.Type = "http://schema.org/HowToSection"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToListItemTrait(ctx, &x.ListItemTrait)
	MapCustomToListItemTrait(ctx, &x.ListItemTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToItemListTrait(ctx, &x.ItemListTrait)
	MapCustomToItemListTrait(ctx, &x.ItemListTrait)


	MapBasicToHowToSectionTrait(ctx, &x.HowToSectionTrait)
	MapCustomToHowToSectionTrait(ctx, &x.HowToSectionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHowToSectionTrait(ctx jsonld.Context, x *schema.HowToSectionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHowToSectionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHowToSectionTrait(ctx jsonld.Context, x *schema.HowToSectionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHowToSectionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}