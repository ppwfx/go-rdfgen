package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsImageObject strings.Replacer
var strconvImageObject strconv.NumError

var basicImageObjectTraitMapping = map[string]func(*schema.ImageObjectTrait, []string){}
var customImageObjectTraitMapping = map[string]func(*schema.ImageObjectTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ImageObject", func(ctx jsonld.Context) (interface{}) {
		return NewImageObjectFromContext(ctx)
	})


	basicImageObjectTraitMapping["http://schema.org/caption"] = func(x *schema.ImageObjectTrait, s []string) {
		x.Caption = s[0]
	}



	basicImageObjectTraitMapping["http://schema.org/representativeOfPage"] = func(x *schema.ImageObjectTrait, s []string) {
		var err error
		x.RepresentativeOfPage, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}



	customImageObjectTraitMapping["http://schema.org/exifData"] = func(x *schema.ImageObjectTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ExifData, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ExifData = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ExifData = s
		}
	}

	customImageObjectTraitMapping["http://schema.org/thumbnail"] = func(x *schema.ImageObjectTrait, ctx jsonld.Context, s string){
		var y schema.ImageObject
		if strings.HasPrefix(s, "_:") {
			y = NewImageObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewImageObject()
			y.Id = s
		}

		x.Thumbnail = &y

		return
	}
}

func NewImageObjectFromContext(ctx jsonld.Context) (x schema.ImageObject) {
	x.Type = "http://schema.org/ImageObject"
	MapBasicToMediaObjectTrait(ctx, &x.MediaObjectTrait)
	MapCustomToMediaObjectTrait(ctx, &x.MediaObjectTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToImageObjectTrait(ctx, &x.ImageObjectTrait)
	MapCustomToImageObjectTrait(ctx, &x.ImageObjectTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToImageObjectTrait(ctx jsonld.Context, x *schema.ImageObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicImageObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToImageObjectTrait(ctx jsonld.Context, x *schema.ImageObjectTrait) () {
	for k, v := range ctx.Current {
		f, ok := customImageObjectTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}