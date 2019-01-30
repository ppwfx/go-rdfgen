package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPublicationVolume strings.Replacer
var strconvPublicationVolume strconv.NumError

var basicPublicationVolumeTraitMapping = map[string]func(*schema.PublicationVolumeTrait, []string){}
var customPublicationVolumeTraitMapping = map[string]func(*schema.PublicationVolumeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PublicationVolume", func(ctx jsonld.Context) (interface{}) {
		return NewPublicationVolumeFromContext(ctx)
	})




	basicPublicationVolumeTraitMapping["http://schema.org/pagination"] = func(x *schema.PublicationVolumeTrait, s []string) {
		x.Pagination = s[0]
	}



	customPublicationVolumeTraitMapping["http://schema.org/pageEnd"] = func(x *schema.PublicationVolumeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PageEnd, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PageEnd = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PageEnd = s
		}
	}

	customPublicationVolumeTraitMapping["http://schema.org/pageStart"] = func(x *schema.PublicationVolumeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PageStart, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PageStart = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PageStart = s
		}
	}

	customPublicationVolumeTraitMapping["http://schema.org/volumeNumber"] = func(x *schema.PublicationVolumeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.VolumeNumber, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.VolumeNumber = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.VolumeNumber = s
		}
	}
}

func NewPublicationVolumeFromContext(ctx jsonld.Context) (x schema.PublicationVolume) {
	x.Type = "http://schema.org/PublicationVolume"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPublicationVolumeTrait(ctx, &x.PublicationVolumeTrait)
	MapCustomToPublicationVolumeTrait(ctx, &x.PublicationVolumeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPublicationVolumeTrait(ctx jsonld.Context, x *schema.PublicationVolumeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPublicationVolumeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPublicationVolumeTrait(ctx jsonld.Context, x *schema.PublicationVolumeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPublicationVolumeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}