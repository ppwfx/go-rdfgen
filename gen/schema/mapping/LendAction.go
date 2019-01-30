package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLendAction strings.Replacer
var strconvLendAction strconv.NumError

var basicLendActionTraitMapping = map[string]func(*schema.LendActionTrait, []string){}
var customLendActionTraitMapping = map[string]func(*schema.LendActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LendAction", func(ctx jsonld.Context) (interface{}) {
		return NewLendActionFromContext(ctx)
	})



	customLendActionTraitMapping["http://schema.org/borrower"] = func(x *schema.LendActionTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Borrower = &y

		return
	}
}

func NewLendActionFromContext(ctx jsonld.Context) (x schema.LendAction) {
	x.Type = "http://schema.org/LendAction"
	MapBasicToTransferActionTrait(ctx, &x.TransferActionTrait)
	MapCustomToTransferActionTrait(ctx, &x.TransferActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLendActionTrait(ctx, &x.LendActionTrait)
	MapCustomToLendActionTrait(ctx, &x.LendActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLendActionTrait(ctx jsonld.Context, x *schema.LendActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLendActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLendActionTrait(ctx jsonld.Context, x *schema.LendActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLendActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}