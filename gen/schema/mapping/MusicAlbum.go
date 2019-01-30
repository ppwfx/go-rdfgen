package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMusicAlbum strings.Replacer
var strconvMusicAlbum strconv.NumError

var basicMusicAlbumTraitMapping = map[string]func(*schema.MusicAlbumTrait, []string){}
var customMusicAlbumTraitMapping = map[string]func(*schema.MusicAlbumTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MusicAlbum", func(ctx jsonld.Context) (interface{}) {
		return NewMusicAlbumFromContext(ctx)
	})






	customMusicAlbumTraitMapping["http://schema.org/byArtist"] = func(x *schema.MusicAlbumTrait, ctx jsonld.Context, s string){
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

	customMusicAlbumTraitMapping["http://schema.org/albumProductionType"] = func(x *schema.MusicAlbumTrait, ctx jsonld.Context, s string){
		var y schema.MusicAlbumProductionType
		if strings.HasPrefix(s, "_:") {
			y = NewMusicAlbumProductionTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicAlbumProductionType()
			y.Id = s
		}

		x.AlbumProductionType = &y

		return
	}

	customMusicAlbumTraitMapping["http://schema.org/albumRelease"] = func(x *schema.MusicAlbumTrait, ctx jsonld.Context, s string){
		var y schema.MusicRelease
		if strings.HasPrefix(s, "_:") {
			y = NewMusicReleaseFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicRelease()
			y.Id = s
		}

		x.AlbumRelease = &y

		return
	}

	customMusicAlbumTraitMapping["http://schema.org/albumReleaseType"] = func(x *schema.MusicAlbumTrait, ctx jsonld.Context, s string){
		var y schema.MusicAlbumReleaseType
		if strings.HasPrefix(s, "_:") {
			y = NewMusicAlbumReleaseTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMusicAlbumReleaseType()
			y.Id = s
		}

		x.AlbumReleaseType = &y

		return
	}
}

func NewMusicAlbumFromContext(ctx jsonld.Context) (x schema.MusicAlbum) {
	x.Type = "http://schema.org/MusicAlbum"
	MapBasicToMusicPlaylistTrait(ctx, &x.MusicPlaylistTrait)
	MapCustomToMusicPlaylistTrait(ctx, &x.MusicPlaylistTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMusicAlbumTrait(ctx, &x.MusicAlbumTrait)
	MapCustomToMusicAlbumTrait(ctx, &x.MusicAlbumTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMusicAlbumTrait(ctx jsonld.Context, x *schema.MusicAlbumTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMusicAlbumTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMusicAlbumTrait(ctx jsonld.Context, x *schema.MusicAlbumTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMusicAlbumTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}