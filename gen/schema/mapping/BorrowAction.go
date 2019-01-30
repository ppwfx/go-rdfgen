package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBorrowAction strings.Replacer
var strconvBorrowAction strconv.NumError

var basicBorrowActionTraitMapping = map[string]func(*schema.BorrowActionTrait, []string){}
var customBorrowActionTraitMapping = map[string]func(*schema.BorrowActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BorrowAction", func(ctx jsonld.Context) (interface{}) {
		return NewBorrowActionFromContext(ctx)
	})



	customBorrowActionTraitMapping["http://schema.org/lender"] = func(x *schema.BorrowActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Lender, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Lender = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Lender = s
		}
	}
}

func NewBorrowActionFromContext(ctx jsonld.Context) (x schema.BorrowAction) {
	x.Type = "http://schema.org/BorrowAction"
	MapBasicToTransferActionTrait(ctx, &x.TransferActionTrait)
	MapCustomToTransferActionTrait(ctx, &x.TransferActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBorrowActionTrait(ctx, &x.BorrowActionTrait)
	MapCustomToBorrowActionTrait(ctx, &x.BorrowActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBorrowActionTrait(ctx jsonld.Context, x *schema.BorrowActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBorrowActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBorrowActionTrait(ctx jsonld.Context, x *schema.BorrowActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBorrowActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}