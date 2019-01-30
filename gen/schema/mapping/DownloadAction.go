package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDownloadAction strings.Replacer
var strconvDownloadAction strconv.NumError

var basicDownloadActionTraitMapping = map[string]func(*schema.DownloadActionTrait, []string){}
var customDownloadActionTraitMapping = map[string]func(*schema.DownloadActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DownloadAction", func(ctx jsonld.Context) (interface{}) {
		return NewDownloadActionFromContext(ctx)
	})

}

func NewDownloadActionFromContext(ctx jsonld.Context) (x schema.DownloadAction) {
	x.Type = "http://schema.org/DownloadAction"
	MapBasicToTransferActionTrait(ctx, &x.TransferActionTrait)
	MapCustomToTransferActionTrait(ctx, &x.TransferActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDownloadActionTrait(ctx, &x.DownloadActionTrait)
	MapCustomToDownloadActionTrait(ctx, &x.DownloadActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDownloadActionTrait(ctx jsonld.Context, x *schema.DownloadActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDownloadActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDownloadActionTrait(ctx jsonld.Context, x *schema.DownloadActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDownloadActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}