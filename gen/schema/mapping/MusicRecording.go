package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicRecording strings.Replacer
var strconvMusicRecording strconv.NumError

var basicMusicRecordingTraitMapping = map[string]func(*schema.MusicRecordingTrait, []string){}
var customMusicRecordingTraitMapping = map[string]func(*schema.MusicRecordingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicRecording", func(ctx jsonld.Context) (interface{}) {
		return NewMusicRecordingFromContext(ctx)
	})






	basicMusicRecordingTraitMapping["http://schema.org/isrcCode"] = func(x *schema.MusicRecordingTrait, s []string) {
		x.IsrcCode = s[0]
	}



	customMusicRecordingTraitMapping["http://schema.org/duration"] = func(x *schema.MusicRecordingTrait, ctx jsonld.Context, s string){
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

	customMusicRecordingTraitMapping["http://schema.org/byArtist"] = func(x *schema.MusicRecordingTrait, ctx jsonld.Context, s string){
		var y schema.MusicGroup
		if strings.HasPrefix(s, "_:") {
			y = NewMusicGroupFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicGroup()
			y.Id = s
		}

		x.ByArtist = &y

		return
	}

	customMusicRecordingTraitMapping["http://schema.org/inAlbum"] = func(x *schema.MusicRecordingTrait, ctx jsonld.Context, s string){
		var y schema.MusicAlbum
		if strings.HasPrefix(s, "_:") {
			y = NewMusicAlbumFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicAlbum()
			y.Id = s
		}

		x.InAlbum = &y

		return
	}

	customMusicRecordingTraitMapping["http://schema.org/inPlaylist"] = func(x *schema.MusicRecordingTrait, ctx jsonld.Context, s string){
		var y schema.MusicPlaylist
		if strings.HasPrefix(s, "_:") {
			y = NewMusicPlaylistFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicPlaylist()
			y.Id = s
		}

		x.InPlaylist = &y

		return
	}

	customMusicRecordingTraitMapping["http://schema.org/recordingOf"] = func(x *schema.MusicRecordingTrait, ctx jsonld.Context, s string){
		var y schema.MusicComposition
		if strings.HasPrefix(s, "_:") {
			y = NewMusicCompositionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicComposition()
			y.Id = s
		}

		x.RecordingOf = &y

		return
	}
}

func NewMusicRecordingFromContext(ctx jsonld.Context) (x schema.MusicRecording) {
	x.Type = "http://schema.org/MusicRecording"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicRecordingTrait(ctx, &x.MusicRecordingTrait)
	MapCustomToMusicRecordingTrait(ctx, &x.MusicRecordingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicRecordingTrait(ctx jsonld.Context, x *schema.MusicRecordingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicRecordingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicRecordingTrait(ctx jsonld.Context, x *schema.MusicRecordingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicRecordingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}