package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTransferAction strings.Replacer
var strconvTransferAction strconv.NumError

var basicTransferActionTraitMapping = map[string]func(*schema.TransferActionTrait, []string){}
var customTransferActionTraitMapping = map[string]func(*schema.TransferActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TransferAction", func(ctx jsonld.Context) (interface{}) {
		return NewTransferActionFromContext(ctx)
	})




	customTransferActionTraitMapping["http://schema.org/fromLocation"] = func(x *schema.TransferActionTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.FromLocation = &y

		return
	}

	customTransferActionTraitMapping["http://schema.org/toLocation"] = func(x *schema.TransferActionTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.ToLocation = &y

		return
	}
}

func NewTransferActionFromContext(ctx jsonld.Context) (x schema.TransferAction) {
	x.Type = "http://schema.org/TransferAction"
	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTransferActionTrait(ctx, &x.TransferActionTrait)
	MapCustomToTransferActionTrait(ctx, &x.TransferActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTransferActionTrait(ctx jsonld.Context, x *schema.TransferActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTransferActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTransferActionTrait(ctx jsonld.Context, x *schema.TransferActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTransferActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}