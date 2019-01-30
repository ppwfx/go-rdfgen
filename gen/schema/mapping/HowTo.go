package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHowTo strings.Replacer
var strconvHowTo strconv.NumError

var basicHowToTraitMapping = map[string]func(*schema.HowToTrait, []string){}
var customHowToTraitMapping = map[string]func(*schema.HowToTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HowTo", func(ctx jsonld.Context) (interface{}) {
		return NewHowToFromContext(ctx)
	})











	customHowToTraitMapping["http://schema.org/supply"] = func(x *schema.HowToTrait, ctx jsonld.Context, s string){
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

	customHowToTraitMapping["http://schema.org/steps"] = func(x *schema.HowToTrait, ctx jsonld.Context, s string){
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

	customHowToTraitMapping["http://schema.org/yield"] = func(x *schema.HowToTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Yield, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Yield = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Yield = s
		}
	}

	customHowToTraitMapping["http://schema.org/estimatedCost"] = func(x *schema.HowToTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EstimatedCost, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EstimatedCost = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EstimatedCost = s
		}
	}

	customHowToTraitMapping["http://schema.org/step"] = func(x *schema.HowToTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Step, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Step = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Step = s
		}
	}

	customHowToTraitMapping["http://schema.org/tool"] = func(x *schema.HowToTrait, ctx jsonld.Context, s string){
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

	customHowToTraitMapping["http://schema.org/performTime"] = func(x *schema.HowToTrait, ctx jsonld.Context, s string){
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

	customHowToTraitMapping["http://schema.org/prepTime"] = func(x *schema.HowToTrait, ctx jsonld.Context, s string){
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

	customHowToTraitMapping["http://schema.org/totalTime"] = func(x *schema.HowToTrait, ctx jsonld.Context, s string){
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
}

func NewHowToFromContext(ctx jsonld.Context) (x schema.HowTo) {
	x.Type = "http://schema.org/HowTo"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHowToTrait(ctx, &x.HowToTrait)
	MapCustomToHowToTrait(ctx, &x.HowToTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHowToTrait(ctx jsonld.Context, x *schema.HowToTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHowToTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHowToTrait(ctx jsonld.Context, x *schema.HowToTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHowToTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}