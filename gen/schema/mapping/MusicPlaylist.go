package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicPlaylist strings.Replacer
var strconvMusicPlaylist strconv.NumError

var basicMusicPlaylistTraitMapping = map[string]func(*schema.MusicPlaylistTrait, []string){}
var customMusicPlaylistTraitMapping = map[string]func(*schema.MusicPlaylistTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicPlaylist", func(ctx jsonld.Context) (interface{}) {
		return NewMusicPlaylistFromContext(ctx)
	})





	customMusicPlaylistTraitMapping["http://schema.org/numTracks"] = func(x *schema.MusicPlaylistTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.NumTracks = &y

		return
	}

	customMusicPlaylistTraitMapping["http://schema.org/track"] = func(x *schema.MusicPlaylistTrait, ctx jsonld.Context, s string){
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

	customMusicPlaylistTraitMapping["http://schema.org/tracks"] = func(x *schema.MusicPlaylistTrait, ctx jsonld.Context, s string){
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
}

func NewMusicPlaylistFromContext(ctx jsonld.Context) (x schema.MusicPlaylist) {
	x.Type = "http://schema.org/MusicPlaylist"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicPlaylistTrait(ctx, &x.MusicPlaylistTrait)
	MapCustomToMusicPlaylistTrait(ctx, &x.MusicPlaylistTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicPlaylistTrait(ctx jsonld.Context, x *schema.MusicPlaylistTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicPlaylistTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicPlaylistTrait(ctx jsonld.Context, x *schema.MusicPlaylistTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicPlaylistTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}