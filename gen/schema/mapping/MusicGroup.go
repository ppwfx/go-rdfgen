package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicGroup strings.Replacer
var strconvMusicGroup strconv.NumError

var basicMusicGroupTraitMapping = map[string]func(*schema.MusicGroupTrait, []string){}
var customMusicGroupTraitMapping = map[string]func(*schema.MusicGroupTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicGroup", func(ctx jsonld.Context) (interface{}) {
		return NewMusicGroupFromContext(ctx)
	})








	customMusicGroupTraitMapping["http://schema.org/genre"] = func(x *schema.MusicGroupTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Genre, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Genre = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Genre = s
		}
	}

	customMusicGroupTraitMapping["http://schema.org/track"] = func(x *schema.MusicGroupTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Track, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Track = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Track = s
		}
	}

	customMusicGroupTraitMapping["http://schema.org/tracks"] = func(x *schema.MusicGroupTrait, ctx jsonld.Context, s string){
		var y schema.MusicRecording
		if strings.HasPrefix(s, "_:") {
			y = NewMusicRecordingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicRecording()
			y.Id = s
		}

		x.Tracks = &y

		return
	}

	customMusicGroupTraitMapping["http://schema.org/album"] = func(x *schema.MusicGroupTrait, ctx jsonld.Context, s string){
		var y schema.MusicAlbum
		if strings.HasPrefix(s, "_:") {
			y = NewMusicAlbumFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicAlbum()
			y.Id = s
		}

		x.Album = &y

		return
	}

	customMusicGroupTraitMapping["http://schema.org/albums"] = func(x *schema.MusicGroupTrait, ctx jsonld.Context, s string){
		var y schema.MusicAlbum
		if strings.HasPrefix(s, "_:") {
			y = NewMusicAlbumFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicAlbum()
			y.Id = s
		}

		x.Albums = &y

		return
	}

	customMusicGroupTraitMapping["http://schema.org/musicGroupMember"] = func(x *schema.MusicGroupTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.MusicGroupMember = &y

		return
	}
}

func NewMusicGroupFromContext(ctx jsonld.Context) (x schema.MusicGroup) {
	x.Type = "http://schema.org/MusicGroup"
	MapBasicToPerformingGroupTrait(ctx, &x.PerformingGroupTrait)
	MapCustomToPerformingGroupTrait(ctx, &x.PerformingGroupTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicGroupTrait(ctx, &x.MusicGroupTrait)
	MapCustomToMusicGroupTrait(ctx, &x.MusicGroupTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicGroupTrait(ctx jsonld.Context, x *schema.MusicGroupTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicGroupTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicGroupTrait(ctx jsonld.Context, x *schema.MusicGroupTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicGroupTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}