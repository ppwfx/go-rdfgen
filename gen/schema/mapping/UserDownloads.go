package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserDownloads strings.Replacer
var strconvUserDownloads strconv.NumError

var basicUserDownloadsTraitMapping = map[string]func(*schema.UserDownloadsTrait, []string){}
var customUserDownloadsTraitMapping = map[string]func(*schema.UserDownloadsTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserDownloads", func(ctx jsonld.Context) (interface{}) {
		return NewUserDownloadsFromContext(ctx)
	})

}

func NewUserDownloadsFromContext(ctx jsonld.Context) (x schema.UserDownloads) {
	x.Type = "http://schema.org/UserDownloads"
	MapBasicToUserInteractionTrait(ctx, &x.UserInteractionTrait)
	MapCustomToUserInteractionTrait(ctx, &x.UserInteractionTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserDownloadsTrait(ctx, &x.UserDownloadsTrait)
	MapCustomToUserDownloadsTrait(ctx, &x.UserDownloadsTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserDownloadsTrait(ctx jsonld.Context, x *schema.UserDownloadsTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserDownloadsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserDownloadsTrait(ctx jsonld.Context, x *schema.UserDownloadsTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserDownloadsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}