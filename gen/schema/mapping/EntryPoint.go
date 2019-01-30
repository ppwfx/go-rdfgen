package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEntryPoint strings.Replacer
var strconvEntryPoint strconv.NumError

var basicEntryPointTraitMapping = map[string]func(*schema.EntryPointTrait, []string){}
var customEntryPointTraitMapping = map[string]func(*schema.EntryPointTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EntryPoint", func(ctx jsonld.Context) (interface{}) {
		return NewEntryPointFromContext(ctx)
	})





	basicEntryPointTraitMapping["http://schema.org/contentType"] = func(x *schema.EntryPointTrait, s []string) {
		x.ContentType = s[0]
	}


	basicEntryPointTraitMapping["http://schema.org/encodingType"] = func(x *schema.EntryPointTrait, s []string) {
		x.EncodingType = s[0]
	}


	basicEntryPointTraitMapping["http://schema.org/httpMethod"] = func(x *schema.EntryPointTrait, s []string) {
		x.HttpMethod = s[0]
	}


	basicEntryPointTraitMapping["http://schema.org/urlTemplate"] = func(x *schema.EntryPointTrait, s []string) {
		x.UrlTemplate = s[0]
	}


	customEntryPointTraitMapping["http://schema.org/actionApplication"] = func(x *schema.EntryPointTrait, ctx jsonld.Context, s string){
		var y schema.SoftwareApplication
		if strings.HasPrefix(s, "_:") {
			y = NewSoftwareApplicationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewSoftwareApplication()
			y.Id = s
		}

		x.ActionApplication = &y

		return
	}

	customEntryPointTraitMapping["http://schema.org/actionPlatform"] = func(x *schema.EntryPointTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ActionPlatform, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ActionPlatform = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ActionPlatform = s
		}
	}

	customEntryPointTraitMapping["http://schema.org/application"] = func(x *schema.EntryPointTrait, ctx jsonld.Context, s string){
		var y schema.SoftwareApplication
		if strings.HasPrefix(s, "_:") {
			y = NewSoftwareApplicationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewSoftwareApplication()
			y.Id = s
		}

		x.Application = &y

		return
	}
}

func NewEntryPointFromContext(ctx jsonld.Context) (x schema.EntryPoint) {
	x.Type = "http://schema.org/EntryPoint"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEntryPointTrait(ctx, &x.EntryPointTrait)
	MapCustomToEntryPointTrait(ctx, &x.EntryPointTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEntryPointTrait(ctx jsonld.Context, x *schema.EntryPointTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEntryPointTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEntryPointTrait(ctx jsonld.Context, x *schema.EntryPointTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEntryPointTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}