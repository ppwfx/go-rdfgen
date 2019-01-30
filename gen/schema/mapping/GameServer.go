package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGameServer strings.Replacer
var strconvGameServer strconv.NumError

var basicGameServerTraitMapping = map[string]func(*schema.GameServerTrait, []string){}
var customGameServerTraitMapping = map[string]func(*schema.GameServerTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GameServer", func(ctx jsonld.Context) (interface{}) {
		return NewGameServerFromContext(ctx)
	})





	customGameServerTraitMapping["http://schema.org/game"] = func(x *schema.GameServerTrait, ctx jsonld.Context, s string){
		var y schema.VideoGame
		if strings.HasPrefix(s, "_:") {
			y = NewVideoGameFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewVideoGame()
			y.Id = s
		}

		x.Game = &y

		return
	}

	customGameServerTraitMapping["http://schema.org/playersOnline"] = func(x *schema.GameServerTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.PlayersOnline = &y

		return
	}

	customGameServerTraitMapping["http://schema.org/serverStatus"] = func(x *schema.GameServerTrait, ctx jsonld.Context, s string){
		var y schema.GameServerStatus
		if strings.HasPrefix(s, "_:") {
			y = NewGameServerStatusFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewGameServerStatus()
			y.Id = s
		}

		x.ServerStatus = &y

		return
	}
}

func NewGameServerFromContext(ctx jsonld.Context) (x schema.GameServer) {
	x.Type = "http://schema.org/GameServer"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGameServerTrait(ctx, &x.GameServerTrait)
	MapCustomToGameServerTrait(ctx, &x.GameServerTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGameServerTrait(ctx jsonld.Context, x *schema.GameServerTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGameServerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGameServerTrait(ctx jsonld.Context, x *schema.GameServerTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGameServerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}