package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsScreeningEvent strings.Replacer
var strconvScreeningEvent strconv.NumError

var basicScreeningEventTraitMapping = map[string]func(*schema.ScreeningEventTrait, []string){}
var customScreeningEventTraitMapping = map[string]func(*schema.ScreeningEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ScreeningEvent", func(ctx jsonld.Context) (interface{}) {
		return NewScreeningEventFromContext(ctx)
	})



	basicScreeningEventTraitMapping["http://schema.org/videoFormat"] = func(x *schema.ScreeningEventTrait, s []string) {
		x.VideoFormat = s[0]
	}



	customScreeningEventTraitMapping["http://schema.org/subtitleLanguage"] = func(x *schema.ScreeningEventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SubtitleLanguage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SubtitleLanguage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SubtitleLanguage = s
		}
	}

	customScreeningEventTraitMapping["http://schema.org/workPresented"] = func(x *schema.ScreeningEventTrait, ctx jsonld.Context, s string){
		var y schema.Movie
		if strings.HasPrefix(s, "_:") {
			y = NewMovieFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMovie()
			y.Id = s
		}

		x.WorkPresented = &y

		return
	}
}

func NewScreeningEventFromContext(ctx jsonld.Context) (x schema.ScreeningEvent) {
	x.Type = "http://schema.org/ScreeningEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToScreeningEventTrait(ctx, &x.ScreeningEventTrait)
	MapCustomToScreeningEventTrait(ctx, &x.ScreeningEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToScreeningEventTrait(ctx jsonld.Context, x *schema.ScreeningEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicScreeningEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToScreeningEventTrait(ctx jsonld.Context, x *schema.ScreeningEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customScreeningEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}