package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicRelease strings.Replacer
var strconvMusicRelease strconv.NumError

var basicMusicReleaseTraitMapping = map[string]func(*schema.MusicReleaseTrait, []string){}
var customMusicReleaseTraitMapping = map[string]func(*schema.MusicReleaseTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicRelease", func(ctx jsonld.Context) (interface{}) {
		return NewMusicReleaseFromContext(ctx)
	})



	basicMusicReleaseTraitMapping["http://schema.org/catalogNumber"] = func(x *schema.MusicReleaseTrait, s []string) {
		x.CatalogNumber = s[0]
	}






	customMusicReleaseTraitMapping["http://schema.org/duration"] = func(x *schema.MusicReleaseTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.Duration = &y

		return
	}

	customMusicReleaseTraitMapping["http://schema.org/creditedTo"] = func(x *schema.MusicReleaseTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.CreditedTo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.CreditedTo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.CreditedTo = s
		}
	}

	customMusicReleaseTraitMapping["http://schema.org/musicReleaseFormat"] = func(x *schema.MusicReleaseTrait, ctx jsonld.Context, s string){
		var y schema.MusicReleaseFormatType
		if strings.HasPrefix(s, "_:") {
			y = NewMusicReleaseFormatTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicReleaseFormatType()
			y.Id = s
		}

		x.MusicReleaseFormat = &y

		return
	}

	customMusicReleaseTraitMapping["http://schema.org/recordLabel"] = func(x *schema.MusicReleaseTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.RecordLabel = &y

		return
	}

	customMusicReleaseTraitMapping["http://schema.org/releaseOf"] = func(x *schema.MusicReleaseTrait, ctx jsonld.Context, s string){
		var y schema.MusicAlbum
		if strings.HasPrefix(s, "_:") {
			y = NewMusicAlbumFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicAlbum()
			y.Id = s
		}

		x.ReleaseOf = &y

		return
	}
}

func NewMusicReleaseFromContext(ctx jsonld.Context) (x schema.MusicRelease) {
	x.Type = "http://schema.org/MusicRelease"
	MapBasicToMusicPlaylistTrait(ctx, &x.MusicPlaylistTrait)
	MapCustomToMusicPlaylistTrait(ctx, &x.MusicPlaylistTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicReleaseTrait(ctx, &x.MusicReleaseTrait)
	MapCustomToMusicReleaseTrait(ctx, &x.MusicReleaseTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicReleaseTrait(ctx jsonld.Context, x *schema.MusicReleaseTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicReleaseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicReleaseTrait(ctx jsonld.Context, x *schema.MusicReleaseTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicReleaseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}