package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAllocateAction strings.Replacer
var strconvAllocateAction strconv.NumError

var basicAllocateActionTraitMapping = map[string]func(*schema.AllocateActionTrait, []string){}
var customAllocateActionTraitMapping = map[string]func(*schema.AllocateActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AllocateAction", func(ctx jsonld.Context) (interface{}) {
		return NewAllocateActionFromContext(ctx)
	})



	customAllocateActionTraitMapping["http://schema.org/purpose"] = func(x *schema.AllocateActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Purpose, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Purpose = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Purpose = s
		}
	}
}

func NewAllocateActionFromContext(ctx jsonld.Context) (x schema.AllocateAction) {
	x.Type = "http://schema.org/AllocateAction"
	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAllocateActionTrait(ctx, &x.AllocateActionTrait)
	MapCustomToAllocateActionTrait(ctx, &x.AllocateActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAllocateActionTrait(ctx jsonld.Context, x *schema.AllocateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAllocateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAllocateActionTrait(ctx jsonld.Context, x *schema.AllocateActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAllocateActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}