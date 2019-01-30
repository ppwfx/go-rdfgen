package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHowToDirection strings.Replacer
var strconvHowToDirection strconv.NumError

var basicHowToDirectionTraitMapping = map[string]func(*schema.HowToDirectionTrait, []string){}
var customHowToDirectionTraitMapping = map[string]func(*schema.HowToDirectionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HowToDirection", func(ctx jsonld.Context) (interface{}) {
		return NewHowToDirectionFromContext(ctx)
	})










	customHowToDirectionTraitMapping["http://schema.org/supply"] = func(x *schema.HowToDirectionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Supply, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Supply = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Supply = s
		}
	}

	customHowToDirectionTraitMapping["http://schema.org/tool"] = func(x *schema.HowToDirectionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Tool, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Tool = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Tool = s
		}
	}

	customHowToDirectionTraitMapping["http://schema.org/performTime"] = func(x *schema.HowToDirectionTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.PerformTime = &y

		return
	}

	customHowToDirectionTraitMapping["http://schema.org/prepTime"] = func(x *schema.HowToDirectionTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.PrepTime = &y

		return
	}

	customHowToDirectionTraitMapping["http://schema.org/totalTime"] = func(x *schema.HowToDirectionTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.TotalTime = &y

		return
	}

	customHowToDirectionTraitMapping["http://schema.org/afterMedia"] = func(x *schema.HowToDirectionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AfterMedia, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AfterMedia = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AfterMedia = s
		}
	}

	customHowToDirectionTraitMapping["http://schema.org/beforeMedia"] = func(x *schema.HowToDirectionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BeforeMedia, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BeforeMedia = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BeforeMedia = s
		}
	}

	customHowToDirectionTraitMapping["http://schema.org/duringMedia"] = func(x *schema.HowToDirectionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DuringMedia, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DuringMedia = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DuringMedia = s
		}
	}
}

func NewHowToDirectionFromContext(ctx jsonld.Context) (x schema.HowToDirection) {
	x.Type = "http://schema.org/HowToDirection"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToListItemTrait(ctx, &x.ListItemTrait)
	MapCustomToListItemTrait(ctx, &x.ListItemTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)


	MapBasicToHowToDirectionTrait(ctx, &x.HowToDirectionTrait)
	MapCustomToHowToDirectionTrait(ctx, &x.HowToDirectionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHowToDirectionTrait(ctx jsonld.Context, x *schema.HowToDirectionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHowToDirectionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHowToDirectionTrait(ctx jsonld.Context, x *schema.HowToDirectionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHowToDirectionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}